package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arthurdiego/goexpert/desafio3/configs"
	"github.com/arthurdiego/goexpert/desafio3/internal/event/handler"
	"github.com/arthurdiego/goexpert/desafio3/internal/infra/graph"
	"github.com/arthurdiego/goexpert/desafio3/internal/infra/grpc/pb"
	"github.com/arthurdiego/goexpert/desafio3/internal/infra/grpc/service"
	"github.com/arthurdiego/goexpert/desafio3/internal/infra/web/webserver"
	"github.com/arthurdiego/goexpert/desafio3/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migration(db)

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrdersUseCase := NewListOrdersUseCase(db)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler(http.MethodPost, "/order", webOrderHandler.Create)
	webserver.AddHandler(http.MethodGet, "/order", webOrderHandler.List)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUsecase:  *listOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func migration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Erro ao criar instância do driver MySQL: %v", err)
	}

	// Instância do Migrator
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../internal/infra/database/migrations",
		"mysql", driver)
	if err != nil {
		log.Fatalf("Erro ao criar instância do migrator: %v", err)
	}

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao desfazer migração: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao executar migração: %v", err)
	}

	// fmt.Println("Migração executada com sucesso")

	// version, dirty, err := m.Version()
	// if err != nil {
	// 	log.Fatalf("Erro ao verificar versão atual do banco de dados: %v", err)
	// }

	// if dirty {
	// 	log.Printf("O banco de dados está em um estado sujo (dirty)")
	// } else {
	// 	log.Printf("Versão atual do banco de dados: %d", version)
	// }

	// rows, err := db.Query("SHOW TABLES")
	// if err != nil {
	// 	log.Fatalf("Erro ao executar SHOW TABLES: %v", err)
	// }
	// defer rows.Close()

	// fmt.Println("Tabelas no banco de dados:")
	// for rows.Next() {
	// 	var tableName string
	// 	err := rows.Scan(&tableName)
	// 	if err != nil {
	// 		log.Fatalf("Erro ao ler nome da tabela: %v", err)
	// 	}
	// 	fmt.Println(tableName)
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Fatalf("Erro ao iterar sobre linhas do resultado: %v", err)
	// }
}

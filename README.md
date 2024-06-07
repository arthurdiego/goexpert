# goexpert

## Desafio 1

Execute na ordem os seguintes comandos para rodar o desafio 1, com maiores informações em: https://github.com/arthurdiego/goexpert/blob/main/desafio1/info.txt

```
cd ./desafio1/
go run ./server/main.go
go run ./client/main.go
```

## Desafio 2

Execute na ordem os seguintes comandos para rodar o desafio 2, com maiores informações em: https://github.com/arthurdiego/goexpert/blob/main/desafio2/info.txt

```
cd ./desafio2/
go run cmd/main.go -cep 56903000
```

## Desafio 2

Execute na ordem os seguintes comandos para rodar o desafio 2, com maiores informações em: https://github.com/arthurdiego/goexpert/blob/main/desafio2/info.txt

```
docker-compose up -d
cd cmd/ordersystem
go run main.go wire_gen.go
```

### WEB

Para testar os comandos web, foram criados os arquivos `desafio3/api/create_order.http` e `desafio3/api/list_orders.http` com a possibilidade de enviar o Request diretamente pelo VSCode. Caso prefira, pode ser utilizado o Postman.

### GRPC

Para executar os comando via GRPC siga os seguintes passos:

```
evans -r repl
package pb
service OrderService
```

Caso queira criar uma Order, exceute o comando `call CreateOrders`, caso queira listar as Orders, execute o comando `call ListOrders`.

### GRAPHQL

Para executar os comandos via GRAPHQL, vá ao navegador e acesse `http://localhost:8080/` e execute as seguintes queries:

Para criar uma Order:

```graphql
mutation createOrder{
  createOrder (input: {id: "gql1", Price: 32.5, Tax:2.3}){
    id
    Price
    Tax
    FinalPrice
  }
}
```

Para listar as Orders:

```graphql
query {
  listOrders{
    id
    Price
    Tax
    FinalPrice
  }
}
```

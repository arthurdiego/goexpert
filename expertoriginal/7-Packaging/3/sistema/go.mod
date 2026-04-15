module github.com/devfullcycle/goexpert/7-Packaging/3/sistema

go 1.19

// go mod edit -replace github.com/devfullcycle/goexpert/7-Packaging/3/math=../math
replace github.com/devfullcycle/goexpert/7-Packaging/3/math => ../math

require github.com/devfullcycle/goexpert/7-Packaging/3/math v0.0.0-00010101000000-000000000000

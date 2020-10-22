module dao_test

go 1.15

replace go-dao => ../

replace go-dao/mock => ../mock

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/lib/pq v1.8.0
	github.com/stretchr/testify v1.6.1
	go-dao v0.0.0-00010101000000-000000000000
	go-dao/mock v0.0.0-00010101000000-000000000000
)

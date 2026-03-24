package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// função para iniciar conexão com banco de dados, retona DB ou erro
func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5434 user=ariel password=bacon dbname=api_rest_go sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	fmt.Println("connected")

	return db, err
}

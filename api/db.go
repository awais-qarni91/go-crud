package api

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DbConnection *sql.DB
var err error

// InitializeDB connect database
func InitializeDB() error {

	DbConnection, err = sql.Open("mysql", "awais:golang456@tcp(localhost:3306)/db-crud")
	err = CreateProductTable()
	if err != nil {
		fmt.Println("Error accessing table")
		return err
	}

	return err
}

func CloseDB() {
	DbConnection.Close()
}

func CreateProductTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "CREATE TABLE IF NOT EXISTS `products` (`id` int primary key NOT NULL AUTO_INCREMENT, `name` varchar(50) NOT NULL, `created_datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP)"
	_, err := DbConnection.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating TodoItem table", err)
		return err
	}
	return err
}

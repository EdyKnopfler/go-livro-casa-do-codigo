package database

import (
	"os"
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"  // carrega porém não referencia
)

func Conectar() (*sql.DB) {
	host := os.Getenv("DB_HOST")
    porta := os.Getenv("DB_PORT")
    usuario := os.Getenv("DB_USER")
    senha := os.Getenv("DB_PASSWORD")

    db, err := sql.Open("postgres",
        fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=encurtador sslmode=disable",
            host, porta, usuario, senha))
    
    if err != nil {
        panic(err)
    }

    db.SetMaxOpenConns(5)  // Depende do quanto contratei no Heroku ;)
    db.SetMaxIdleConns(2)

    err = db.Ping()

    if err != nil {
        panic(err)
    }

    return db
}

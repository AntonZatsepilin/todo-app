package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable       = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "user_lists"
	todoItemsTable  = "todo_items"
	listsItemTable  = "listd_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

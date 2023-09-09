package lib

import (
	"github.com/jackc/pgx"
	"log"
	"strconv"
)

type DataBase struct {
	*pgx.Conn
}

func NewDatabase(env Env) DataBase {
	port, _ := strconv.Atoi(env.DBPort)
	config := pgx.ConnConfig{
		Host:     env.DBHost,
		Port:     uint16(port),
		User:     env.DBUsername,
		Password: env.DBPassword,
		Database: env.DBName,
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД!\n%+v", err)
	}

	log.Println("Подключение к БД произошло успешно!")

	return DataBase{conn}
}

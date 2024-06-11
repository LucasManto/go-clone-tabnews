package main

import (
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

func Query(sql string, dest any) error {
	portEnvVarValue := os.Getenv("POSTGRES_PORT")
	port, err := strconv.Atoi(portEnvVarValue)
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     uint16(port),
		User:     os.Getenv("POSTGRES_USER"),
		Database: os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.QueryRow(sql).Scan(dest)
}

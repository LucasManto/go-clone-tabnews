package main

import (
	"context"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func query(ctx context.Context, sql string, dest any) error {
	sb := strings.Builder{}
	sb.WriteString("host=" + os.Getenv("POSTGRES_HOST"))
	sb.WriteString(" port=" + os.Getenv("POSTGRES_PORT"))
	sb.WriteString(" user=" + os.Getenv("POSTGRES_USER"))
	sb.WriteString(" password=" + os.Getenv("POSTGRES_PASSWORD"))
	sb.WriteString(" dbname=" + os.Getenv("POSTGRES_DB"))
	connString := sb.String()

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	return conn.QueryRow(ctx, sql).Scan(dest)
}

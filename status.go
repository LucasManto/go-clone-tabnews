package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type statusResponse struct {
	UpdatedAt    *time.Time    `json:"updated_at,omitempty"`
	Dependencies *dependencies `json:"dependencies,omitempty"`
}

type dependencies struct {
	Database *database `json:"database,omitempty"`
}

type database struct {
	Version         string `json:"database_version,omitempty"`
	MaxConnections  int    `json:"max_connections,omitempty"`
	OpenConnections int64  `json:"open_connections,omitempty"`
}

func status(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	updatedAt := time.Now().UTC()
	databaseVersionQueryResult, err := query(ctx, "SHOW server_version;")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	databaseVersion := databaseVersionQueryResult[0]["server_version"].(string)

	databaseMaxConnectionsQueryResult, err := query(ctx, "SHOW max_connections;")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	databaseMaxConnections, err := strconv.Atoi(databaseMaxConnectionsQueryResult[0]["max_connections"].(string))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	databaseOpenConnections, err := query(ctx, "SELECT count(*) FROM pg_stat_activity WHERE datname = $1;", os.Getenv("POSTGRES_DB"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	openConnections := databaseOpenConnections[0]["count"].(int64)

	response := statusResponse{
		UpdatedAt: &updatedAt,
		Dependencies: &dependencies{
			Database: &database{
				Version:         databaseVersion,
				MaxConnections:  databaseMaxConnections,
				OpenConnections: openConnections,
			},
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

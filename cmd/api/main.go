package main

import (
	"errors"
	"local/panda-killer/cmd/config"
	"local/panda-killer/pkg/domain/usecase"
	"local/panda-killer/pkg/gateway/algorithms"
	"local/panda-killer/pkg/gateway/db/postgres"
	"local/panda-killer/pkg/gateway/repository"
	"local/panda-killer/pkg/gateway/rest"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Server is starting ...")

	conn, err := waitPostgres()
	if err != nil {
		panic(err)
	}
	postgres.RunMigrations()

	accountRepo := repository.NewAccountRepo(conn)
	transferRepo := repository.NewTransferRepo(conn)
	router := rest.CreateRouter(
		usecase.NewAccountUsecase(accountRepo, algorithms.AccountSecurityAlgorithmsImpl{}),
		usecase.NewTransferUsecase(transferRepo, accountRepo),
	)

	port, err := config.GetRestApiPort()
	if err != nil {
		panic(err)
	}

	log.Info("Server have started!")
	err = http.ListenAndServe(port, router)
	log.Fatal(err)
}

func waitPostgres() (*pgx.Conn, error) {
	for i := 0; i < 3; i++ {
		conn, err := postgres.OpenConnection()
		if err == nil {
			return conn, nil
		}
		time.Sleep(3 * time.Second)
	}
	return nil, errors.New("failed to connect to Postgres in 9 seconds")
}

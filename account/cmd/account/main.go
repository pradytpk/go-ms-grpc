package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pradytpk/go-ms-grpc/account"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(i int) error {
		fmt.Println("Database URL:", cfg.DatabaseURL)
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	defer r.Close()
	log.Println("listening on port 8080")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}

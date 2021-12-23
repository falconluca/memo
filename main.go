package main

import (
	"fmt"
	"github.com/shaohsiung/memo/internal/pkg/config"
	"log"
)

func main() {
	conf, err := config.Load("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", conf.DSN)
	fmt.Printf("%+v\n", conf.GRPC)
}

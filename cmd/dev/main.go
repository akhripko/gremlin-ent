package main

import (
	"context"
	"log"

	"github.com/akhripko/gremlin-ent/src/options"
	"github.com/akhripko/gremlin-ent/src/service"
	_ "github.com/facebook/ent/entc"
)

func main() {
	// read service config from os env
	config := options.ReadEnv()

	// init logger

	log.Print("begin")
	log.Print("gremlin addr: ", config.GremlinAddr)

	ctx := context.Background()

	_, err := service.New(ctx, config.GremlinAddr)
	if err != nil {
		log.Fatal("init service err:", err.Error())
	}

	//if err := fillDB(s); err != nil {
	//	log.Fatal("fillDB err:", err.Error())
	//}

	log.Print("end")
}

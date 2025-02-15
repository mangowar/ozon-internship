package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"shortener/config"
	"shortener/server"
	"shortener/shorten"
	"shortener/storage"
	"time"
)

func main() {
	var (
		srv        *shorten.Service
		flag_value string
		st         shorten.Storage
	)

	flag.StringVar(&flag_value, "s", "", "contains type of storage")
	flag.Parse()

	if flag_value == "database" {
		var err error
		for i := 0; i < 10; i++ {
			st, err = storage.New(config.Get().DB.DSN)
			if err == nil {
				break
			}
			time.Sleep(time.Second)
		}
		if err != nil {
			fmt.Println(err)
			log.Fatal("Can't access database")
		}
	} else if flag_value == "memory" {
		st = storage.InternalMemory{}
	} else {
		log.Fatalf("Wrong flag: %s", flag_value)
	}

	srv = shorten.NewService(st)
	http.HandleFunc("/", server.MainHandler(srv))
	go func() {
		fmt.Println("Starting server...")
		if err := http.ListenAndServe(config.Get().ListenAddr(), nil); err != nil {
			fmt.Printf("Listen and Serve: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	log.Println("server started")
	<-quit

}

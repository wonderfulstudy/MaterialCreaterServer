package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wonderfulstudy/MaterialCreaterServer/db"
	"github.com/wonderfulstudy/MaterialCreaterServer/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}

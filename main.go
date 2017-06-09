package main

import (
    "log"
    "net/http"
	
	"rbacapi/repository"
	"rbacapi/api"
)

func main() {

	repository.LoadData()

    router := api.NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
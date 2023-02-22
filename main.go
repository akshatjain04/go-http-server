package main

import (
	"fmt"
	"net/http"
	"taalhach/go-http-server/configs"
	"taalhach/go-http-server/database"
)

func main() {
	dbConfigs, err := configs.ParseDBConfigs()
	if err != nil {
		panic(err)
	}

	dbConn, err := database.ConnectDatabase(dbConfigs)
	if err != nil {
		panic(err)
	}

	router := NewRouter(dbConn)
	http.HandleFunc("/ping", router.ping)
	http.HandleFunc("/nonce", router.dbNonce)
	fmt.Println("Listening on port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"taalhach/go-http-server/configs"
	"taalhach/go-http-server/database"
	"taalhach/go-http-server/router"
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

	router := router.NewRouter(dbConn)
	http.HandleFunc("/ping", router.Ping)
	http.HandleFunc("/nonce", router.DBNonce)
	fmt.Println("Listening on port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}

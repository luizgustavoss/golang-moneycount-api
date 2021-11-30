package main

import (
	"fmt"
	"log"
	"moneycount-api/src/interfaces/rest/router"
	"net/http"
)

func main() {
	fmt.Println("MoneyCount API")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}

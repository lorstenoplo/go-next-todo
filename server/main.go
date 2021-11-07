package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lorstenoplo/go-next-todo/router"
)

func main() {
	r := router.Router()

	fmt.Println("Server is running on port 9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}

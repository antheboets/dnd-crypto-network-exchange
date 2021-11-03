package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

	fmt.Println("starting server")

	http.Handle("/", http.FileServer(http.Dir("./resources/test")))

    log.Fatal(http.ListenAndServe(":8081", nil))

}
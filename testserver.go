package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		startServer()
		wg.Done()
	}()

	wg.Wait()
}
func startServer() {
	fmt.Println("starting server")
	http.Handle("/", http.FileServer(http.Dir("./resources/test")))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

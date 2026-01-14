package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	app, err := App()
	if err != nil {
		panic(err)
	}
	defer app.Close()

	fmt.Println("Server started on port", app.Config.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: nil, // TODO: Add handler
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

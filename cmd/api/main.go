package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nerdydave2017/filesharing/internal/bootstrap"
)

func main() {
	app, err := bootstrap.App()
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

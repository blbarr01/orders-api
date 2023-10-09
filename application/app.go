package application

import (
	"context"
	"fmt"
	"net/http"
	"github.com/redis/go-redis/v9"
)

// define struture of the app
type App struct {
	router http.Handler
	rdb *redis.Client
}

//create construtor
func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb: redis.NewClient(&redis.Options{}),
	}
	return app
}

//main function to run the app
func (a* App) Start(ctx context.Context) error {
	server := http.Server{
		Addr: ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil{
		return fmt.Errorf("failed to connect to redis %w", err)	
	}

	fmt.Printf("server listening on port %s", server.Addr)

	err = server.ListenAndServe()
	if err != nil{
		return fmt.Errorf("failed to start server: %w", err) 
	}
	return nil
}


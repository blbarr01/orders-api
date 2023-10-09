package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/blbarr01/orders-api/application"
)

func main(){
	app := application.New()
	// background context only to be used w/in main
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)	
	defer cancel()

	err := app.Start(ctx)
	if err != nil{
		fmt.Println("failed to start the app", err)
	}


}





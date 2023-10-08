package main

import (
	"fmt"
	"context"
	"github.com/blbarr01/orders-api/application"
)

func main(){
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil{
		fmt.Println("failed to start the app", err)
	}


}





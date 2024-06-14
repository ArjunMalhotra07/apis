package gochi

import (
	"context"
	"fmt"

	"github.com/ArjunMalhotra07/apis.git/application"
)

func MainFunction() {
	app := application.New()
	err := app.StartServer(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
	
}

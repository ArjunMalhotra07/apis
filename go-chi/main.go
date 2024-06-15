package gochi

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ArjunMalhotra07/apis.git/application"
)

func MainFunction() {
	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.StartServer(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}

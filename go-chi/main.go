package gochi

import (
	"context"
	"fmt"

	"github.com/ArjunMalhotra07/apis.git/application"
)

func MainFunction() {
	app := application.New()
	// //! Method 1
	err := app.StartServer(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
	//! Method 2
	/*
		err := http.ListenAndServe(":8080", router)
	*/

}

package main

import (
	"context"
	"fmt"
	"github.com/baijugoradia/gohelloworld/application"
)

func main() {
	fmt.Println("Server Starting")
	app := new(application.App)

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Errorf("Some issue with serve : %s", err)
	}
}

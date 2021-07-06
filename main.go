package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func auth() (nil, error) {
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
}

func main() {
	fmt.Println("Hello, World")
}

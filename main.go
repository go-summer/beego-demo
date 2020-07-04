package main

import (
	"context"
	"github.com/go-summer/summer"
	"net/http"
)

func main() {
	summer.App.SetController("/hello", func(ctx context.Context, request *http.Request) string {
		return "hello"
	})
	summer.App.SetPort("8080")
	summer.App.Run()
}

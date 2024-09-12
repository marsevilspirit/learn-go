package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "userName", "mars")

	handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
	userName := ctx.Value("userName")
	fmt.Println("User Name:", userName)
}

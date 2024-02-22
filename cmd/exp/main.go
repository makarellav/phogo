package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	withValue := context.WithValue(ctx, "abc", 123)

	value, ok := withValue.Value("abc").(int)

	if !ok {
		fmt.Println("bruh")
	}

	fmt.Println(value)
}

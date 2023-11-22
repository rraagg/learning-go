package main

import (
	"chapterTwelve"
	"context"
	"fmt"
)

func logic(ctx context.Context, info string) (string, error) {
	fmt.Printf("I did some interesting %s \n", info)
	return "finished", nil
}

func main() {
	ss := chapterTwelve.SlowServer()
	fmt.Println("hello")

	ctx := context.Background()
	logicResult, err := logic(ctx, "stuff")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(logicResult)
}

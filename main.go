package main

import (
	"context"
	"projeto-final/infrastructure/setup"
	"sync"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	setup.
		NewConfig().
		InitLogger().
		WithAppConfig().
		WithDB().
		WithRouter().
		WithWebServer().
		Start(ctx, &wg)

	wg.Wait()
}

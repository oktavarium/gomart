package main

import (
	"fmt"

	"github.com/oktavarium/gomart/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(fmt.Errorf("app closed in wrong way: %w", err))
	}
}

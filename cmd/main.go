package main

import (
	"fmt"
	"github.com/douyu/jupiter"
	"home/internal/engine"
	"log"
)

func main() {
	eng := engine.NewEngine()
	eng.RegisterHooks(jupiter.StageAfterStop, func() error {
		fmt.Println("exit jupiter app ...")
		return nil
	})

	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}

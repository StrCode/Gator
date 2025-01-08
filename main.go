package main

import (
	"fmt"
	"log"
	"os"

	"github.com/StrCode/Gator/internal/comm"
	"github.com/StrCode/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	newState := comm.State{
		Cfg: &cfg,
	}

	items := make(map[string]func(*comm.State, comm.Command) error)
	commands := comm.Commands{
		Items: items,
	}

	commands.Register("login", comm.HandlerLogin)

	words := os.Args

	if len(words) < 2 {
		log.Fatalf("Not enough arguments were provided")
		os.Exit(0)
	}

	commands.Run(&newState, comm.Command{
		words[1],
		words[1:],
	})

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}

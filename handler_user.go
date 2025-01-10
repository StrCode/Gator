package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/StrCode/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	if user.Name == "" {
		fmt.Println("couldn't set find the user: %v", name)
		os.Exit(1)
	}

	err_n := s.cfg.SetUser(name)
	if err_n != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		uuid.New(),
		time.Now().UTC(),
		time.Now().UTC(),
		name,
	})
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s.cfg.SetUser(user.Name)
	fmt.Println("User has been created successfully!")
	fmt.Printf("New user: %+v\n", user)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}

		fmt.Printf("* %v\n", user.Name)
	}

	return nil
}

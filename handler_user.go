package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tumpech/gator/internal/database"
)

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]
	arg := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username}
	dbUser, err := s.db.CreateUser(context.Background(), arg)
	if err != nil {
		return fmt.Errorf("couldn't create user'%s': %w", username, err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User is created\n")
	printUser(dbUser)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user '%s' doesn't exists", username)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("User is set to %s\n", username)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users: %w", err)
	}

	fmt.Println("Users table is reset.")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users: %w", err)
	}

	currentUser := s.cfg.CurrentUserName

	for _, dbUser := range dbUsers {
		if dbUser.Name == currentUser {
			fmt.Printf("* %s (current)\n", dbUser.Name)
		} else {
			fmt.Printf("* %s\n", dbUser.Name)
		}
	}

	return nil
}

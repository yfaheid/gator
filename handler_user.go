package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/yfaheid/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("incorrect amount of args")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		os.Exit(1)
		return err
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("user set to " + cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("incorrect amount of args")
	}
	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0]})
	if err != nil {
		os.Exit(1)
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("user set to " + cmd.args[0])
	fmt.Println(newUser)
	return nil
}

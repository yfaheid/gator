package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/yfaheid/gator/internal/config"
	"github.com/yfaheid/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	newState := state{cfg: &cfg, db: dbQueries}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	if len(os.Args) < 2 {
		fmt.Println("no command")
		os.Exit(1)
	}
	name := os.Args[1]
	args := os.Args[2:]
	cmd := command{name: name, args: args}

	err = cmds.run(&newState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

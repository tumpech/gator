package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/tumpech/gator/internal/config"
	"github.com/tumpech/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{cfg: &cfg}
	db, err := sql.Open("postgres", programState.cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %s", err)
	}
	dbQueries := database.New(db)
	programState.db = dbQueries
	cmds := NewCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

	programArgs := os.Args
	if len(programArgs) < 2 {
		log.Fatal("there must be a command")
	}
	cmd := command{programArgs[1], programArgs[2:]}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatalf("error executing command: %v", err)
	}
}

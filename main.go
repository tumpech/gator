package main

import (
	"context"
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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
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
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnFollow))

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

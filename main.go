package main

import (
	"log"
	"os"

	"github.com/tumpech/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{cfg: &cfg}
	cmds := NewCommands()
	cmds.register("login", handlerLogin)

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

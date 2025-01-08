package comm

import (
	"fmt"
	"os"

	"github.com/StrCode/Gator/internal/config"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Items map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		fmt.Errorf("a username is required")
		os.Exit(1)
	}

	s.Cfg.SetUser(cmd.Args[1])
	fmt.Printf("The user has been set")
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Items[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	if s == nil {
		return fmt.Errorf("no state has been initialized")
	}

	newfunc := c.Items[cmd.Name]

	if err := newfunc(s, cmd); err != nil {
		return fmt.Errorf("this is not working")
	}

	return nil
}

package command

func NewStatusCommand() *Command {
	return &Command{
		[]string{"status"},
		StatusHandler,
		[]func() error{},
		[]func() error{},
	}
}

func StatusHandler() error {
	return nil
}

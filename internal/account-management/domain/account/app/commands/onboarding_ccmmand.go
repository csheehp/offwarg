package commands

type accountOnboardingCommand struct{}

func (a *accountOnboardingCommand) Handle() error {
	return nil
}

func NewAccountOnboardingCommand() *accountOnboardingCommand {
	return &accountOnboardingCommand{}
}

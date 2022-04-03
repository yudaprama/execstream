package execstream

func NewExecutorFromName() (*executor, error) {
	return &executor{"bash", []string{"-c"}}, nil
}

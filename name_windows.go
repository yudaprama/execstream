package execstream

func NewExecutorFromName() (*executor, error) {
	return &executor{"cmd", []string{"/c"}}, nil
}

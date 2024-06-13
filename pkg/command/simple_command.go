package command

type SimpleCommand interface {
	Execute(invocation SimpleCommandInvocation) error
	Suggest(invocation SimpleCommandInvocation) []string
}

type SimpleCommandInvocation struct {
	Arguments []string
	Source    CommandSource
}

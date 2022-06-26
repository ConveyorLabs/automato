package core

type OnEvent struct {
	EventSignature    string
	ExecutionFunction func() string
}

func (o *OnEvent) evaluate() bool {

	return true
}

package core

type OnEvent struct {
	EventSignature    bool
	ExecutionFunction func() bool
}

func (o *OnEvent) evaluate() bool {

	return true
}

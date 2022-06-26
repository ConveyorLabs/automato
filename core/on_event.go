package core

type OnEvent struct {
	EventSignature    bool
	ExecutionFunction func() bool
}

func (o *OnEvent) evaluate() bool {

	return true
}

func (o *OnEvent) execute() bool {
	return o.ExecutionFunction()
}

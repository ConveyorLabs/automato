package core

type AutomationTask struct {
	Task Trigger
}

type Trigger interface {
	evaluate() bool
	execute() bool
}

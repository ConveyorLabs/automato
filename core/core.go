package core

type AutomationTask struct {
	Task Trigger
}

type Trigger interface {
	evaluate() bool
	execute() bool
}

func GenerateAutomationTasks() []AutomationTask {
	return []AutomationTask{}
}

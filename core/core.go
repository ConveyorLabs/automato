package core

type AutomationTask struct {
	AutomationTask Task
}

func newAutomationTask() AutomationTask {

}

type Task struct {
	trigger Trigger
	action  Action
}

type Action struct {
	isTX           bool
	messageContent MessageContent
}

type MessageContent struct {
	address           bool
	functionSignature bool
}

type Trigger interface {
	evaluate() bool
}

func GenerateAutomationTasks(ast bool) []AutomationTask {

	for trigger := range ast {
		//create new automation task

		//add to automation task list
	}

	return []AutomationTask{}
}

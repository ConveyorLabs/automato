package core

import yamlParser "automato/yaml_parser"

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

func GenerateAutomationTasks(ast *yamlParser.YamlFile) []AutomationTask {

	for _, at := range ast.AutomationTasks {
		automationTask := at
		trigger := automationTask.Trigger
		action := automationTask.Actions

	}
	//create new automation task

	//add to automation task list

	return []AutomationTask{}
}

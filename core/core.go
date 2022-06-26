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

		trigger := at.Trigger

		//if the trigger is a when block
		if trigger == yamlParser.WHENBLockType{
			
			newWhenBlock:=WhenBlock{}


		} else if trigger == &yamlParser.Trigger{}{
			
		}else if trigger == OtherTypefromparser{}



	}
	//create new automation task

	//add to automation task list

	return []AutomationTask{}
}

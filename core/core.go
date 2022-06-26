package core

import (
	yamlParser "automato/yaml_parser"
	"math/big"
)

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
		//Generate all automation tasks in new Automation tasks array  
		trigger := at.Trigger
		actions := at.Actions
		//if the trigger is a when block
		if trigger == yamlParser.WHENBLockType{
			
			newWhenBlock:=WhenBlock{big.NewInt(trigger.WhenBlock)}
			for _, action := range at.Actions.Actions {
				 
			}


		} else if trigger == yamlParser.Trigger.OnEvent{
			// newOnEvent:= OnEvent{trigger.OnEvent}

		}else if trigger == yamlParser.Trigger.SecondsInterval{
			newSecondsInterval := 
		}

		// trigger := automationTask.Trigger
		// action := automationTask.Actions
		// messageContent := MessageContent{}
		// at.Actions.Actions
	}
	//create new automation task

	//add to automation task list

	return []AutomationTask{}
}

package main

type ActionAvailableFn func() bool
type ActionCommandFn func(*Node)

type Action struct {
	Name        string
	Description string
	NodeType    string
	CommandFn   ActionCommandFn
}

type ActionInstance struct {
	Action *Action
	Node   *Node
}

var actions []Action
var actionInstances []ActionInstance

func RegisterActions() {
	RegisterActivePortScan()
	RegisterPassivePortScan()
}

func RegisterAction(action Action) {
	actions = append(actions, action)
}

func (action *Action) Instantiate(node *Node) ActionInstance {
	instance := ActionInstance{
		action,
		node,
	}

	actionInstances = append(actionInstances, instance)
	instance.Run()
	return instance
}

func (instance *ActionInstance) Run() {
	instance.Action.CommandFn(instance.Node)
}

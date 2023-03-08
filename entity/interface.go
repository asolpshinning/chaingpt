package entity

type Agent struct {
	Name   string
	Output string
}

type AgentResponse struct {
	Input  string
	Output string
}

type Tool struct {
	Name  string
	Value string
}

type Chain struct {
	Name string
}

func CreateNewAgent(name string) *Agent {
	return &Agent{
		Name: name,
	}
}

package entity

type Agent struct {
	Name   string
	Output string
}

type AgentResponse struct {
	Input        string
	Output       string
	Satisfactory bool
}

type Tool struct {
	Type  string
	Name  string
	Value string
}

type ChainResponse struct {
	Response string
}

func CreateNewAgent(name string) *Agent {
	return &Agent{
		Name: name,
	}
}

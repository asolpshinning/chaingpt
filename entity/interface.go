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
	Name  string
	Value string
}

type Chain interface {
	BaseChain(from *Agent, input *AgentResponse, tools []Tool) error
}

type ChainResponse struct {
	Response string
}

func CreateNewAgent(name string) *Agent {
	return &Agent{
		Name: name,
	}
}

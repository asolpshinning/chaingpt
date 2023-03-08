package chains

import "github.com/asolpshinning/chaingpt/entity"

type chain struct {
	Name string
}

func CreateNewChain(name string) entity.Chain {
	return &chain{Name: name}
}

func (c *chain) BaseChain(from *entity.Agent, input *entity.AgentResponse, tools []entity.Tool) error {
	// if the tools are empty, then the call is not coming from an

	return nil
}

package gocdapi

import "fmt"

type AgentService struct {
	client *Client
}

//Agent represents a Agent index as in the response when listing agents
type Agent struct {
	Os           string   `json:"operating_system"`
	Environments []string `json:"environments"`
	UUID         string   `json:"uuid"`
	AgentName    string   `json:"agent_name"`
	Resources    []string `json:"resources"`
	Sandbox      string   `json:"sandbox"`
	Status       string   `json:"status"`
	BuildLocator string   `json:"build_locator"`
	IPAddress    string   `json:"ip_address"`
	Enabled      bool     `json:"enabled"`
}

type Embedded struct {
	Agents []Agent `json:"agents"`
}

type AgentList struct {
	Embedded Embedded `json:"_embedded"`
}

type Message struct {
	Message string `json:"message"`
}

func (s *AgentService) GetAll() (*AgentList, *Response, error) {
	u := "agents"

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		fmt.Println("ERROR 1")
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/vnd.go.cd.v1+json")

	agents := new(*AgentList)
	response, err := s.client.Do(req, agents)
	if err != nil {
		fmt.Println("ERROR 2")
		return nil, response, err
	}
	return *agents, response, err

}

func (s *AgentService) Get(UUID string) (Agent, *Response, error) {
	u := "agents/" + UUID

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return Agent{}, nil, err
	}

	req.Header.Set("Accept", "application/vnd.go.cd.v1+json")

	agent := new(Agent)
	response, err := s.client.Do(req, agent)

	if err != nil {
		fmt.Println("Error Do agents")
		return Agent{}, response, err
	}

	return *agent, response, err

}

func (s *AgentService) Delete(UUID string) (Message, *Response, error) {
	u := "agents/" + UUID

	req, err := s.client.NewRequest("DELETE", u, nil)

	if err != nil {
		return Message{Message: "ERROR CREATING REQUEST"}, nil, err
	}

	message := new(Message)
	response, err := s.client.Do(req, message)

	if err != nil {
		return Message{Message: "ERROR EXECUTING REQUEST"}, response, err
	}

	return *message, response, err

}

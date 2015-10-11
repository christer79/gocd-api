package gocdapi

type PipelineGroupService struct {
	client *Client
}

//Stage holds informaton about a Stage of a Pipeline
type Stage struct {
	Name string `json:"name"`
}

//Material represents information about material such as Git or Mercurial repo
type Material struct {
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

//Pipeline represents information about a single pipeline
type Pipeline struct {
	Name      string     `json:"name"`
	Label     string     `json:"label"`
	Materials []Material `json:"materials"`
	Stages    []Stage    `json:"stages"`
}

//PipelineGroup holds informationa bout a single pipeline group
type PipelineGroup struct {
	Name      string     `json:"name"`
	Pipelines []Pipeline `json:"pipelines"`
}

func (s *PipelineGroupService) Get() ([]PipelineGroup, *Response, error) {
	u := "config/pipeline_groups"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pipelineGroups := new([]PipelineGroup)
	response, err := s.client.Do(req, pipelineGroups)
	if err != nil {
		return nil, response, err
	}

	return *pipelineGroups, response, err

}

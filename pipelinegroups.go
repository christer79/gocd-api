package gocdapi

type PipelineGroupService struct {
	client *Client
}

func (s *PipelineGroupService) Get() ([]PipelineGroup, *Response, error) {
	u := "api/config/pipeline_groups"

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

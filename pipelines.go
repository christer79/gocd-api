package gocdapi

import (
	"fmt"
	"strconv"
)

type PipelineService struct {
	client *Client
}

func (s *PipelineService) GetInstance(coordinate Coordinate) (PipelineInstance, *Response, error) {
	u := "api/pipelines/" + coordinate.PipelineName + "/instance/" + strconv.Itoa(coordinate.PipelineCount)
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return PipelineInstance{}, nil, err
	}

	//req.Header.Set("Accept", "application/vnd.go.cd.v1+json")

	pipelineinstance := new(PipelineInstance)
	response, err := s.client.Do(req, pipelineinstance)

	if err != nil {
		fmt.Println("Error Do agents")
		return PipelineInstance{}, response, err
	}

	return *pipelineinstance, response, err

}

func (s *PipelineService) GetHistory(coordinate Coordinate) (PipelineHistory, *Response, error) {
	u := "api/pipelines/" + coordinate.PipelineName + "/history"
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return PipelineHistory{}, nil, err
	}

	//req.Header.Set("Accept", "application/vnd.go.cd.v1+json")

	pipelinehistory := new(PipelineHistory)
	response, err := s.client.Do(req, pipelinehistory)

	if err != nil {
		fmt.Println("Error Do agents")
		return PipelineHistory{}, response, err
	}

	return *pipelinehistory, response, err

}

func (s *PipelineService) GetStatus(coordinate Coordinate) (PipelineStatus, *Response, error) {
	u := "api/pipelines/" + coordinate.PipelineName + "/history"
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return PipelineStatus{}, nil, err
	}

	//req.Header.Set("Accept", "application/vnd.go.cd.v1+json")

	pipelinestatus := new(PipelineStatus)
	response, err := s.client.Do(req, pipelinestatus)

	if err != nil {
		fmt.Println("Error Do agents")
		return PipelineStatus{}, response, err
	}

	return *pipelinestatus, response, err

}

package gocdapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type ArtifactService struct {
	client *Client
}

func (s *ArtifactService) GetAllArtifacts(coordinate Coordinate) ([]Artifact, *Response, error) {
	u := "files/" + coordinate.PipelineName + "/" + strconv.Itoa(coordinate.PipelineCount) + "/" + coordinate.StageName + "/" + strconv.Itoa(coordinate.StageCount) + "/" + coordinate.JobName + ".json"
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return []Artifact{}, nil, err
	}

	artifacts := new([]Artifact)
	response, err := s.client.Do(req, artifacts)
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	fmt.Println(contents)
	if err != nil {
		fmt.Println("Error Do agents")
		return []Artifact{}, response, err
	}

	return *artifacts, response, err

}

func (s *ArtifactService) GetArtifactDirectory(coordinate Coordinate) (*bytes.Buffer, *Response, error) {
	//http://localhost:8153/go/files/AlaskaMultiNodeLassiBlue/4/BootStrap/1/BootStrap/artifacts/servers.sh
	data := new(bytes.Buffer)

	u := "files/" + coordinate.PipelineName + "/" + strconv.Itoa(coordinate.PipelineCount) + "/" + coordinate.StageName + "/" + strconv.Itoa(coordinate.StageCount) + "/" + coordinate.JobName + "/" + coordinate.FilePath + ".zip"
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return data, nil, err
	}

	response, err := s.client.Do(req, data)

	if err != nil {
		fmt.Println("Error Do agents")
		return data, response, err
	}

	return data, response, err

}

func (s *ArtifactService) GetArtifactFile(coordinate Coordinate) (*bytes.Buffer, *Response, error) {
	//TODO: This call might return 202 expecting us to poll every now and then for the file.

	//http://localhost:8153/go/files/AlaskaMultiNodeLassiBlue/4/BootStrap/1/BootStrap/artifacts/servers.sh
	data := new(bytes.Buffer)

	u := "files/" + coordinate.PipelineName + "/" + strconv.Itoa(coordinate.PipelineCount) + "/" + coordinate.StageName + "/" + strconv.Itoa(coordinate.StageCount) + "/" + coordinate.JobName + "/" + coordinate.FilePath

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return data, nil, err
	}

	response, err := s.client.Do(req, data)

	if err != nil {
		fmt.Println("Error Do agents")
		return data, response, err
	}

	return data, response, err

}

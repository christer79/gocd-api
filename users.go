package gocdapi

import "fmt"

type UserService struct {
	client *Client
}

func (s *UserService) GetAll() ([]User, *Response, error) {
	u := "api/users"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		fmt.Println("ERROR CREATOING REQUEST")
		return nil, nil, err

	}

	users := new([]User)
	response, err := s.client.Do(req, users)
	if err != nil {
		fmt.Println("ERROR EXECUTING REQUEST")
		var body []byte
		response.Body.Read(body)
		fmt.Println(body)
		return nil, response, err
	}
	var body []byte
	response.Body.Read(body)
	fmt.Println(body)
	return *users, response, err

}

package gocdapi

import "fmt"

type UserService struct {
	client *Client
}

func (s *UserService) GetAll() ([]User, *Response, error) {
	u := "api/users"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		fmt.Println("ERROR CREATING REQUEST")
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/vnd.go.cd.v1+json")

	users := new(UserList)
	response, err := s.client.Do(req, users)
	if err != nil {
		fmt.Println("ERROR EXECUTING REQUEST")
		return nil, response, err
	}

	return users.UEmbedded.Users, response, err

}

func (s *UserService) Get(loginName string) (*User, *Response, error) {
	u := "api/users/" + loginName

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		fmt.Println("ERROR CREATING REQUEST")
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/vnd.go.cd.v1+json")

	user := new(User)
	response, err := s.client.Do(req, user)
	if err != nil {
		fmt.Println("ERROR EXECUTING REQUEST")
		return nil, response, err
	}

	return user, response, err

}

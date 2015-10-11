package gocdapi

import "fmt"

type UserService struct {
	client *Client
}

type User struct {
	LoginName      string   `json:"login_name"`
	DisplayName    string   `json:"DisplayName"`
	Enabled        string   `json:"enabled"`
	Email          string   `json:"email"`
	EmailMe        bool     `json:"email_me"`
	CheckinAliases []string `json:"checkin_aliases"`
}

func (s *UserService) GetAll() ([]User, *Response, error) {
	u := "users"

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

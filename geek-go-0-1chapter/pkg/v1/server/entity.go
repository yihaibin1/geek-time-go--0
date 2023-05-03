package server


type signRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type signResponse struct {
	UserId string `json:"userid"`
	Name   string `json:"name"`
}
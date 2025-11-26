package models

type Workspace struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Credentials []Credential `json:"credentials"`
}

type Credential struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	AppName  string `json:"appName"`
	Url      string `json:"url"`
	Note     string `json:"note"`
}

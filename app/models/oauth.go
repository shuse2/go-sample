package models

type RequestToken struct {
	Token  string
	Secret string
	Url    string
}

type AccessToken struct {
	Secret         string
	Token          string
	AdditionalData map[string]string
}

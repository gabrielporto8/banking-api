package models

type Authentication struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Token struct {
	Token string `json:"token"`
}
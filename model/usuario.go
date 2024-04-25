package model

type Usuario struct {
	Nome  string `json:"nome"`
	ID    string `json:"id"`
	Email string `json:"email"`
	senha string `json:"senha"`
}

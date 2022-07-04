package models

type Contacts struct {
	Contacts []struct {
		Nome       string `json:"Name"`
		Celular    string `json:"Cellphone"`
		Id_cliente int    `json:"Id_Cliente"`
	} `json:"contacts"`
}

type ContactsList struct {
	Contacts []struct {
		Nome    string `json:"Name"`
		Celular string `json:"Cellphone"`
	} `json:"contacts"`
}

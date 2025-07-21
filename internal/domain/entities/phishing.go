package entities

type Phising struct {
	Domain string `json:"domain"`
	Active bool   `json:"active"`
}

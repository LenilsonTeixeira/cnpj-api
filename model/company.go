package model

import "time"

type Company struct {
	CNPJ        string    `json:"cnpj"`
	Name        string    `json:"name"`
	FantasyName string    `json:"fantasyName"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"createdAt"`
}

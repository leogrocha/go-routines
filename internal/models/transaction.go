package models

import "time"

type Transaction struct {
	Account Account   `json:"account"`
	Value   float64   `json:"value"`
	DataAt  time.Time `json:"dataAt"`
}

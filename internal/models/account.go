package models

type Account struct {
	Number   string  `json:"number"`
	Agency   string  `json:"agency"`
	BankName string  `json:"bankName"`
	BankCode string  `json:"bankCode"`
	Person   Person  `json:"person"`
	Balance  float64 `json:"balance"`
}

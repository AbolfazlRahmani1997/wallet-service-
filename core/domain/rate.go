package domain

type Rate struct {
	Id       string   `json:"id" gorm:"primaryKey index"`
	Amount   float64  `json:"Amount"`
	Currency Currency `json:"Currency"`
}

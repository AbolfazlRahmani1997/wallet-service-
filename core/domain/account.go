package domain

type Account struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	NationalCode string `json:"national_code"`
	PhoneNumber  string `json:"phone_number"`
}

package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/core/domain"
	"payment/core/service"
)

type AccountHandler struct {
	AccountService *service.AccountService
}
type AccountJson struct {
	NationalCode string `json:"national_code"`
	PhoneNumber  string `json:"phone_number"`
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

func (handler *AccountHandler) CreateAccount(c *gin.Context) {
	var accountJson AccountJson
	err := c.ShouldBindJSON(&accountJson)
	if err != nil {
		return
	}

	err = handler.AccountService.SaveAccount(&domain.Account{NationalCode: accountJson.NationalCode, PhoneNumber: accountJson.PhoneNumber})
	if err != nil {
		return
	}
}

func (handler *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {

}

func (handler *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {

}
func (handler *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {

}

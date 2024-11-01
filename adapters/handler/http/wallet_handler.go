package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/core/service"
)

type Handler struct {
	WalletService service.WalletService
}

func NewWalletHandler(service service.WalletService) *Handler {
	return &Handler{WalletService: service}
}

func (h *Handler) CreateWallet(c *gin.Context) {
	var json CreateWalletRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	wallet, err := h.WalletService.CreateWallet()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wallet})
}
func (h *Handler) Transfer(c *gin.Context) {
	var json TransferRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transfer, err := h.WalletService.Transfer(json.Amount, json.WalletOrigin, json.WalletDestination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": transfer})
}

type CreateWalletRequest struct {
	Password string `json:"password" binding:"required"`
}

type TransferRequest struct {
	WalletOrigin      uint    `json:"wallet_origin" binding:"required"`
	WalletDestination uint    `json:"wallet_destination" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
}

type ChargingRequest struct {
	WalletOrigin uint    `json:"wallet_origin" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
}

type ChangeRequest struct {
	WalletOrigin uint    `json:"wallet_origin" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	Currency     string  `json:"currency" binding:"required"`
}

func (ws Handler) Charging(c *gin.Context) {
	var json ChargingRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	document, err := ws.WalletService.CashIn(json.WalletOrigin, json.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": document})

}

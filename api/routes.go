package api

import (
	"github.com/gin-gonic/gin"
	"payment/adapters/handler/http"
)

var engine *gin.Engine

func InitRouter(walletHandler *http.Handler) {
	gin.SetMode(gin.DebugMode)
	engine = gin.Default()

	walletRoute := engine.Group("wallet")
	walletRoute.POST("/", walletHandler.CreateWallet)
	walletRoute.POST("/transfer", walletHandler.Transfer)
	walletRoute.POST("/charging", walletHandler.Charging)
	walletRoute.POST("/check_out", walletHandler.Charging)

}

func Start(addr string) error {
	return engine.Run(addr)
}
package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/core/service"
)

type DocumentHandler struct {
	DocumentHandler *service.DocumentService
}

func NewDocumentHandler(service *service.DocumentService) DocumentHandler {
	return DocumentHandler{DocumentHandler: service}
}

func (receiver DocumentHandler) GetAll(c *gin.Context) {
	all, err := receiver.DocumentHandler.GetAll()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, all)
}

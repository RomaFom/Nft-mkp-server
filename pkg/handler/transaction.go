package handler

import (
	"app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionInput struct {
	TxHash string `json:"tx_hash" binding:"required"`
	Wallet string `json:"wallet" binding:"required"`
}

func (h *Handler) addTransaction(c *gin.Context) {
	var input TransactionInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	txId, err := h.services.Transaction.CreateTransaction(input.Wallet, input.TxHash)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"txId": txId,
	})
}

type getAllTransactionsResponse struct {
	Data []app.Transaction `json:"data"`
}

func (h *Handler) getAllTransactions(c *gin.Context) {
	txList, err := h.services.Transaction.GetAllTransactions()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTransactionsResponse{
		Data: txList,
	})
}

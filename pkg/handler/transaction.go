package handler

import (
	"app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionInput struct {
	TxHash string `json:"tx_hash" binding:"required"`
	Wallet string `json:"wallet" binding:"required"`
	UserId int    `json:"user_id" binding:"required"`
	ItemId int    `json:"item_id" binding:"required"`
}

func (h *Handler) addTransaction(c *gin.Context) {
	var input app.Transaction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	txId, err := h.services.Transaction.CreateTransaction(input)

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

func (h *Handler) getNftTransactions(c *gin.Context) {
	nftId, err := getId(c)

	if err != nil {
		return
	}

	txList, err := h.services.Transaction.GetNftTransactions(nftId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTransactionsResponse{
		Data: txList,
	})

}

func (h *Handler) buyItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input app.BuyItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tx := app.Transaction{
		TxHash: input.TxHash,
		Wallet: input.Wallet,
		UserId: userId,
		ItemId: input.ItemId,
		NftId:  input.NftId,
	}

	txId, err := h.services.Transaction.CreateTransaction(tx)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.services.Marketplace.BuyItem(input.ItemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"txId": txId,
		"item": item,
	})

}

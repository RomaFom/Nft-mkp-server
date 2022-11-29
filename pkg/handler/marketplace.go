package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getItemCount(c *gin.Context) {
	count, err := h.services.Marketplace.GetItemCount()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"count": count,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	params := getPaginationParams(c)
	items, err := h.services.Marketplace.GetItemsForSale(params.page, params.size)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getMyListings(c *gin.Context) {
	wallet, err := h.getUserWallet(c)
	if err != nil {
		return
	}

	items, err := h.services.Marketplace.GetMyListings(wallet)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})

}

func (h *Handler) getMyPurchases(c *gin.Context) {
	wallet, err := h.getUserWallet(c)
	if err != nil {
		return
	}

	items, err := h.services.Marketplace.GetMyPurchases(wallet)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})

}

func (h *Handler) init(c *gin.Context) {
	h.services.ValidateSCItems()
	//if err != nil {
	//	newErrorResponse(c, http.StatusNotFound, "Error validating items")
	//	return
	//}
}

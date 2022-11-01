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
	count, err := h.services.Marketplace.GetItemCount()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, err := h.services.Marketplace.GetMarketplaceItems(count)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})
}

package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "user_Id"
	wallet     = "wallet"
	id         = "id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	//	parse token
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)

}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id invalid type")
		return 0, errors.New("user id not found")
	}
	return idInt, nil

}

func (h *Handler) getUserWallet(c *gin.Context) (string, error) {
	header := c.GetHeader(wallet)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty wallet header")
		return "", errors.New("empty wallet header")
	}

	return header, nil
}

func getId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found")
		return 0, errors.New("id not found")
	}

	return id, nil
}

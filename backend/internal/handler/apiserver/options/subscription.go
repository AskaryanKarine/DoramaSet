package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	errors2 "DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type subResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Duration    string `json:"duration"`
}

func durationToString(t time.Duration) string {
	d := t.Round(time.Minute)
	h := d / time.Hour
	month := (h / 24) / 30
	return fmt.Sprintf("%d month", month)
}

func makeSubResponse(sub model.Subscription) subResponse {
	return subResponse{
		Id:          sub.Id,
		Name:        sub.Name,
		Description: sub.Description,
		Cost:        sub.Cost,
		Duration:    durationToString(sub.Duration),
	}
}

func (h *Handler) getAllSubs(c *gin.Context) {
	data, err := h.Services.GetAll()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response := make([]subResponse, 0)
	for _, el := range data {
		response = append(response, makeSubResponse(el))
	}
	c.JSON(http.StatusOK, gin.H{"Data": response})
}

func (h *Handler) getInfoSub(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.Services.GetInfo(id)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": makeSubResponse(*data)})
}

func (h *Handler) subscribe(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.SubscribeUser(token, id)
	if err != nil && errors.As(err, &errors2.BalanceError{}) {
		_ = c.AbortWithError(http.StatusPaymentRequired, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sub, err := h.Services.GetInfo(id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"Data": makeSubResponse(*sub)})
}

func (h *Handler) unsubscribe(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.UnsubscribeUser(token)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

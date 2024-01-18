package options

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"DoramaSet/internal/handler/apiserver/middleware"
	errors2 "DoramaSet/internal/logic/errors"
	"DoramaSet/internal/tracing"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllSubs(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id/episode")
	defer span.End()
	data, err := h.Services.GetAll(ctx)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response := make([]DTO.Subscription, 0)
	for _, el := range data {
		response = append(response, DTO.MakeSubResponse(el))
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) getInfoSub(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id/episode")
	defer span.End()
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.Services.GetInfo(ctx, id)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeSubResponse(*data)})
}

func (h *Handler) subscribe(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id/episode")
	defer span.End()
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

	err = h.Services.SubscribeUser(ctx, token, id)
	if err != nil && errors.As(err, &errors2.BalanceError{}) {
		_ = c.AbortWithError(http.StatusPaymentRequired, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sub, err := h.Services.GetInfo(ctx, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeSubResponse(*sub)})
}

func (h *Handler) unsubscribe(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id/episode")
	defer span.End()
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.UnsubscribeUser(ctx, token)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/tracing"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) changeColor(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /user/color")
	defer span.End()
	color := c.Query("color")

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.ChangeAvatarColor(ctx, token, color)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) changeEmoji(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /user/emoji")
	defer span.End()
	color := c.Query("emoji")

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.ChangeEmoji(ctx, token, color)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) earnPoint(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /user/earn")
	defer span.End()
	var req struct {
		Points string `json:"points"`
	}
	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	points, err := strconv.Atoi(req.Points)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	user, err := h.Services.AuthByToken(ctx, token)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.Services.EarnPoint(ctx, user, points)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

package options

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createPicture(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /picture/")
	defer span.End()
	var req DTO.Picture

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	newPicture := DTO.MakePicture(req)
	err = h.Services.CreatePicture(ctx, token, newPicture)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": DTO.MakePictureResponse(*newPicture)})
}

func (h *Handler) addPictureToStaff(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /staff/:id")
	defer span.End()
	rowDId := c.Param("id")
	SId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var req struct {
		Id int `json:"id"`
	}

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.AddPictureToStaff(ctx, token, model.Picture{Id: req.Id}, SId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) addPictureToDorama(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /dorama/:id/picture")
	defer span.End()
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req struct {
		Id int `json:"id"`
	}

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.AddPictureToDorama(ctx, token, model.Picture{Id: req.Id}, DId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

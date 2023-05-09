package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createPicture(c *gin.Context) {
	var req model.Picture

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.CreatePicture(token, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": req})
}

func (h *Handler) addPictureToStaff(c *gin.Context) {
	rowDId := c.Param("id")
	SId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	rowPId := c.Query("id")
	PId, err := strconv.Atoi(rowPId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.AddPictureToStaff(token, model.Picture{Id: PId}, SId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) addPictureToDorama(c *gin.Context) {
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rowPId := c.Query("id")
	PId, err := strconv.Atoi(rowPId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.AddPictureToDorama(token, model.Picture{Id: PId}, DId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getEpisodeList(c *gin.Context) {
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.Services.GetEpisodeList(DId)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (h *Handler) createEpisode(c *gin.Context) {
	var req model.Episode
	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.CreateEpisode(token, &req, DId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": req})
}

func (h *Handler) markWatchingEpisode(c *gin.Context) {
	rowId := c.Query("id")
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

	err = h.Services.MarkWatchingEpisode(token, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

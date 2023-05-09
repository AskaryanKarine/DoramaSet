package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getPublicList(c *gin.Context) {
	data, err := h.Services.GetPublicLists()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (h *Handler) getListById(c *gin.Context) {
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

	data, err := h.Services.GetListById(token, id)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (h *Handler) createList(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var req model.List
	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Services.CreateList(token, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": req})
}

func (h *Handler) addToList(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	rowLId := c.Param("id")
	LId, err := strconv.Atoi(rowLId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rowDId := c.Query("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Services.AddToList(token, LId, DId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) delFromList(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	rowLId := c.Param("id")
	LId, err := strconv.Atoi(rowLId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rowDId := c.Query("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Services.DelFromList(token, LId, DId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) delList(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	rowId := c.Query("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Services.DelList(token, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getUserLists(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	data, err := h.Services.GetUserLists(token)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (h *Handler) getUserFavList(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	data, err := h.Services.GetFavList(token)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (h *Handler) addToFav(c *gin.Context) {
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

	err = h.Services.AddToFav(token, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

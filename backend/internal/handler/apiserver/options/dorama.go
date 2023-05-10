package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllDorama(c *gin.Context) {
	dorama, err := h.Services.GetAllDorama()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": dorama})
}

func (h *Handler) createDorama(c *gin.Context) {
	var req model.Dorama

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.CreateDorama(token, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": req})
}

func (h *Handler) updateDorama(c *gin.Context) {
	var req model.Dorama

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.UpdateDorama(token, req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": req})
}

func (h *Handler) getDoramaById(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dorama, err := h.Services.GetDoramaById(id)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": dorama})
}

func (h *Handler) findDoramaByName(c *gin.Context) {
	rowName := c.Query("name")

	dorama, err := h.Services.GetDoramaByName(rowName)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": dorama})
}

func (h *Handler) addStaffToDorama(c *gin.Context) {
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rowSId := c.Query("id")
	SId, err := strconv.Atoi(rowSId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.AddStaffToDorama(token, DId, SId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getStaffListByDorama(c *gin.Context) {
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		return
	}

	data, err := h.Services.GetStaffListByDorama(DId)
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

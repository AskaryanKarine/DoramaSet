package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getStaffList(c *gin.Context) {
	data, err := h.Services.GetStaffList()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (h *Handler) findStaffByName(c *gin.Context) {
	rowName := c.Query("name")

	data, err := h.Services.GetListByName(rowName)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (h *Handler) getStaffById(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	data, err := h.Services.GetStaffById(id)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{"Data": data})

}

func (h *Handler) createStaff(c *gin.Context) {
	var req struct {
		Name     string `json:"name,omitempty"`
		Birthday string `json:"birthday,omitempty"`
		Type     string `json:"type,omitempty"`
		Gender   string `json:"gender,omitempty"`
	}

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	t, err := time.Parse("02/01/2006", req.Birthday)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newStaff := model.Staff{
		Name:     req.Name,
		Birthday: t,
		Type:     req.Type,
		Gender:   req.Gender,
		Photo:    nil,
	}

	err = h.Services.CreateStaff(token, &newStaff)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": newStaff})
}

func (h *Handler) updateStaff(c *gin.Context) {
	var req struct {
		Id       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Birthday string `json:"birthday,omitempty"`
		Type     string `json:"type,omitempty"`
		Gender   string `json:"gender,omitempty"`
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

	t, err := time.Parse("02/01/2006", req.Birthday)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newStaff := model.Staff{
		Id:       req.Id,
		Name:     req.Name,
		Birthday: t,
		Type:     req.Type,
		Gender:   req.Gender,
		Photo:    nil,
	}

	err = h.Services.UpdateStaff(token, newStaff)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": newStaff})
}

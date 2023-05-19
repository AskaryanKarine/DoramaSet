package options

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"DoramaSet/internal/handler/apiserver/middleware"
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
	var res []DTO.Staff
	for _, d := range data {
		res = append(res, DTO.MakeStaffResponse(d))
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
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
		return
	}
	var res []DTO.Staff
	for _, d := range data {
		res = append(res, DTO.MakeStaffResponse(d))
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeStaffResponse(*data)})

}

func (h *Handler) createStaff(c *gin.Context) {
	var req DTO.Staff

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	t, err := time.Parse("02.01.2006", req.Birthday)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newStaff := DTO.MakeStaff(req, t)

	err = h.Services.CreateStaff(token, newStaff)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeStaffResponse(*newStaff)})
}

func (h *Handler) updateStaff(c *gin.Context) {
	var req DTO.Staff

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	t, err := time.Parse("02.01.2006", req.Birthday)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	staff := DTO.MakeStaff(req, t)
	staff.Id = req.Id

	err = h.Services.UpdateStaff(token, *staff)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeStaffResponse(*staff)})
}

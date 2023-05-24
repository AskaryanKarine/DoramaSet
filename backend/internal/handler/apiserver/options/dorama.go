package options

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"DoramaSet/internal/handler/apiserver/middleware"
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

	var response []DTO.Dorama
	for _, d := range dorama {
		var review []DTO.Review
		for _, r := range d.Reviews {
			info, err := h.Services.GetPublicInfo(r.Username)
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			review = append(review, DTO.MakeReviewResponse(r, *info))
		}
		res := DTO.MakeDoramaResponse(d)
		res.Reviews = review
		response = append(response, res)
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) createDorama(c *gin.Context) {
	var req DTO.Dorama

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	newDorama := DTO.MakeDorama(req)
	err = h.Services.CreateDorama(token, newDorama)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeDoramaResponse(*newDorama)})
}

func (h *Handler) updateDorama(c *gin.Context) {
	var req DTO.Dorama

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	data := DTO.MakeDorama(req)
	data.Id = req.Id
	err = h.Services.UpdateDorama(token, *data)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": req})
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

	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeDoramaResponse(*dorama)})
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

	var response []DTO.Dorama
	for _, d := range dorama {
		response = append(response, DTO.MakeDoramaResponse(d))
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) addStaffToDorama(c *gin.Context) {
	var req DTO.Id

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

	err = h.Services.AddStaffToDorama(token, DId, req.Id)
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

	var response []DTO.Staff
	for _, d := range data {
		response = append(response, DTO.MakeStaffResponse(d))
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) CreateReview(c *gin.Context) {
	var req DTO.Review
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		return
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

	newReview := DTO.MakeReview(req)
	err = h.Services.AddReview(token, DId, newReview)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info, err := h.Services.GetPublicInfo(req.Username)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := DTO.MakeReviewResponse(*newReview, *info)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) DeleteReview(c *gin.Context) {
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.DeleteReview(token, DId)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

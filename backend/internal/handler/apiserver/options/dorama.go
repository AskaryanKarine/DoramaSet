package options

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/tracing"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllDorama(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/")
	defer span.End()
	dorama, err := h.Services.GetAllDorama(ctx)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var response []DTO.Dorama
	for _, d := range dorama {
		var review []DTO.Review
		for _, r := range d.Reviews {
			info, err := h.Services.GetPublicInfo(ctx, r.Username)
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
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /dorama/")
	defer span.End()
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
	err = h.Services.CreateDorama(ctx, token, newDorama)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeDoramaResponse(*newDorama)})
}

func (h *Handler) updateDorama(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "PUT /dorama/")
	defer span.End()
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
	err = h.Services.UpdateDorama(ctx, token, *data)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *Handler) getDoramaById(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id")
	defer span.End()
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dorama, err := h.Services.GetDoramaById(ctx, id)
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
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /find/dorama/")
	defer span.End()
	rowName := c.Query("name")

	dorama, err := h.Services.GetDoramaByName(ctx, rowName)
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
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /dorama/:id/staff")
	defer span.End()
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

	err = h.Services.AddStaffToDorama(ctx, token, DId, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getStaffListByDorama(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id/staff")
	defer span.End()
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		return
	}

	data, err := h.Services.GetStaffListByDorama(ctx, DId)
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
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /dorama/:id/review")
	defer span.End()
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
	err = h.Services.AddReview(ctx, token, DId, newReview)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info, err := h.Services.GetPublicInfo(ctx, req.Username)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := DTO.MakeReviewResponse(*newReview, *info)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) DeleteReview(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "DELETE /dorama/:id/review")
	defer span.End()
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

	err = h.Services.DeleteReview(ctx, token, DId)
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

package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type listResponse struct {
	Id          int            `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	CreatorName string         `json:"creator_name,omitempty"`
	Type        string         `json:"type,omitempty"`
	Doramas     []model.Dorama `json:"doramas,omitempty"`
}

func makeListResponse(list model.List) listResponse {
	str, _ := constant.GetTypeList(list.Type)
	return listResponse{
		Id:          list.Id,
		Name:        list.Name,
		Description: list.Description,
		CreatorName: list.CreatorName,
		Type:        str,
		Doramas:     list.Doramas,
	}
}

func (h *Handler) getPublicList(c *gin.Context) {
	data, err := h.Services.GetPublicLists()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var result []listResponse
	for _, d := range data {
		result = append(result, makeListResponse(d))
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": makeListResponse(*data)})
}

func (h *Handler) createList(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var req listResponse
	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	model := model.List{
		Name:        req.Name,
		Description: req.Description,
		Type:        constant.ListType[req.Type],
		Doramas:     nil,
	}

	err = h.Services.CreateList(token, &model)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": model})
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

	var req struct {
		Id int `json:"id"`
	}

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Services.AddToList(token, LId, req.Id)
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

	var result []listResponse
	for _, d := range data {
		result = append(result, makeListResponse(d))
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
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

	var result []listResponse
	for _, d := range data {
		result = append(result, makeListResponse(d))
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
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

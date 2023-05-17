package options

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"DoramaSet/internal/handler/apiserver/middleware"
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

	var response []DTO.Episode
	for _, d := range data {
		response = append(response, DTO.MakeEpisodeRequest(d))
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) createEpisode(c *gin.Context) {
	var req DTO.Episode

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

	newEpisode := DTO.MakeEpisode(req)
	err = h.Services.CreateEpisode(token, newEpisode, DId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeEpisodeRequest(*newEpisode)})
}

func (h *Handler) markWatchingEpisode(c *gin.Context) {
	var req DTO.Id

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.MarkWatchingEpisode(token, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getEpisodeWithStatus(c *gin.Context) {
	var response []DTO.WatchingResponse

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

	watching, err := h.Services.GetWatchingEpisode(token, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	all, err := h.GetEpisodeList(id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(token) == 0 {
		for _, d := range all {
			response = append(response, DTO.MakeWatchingResponse(d, false))
		}
		c.JSON(http.StatusOK, gin.H{"data": response})
	}

	for _, e := range all {
		w := false
		for i, watch := range watching {
			if e.Id == watch.Id {
				w = true
				watching = append(watching[:i], watching[i+1:]...)
				break
			}
		}
		response = append(response, DTO.MakeWatchingResponse(e, w))

	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

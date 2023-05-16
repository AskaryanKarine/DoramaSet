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
	var req struct {
		Id int `json:"id,omitempty"`
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

	err = h.Services.MarkWatchingEpisode(token, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type watchingResponse struct {
	Episode  model.Episode `json:"episode"`
	Watching bool          `json:"watching"`
}

func (h *Handler) getEpisodeWithStatus(c *gin.Context) {
	var response []watchingResponse
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

	if len(token) == 0 {
		c.JSON(http.StatusOK, gin.H{"Data": response})
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

	for _, e := range all {
		w := false
		for i, watch := range watching {
			if e.Id == watch.Id {
				w = true
				watching = append(watching[:i], watching[i+1:]...)
				break
			}
		}
		response = append(response, watchingResponse{Episode: e, Watching: w})

	}

	c.JSON(http.StatusOK, gin.H{"Data": response})
}

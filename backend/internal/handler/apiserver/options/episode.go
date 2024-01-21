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

func (h *Handler) getEpisodeList(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id/episode")
	defer span.End()
	rowDId := c.Param("id")
	DId, err := strconv.Atoi(rowDId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.Services.GetEpisodeList(ctx, DId)
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
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /dorama/:id/episode")
	defer span.End()
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
	err = h.Services.CreateEpisode(ctx, token, newEpisode, DId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DTO.MakeEpisodeRequest(*newEpisode)})
}

func (h *Handler) markWatchingEpisode(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /user/episode")
	defer span.End()
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

	err = h.Services.MarkWatchingEpisode(ctx, token, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getEpisodeWithStatus(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /user/episode")
	defer span.End()

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

	all, err := h.GetEpisodeList(ctx, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(token) == 0 {
		for _, d := range all {
			response = append(response, DTO.MakeWatchingResponse(d, false))
		}
		c.JSON(http.StatusOK, gin.H{"data": response})
		return
	}

	watching, err := h.Services.GetWatchingEpisode(ctx, token, id)
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
		response = append(response, DTO.MakeWatchingResponse(e, w))

	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

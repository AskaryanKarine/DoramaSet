package options

import (
	"DoramaSet/internal/logic/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userResponse struct {
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Points   int         `json:"points,omitempty"`
	IsAdmin  bool        `json:"isAdmin,omitempty"`
	Sub      subResponse `json:"subs"`
	LastSubs string      `json:"lastSubs,omitempty"`
	RegData  string      `json:"regData,omitempty"`
}

func makeUserResponse(user model.User) userResponse {
	return userResponse{
		Username: user.Username,
		Email:    user.Email,
		Points:   user.Points,
		IsAdmin:  user.IsAdmin,
		Sub:      makeSubResponse(*user.Sub),
		LastSubs: user.LastSubscribe.String(),
		RegData:  user.RegData.String(),
	}
}

func (h *Handler) login(c *gin.Context) {
	var req model.User

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.Services.Login(req.Username, req.Password)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.Services.AuthByToken(token)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  makeUserResponse(*user),
	})
}

func (h *Handler) registration(c *gin.Context) {
	var req model.User

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.Services.Registration(&req)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  makeUserResponse(req),
	})
}

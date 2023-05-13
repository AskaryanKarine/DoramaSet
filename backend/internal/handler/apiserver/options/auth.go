package options

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type userResponse struct {
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Points   int         `json:"points,omitempty"`
	IsAdmin  bool        `json:"isAdmin,omitempty"`
	Sub      subResponse `json:"sub"`
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
		LastSubs: user.LastSubscribe.Format("1 January 2006"),
		RegData:  user.RegData.Format("1 January 2006"),
	}
}

func setCookie(c *gin.Context, token string, tokenExp int) {

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Duration(tokenExp) * time.Hour),
		HttpOnly: false,
	})
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

	// c.SetCookie("token", token, 3600*h.tokenExprHour, "/", "localhost", false, true)
	setCookie(c, token, h.tokenExprHour)
	c.JSON(http.StatusOK, gin.H{
		"user": makeUserResponse(*user),
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

	// c.SetCookie("token", token, 3600*h.tokenExprHour, "/", "localhost", false, true)
	setCookie(c, token, h.tokenExprHour)
	c.JSON(http.StatusOK, gin.H{
		"user": makeUserResponse(req),
	})
}

func (h *Handler) getUserByCookieToken(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	user, err := h.Services.AuthByToken(cookie)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": makeUserResponse(*user)})
}

func (h *Handler) logout(c *gin.Context) {
	// c.SetCookie("token", "", -1, "/", "localhost", false, true)
	setCookie(c, "", -1)
	c.JSON(http.StatusOK, gin.H{})
}

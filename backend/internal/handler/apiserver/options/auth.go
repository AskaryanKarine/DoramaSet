package options

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func setCookie(c *gin.Context, token string, tokenExp int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Duration(tokenExp) * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
}

func (h *Handler) login(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /auth/login")
	defer span.End()
	var req DTO.Auth

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.Services.Login(ctx, req.Login, req.Password)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.Services.AuthByToken(ctx, token)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	setCookie(c, token, h.tokenExprHour)
	c.JSON(http.StatusOK, gin.H{"user": DTO.MakeUserResponse(*user)})
}

func (h *Handler) registration(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "POST /auth/registration")
	defer span.End()
	var req DTO.Auth

	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newUser := DTO.MakeUser(req)
	token, err := h.Services.Registration(ctx, newUser)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	setCookie(c, token, h.tokenExprHour)
	c.JSON(http.StatusOK, gin.H{"user": DTO.MakeUserResponse(*newUser)})
}

func (h *Handler) getUserByCookieToken(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /auth/")
	defer span.End()
	cookie, err := c.Cookie("token")
	if err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	user, err := h.Services.AuthByToken(ctx, cookie)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": DTO.MakeUserResponse(*user)})
}

func (h *Handler) logout(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /auth/logout")
	defer span.End()
	setCookie(c, "", -1)
	c.JSON(http.StatusOK, gin.H{})
}

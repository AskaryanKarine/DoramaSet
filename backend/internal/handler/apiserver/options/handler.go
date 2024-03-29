package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/handler/apiserver/services"
	"DoramaSet/internal/tracing"
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	services.Services
	mode          string
	tokenExprHour int
}

func NewHandler(services services.Services, mode string, tokenExprHour int) *Handler {
	return &Handler{
		Services:      services,
		mode:          mode,
		tokenExprHour: tokenExprHour,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(h.mode)
	router := gin.Default()

	prometheus.MustRegister(requestDuration, requestsTotal)
	router.Use(prometheusMiddleware)
	router.Use(middleware.ErrorHandler)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
			"withCredentials", "Set-Cookie", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	home := router.Group("/", h.updateUserDataByToken)
	{
		auth := home.Group("/auth")
		{
			auth.GET("/", h.getUserByCookieToken)
			auth.POST("/registration", h.registration) // guest
			auth.POST("/login", h.login)               // guest
			auth.GET("/logout", h.logout)
		}

		user := home.Group("/user")
		{
			user.GET("/list", h.getUserLists)       // user
			user.POST("/favorite", h.addToFav)      // user
			user.GET("/favorite", h.getUserFavList) // user
			user.GET("/color", h.changeColor)
			user.GET("/emoji", h.changeEmoji)
			user.POST("/earn", h.earnPoint)
			user.GET("/episode", h.getEpisodeWithStatus)
			user.POST("/episode", h.markWatchingEpisode) // user
		}

		subscription := home.Group("/subscription")
		{
			subscription.GET("/", h.getAllSubs)    // guest
			subscription.GET("/:id", h.getInfoSub) // guest
			subscription.POST("/:id", h.subscribe) // user
			subscription.POST("/", h.unsubscribe)  // user
		}

		dorama := home.Group("/dorama")
		{
			dorama.GET("/", h.getAllDorama)                   // guest
			dorama.GET("/:id", h.getDoramaById)               // guest
			dorama.POST("/", h.createDorama)                  // admin
			dorama.PUT("/", h.updateDorama)                   // admin
			dorama.POST("/:id/staff", h.addStaffToDorama)     // admin
			dorama.GET("/:id/staff", h.getStaffListByDorama)  // guest
			dorama.GET("/:id/episode", h.getEpisodeList)      // guest
			dorama.POST("/:id/episode", h.createEpisode)      // admin
			dorama.POST("/:id/picture", h.addPictureToDorama) // admin
			dorama.POST("/:id/review", h.CreateReview)
			dorama.DELETE("/:id/review", h.DeleteReview)
		}

		list := home.Group("/list")
		{
			list.GET("/public", h.getPublicList) // guest
			list.GET("/:id", h.getListById)      // guest
			list.POST("/", h.createList)         // user
			list.POST("/:id", h.addToList)       // user
			list.DELETE("/:id", h.delFromList)   // user
		}

		staff := home.Group("/staff")
		{
			staff.GET("/", h.getStaffList)                  // guest
			staff.POST("/", h.createStaff)                  // admin
			staff.PUT("/", h.updateStaff)                   // admin
			staff.GET("/:id", h.getStaffById)               // guest
			staff.POST("/:id/picture", h.addPictureToStaff) // admin

		}

		picture := home.Group("/picture")
		{
			picture.POST("/", h.createPicture) // admin
		}

		find := home.Group("/find")
		{
			find.GET("/dorama/", h.findDoramaByName) // guest
			find.GET("/staff/", h.findStaffByName)   // guest
		}
	}

	router.GET("/metrics", prometheusHandler())

	return router
}

func (h *Handler) updateUserDataByToken(c *gin.Context) {
	ctx := context.Background()
	ctx, span := tracing.StartSpanFromContext(ctx, "GET /dorama/:id/episode")
	defer span.End()
	var token string
	cook, err := c.Cookie("token")
	if errors.Is(err, http.ErrNoCookie) {
		token = ""
	} else if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid cookie"))
		return
	} else {
		token = cook

		err := h.Services.UpdateActive(ctx, token)
		if err != nil {
			setCookie(c, "", -1)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		err = h.Services.UpdateSubscribe(ctx, token)
		if err != nil {
			setCookie(c, "", -1)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}

	c.Set("userToken", cook)
}

func fatalDB(e error) bool {
	return strings.Contains(e.Error(), "connect")
}

package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/handler/apiserver/services"
	"DoramaSet/internal/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Handler struct {
	log logger.Logger
	services.Services
}

func NewHandler(log logger.Logger, services services.Services) *Handler {
	return &Handler{log, services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(middleware.ErrorHandler)

	home := router.Group("/", h.updateUserActiveByToken)
	{
		auth := home.Group("/auth")
		{
			auth.POST("/registration", h.registration) // guest
			auth.GET("/login", h.login)                // guest
		}

		user := home.Group("/user")
		{
			user.POST("/episode", h.markWatchingEpisode) // user
			user.GET("/list", h.getUserLists)            // user
			user.POST("/favorite/", h.addToFav)          // user
			user.GET("/favorite", h.getUserFavList)      // user
		}

		subscription := home.Group("/subscription")
		{
			subscription.GET("/", h.getAllSubs)    // guest
			subscription.GET("/:id", h.getInfoSub) // guest
			subscription.PUT("/:id", h.subscribe)  // user
			subscription.PUT("/", h.unsubscribe)   // user
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
		}

		list := home.Group("/list")
		{
			list.GET("/public", h.getPublicList) // guest
			list.GET("/:id", h.getListById)      // guest
			list.POST("/", h.createList)         // user
			list.POST("/:id", h.addToList)       // user
			list.DELETE("/:id", h.delFromList)   // user
			list.DELETE("/", h.delList)          // user
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

	return router
}

func (h *Handler) updateUserActiveByToken(c *gin.Context) {
	var token string
	header := c.GetHeader("Authorization")
	if header == "" {
		token = ""
	} else {
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			_ = c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid auth header"))
			return
		}

		if len(headerParts[1]) == 0 {
			_ = c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("token is empty"))
			return
		}
		token = headerParts[1]
		err := h.Services.UpdateActive(token)
		if err != nil && fatalDB(err) {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
		}
	}

	c.Set("userToken", token)
}

func fatalDB(e error) bool {
	return strings.Contains(e.Error(), "connect")
}

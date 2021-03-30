package rest

import (
	"inflix-main/controller"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SetupRouter - sets the REST routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	v1 := r.Group("/v1")
	{
		v1.GET("album", controller.GetAlbums)
		v1.POST("album", controller.CreateAlbum)
		v1.GET("album/:id", controller.GetAlbum)
		v1.PUT("album/:id", controller.UpdateAlbum)
		v1.DELETE("album/:id", controller.DeleteAlbum)
	}

	return r
}

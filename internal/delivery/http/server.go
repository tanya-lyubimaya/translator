package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tanya-lyubimaya/translator/internal/domain"
	"net/http"
	"strings"
)

var (
	origins = map[string]bool{
		"http://localhost":      true,
		"http://localhost:8080": true,
		"http://localhost:8081": true,
	}
)

type Server struct {
	router *gin.Engine
	uc     domain.UseCase
}

func New(uc domain.UseCase) (*Server, error) {
	router := gin.Default()
	router.Use(CORSMiddleware())
	server := &Server{uc: uc, router: router}

	router.POST("/get-translation", server.GetTranslation)
	router.GET("/get-languages", server.GetLanguages)
	return server, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for v := range origins {
			if strings.Index(v, c.Request.Referer()) > -1 {
				if c.Request.Referer()[len(c.Request.Referer())-1:] == "/" {
					c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Referer()[0:len(c.Request.Referer())-1])
				} else {
					c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Referer())
				}
				break
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) GetTranslation(c *gin.Context) {
	var model domain.TranslateRequestModel
	err := c.ShouldBindJSON(&model)
	output, err := s.uc.GetTranslation(c, model)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, output)
}

func (s *Server) GetLanguages(c *gin.Context) {
	l, err := s.uc.GetLanguages(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, l)
}

func (s *Server) Serve(port string) error {
	return s.router.Run(port)
}

func (s *Server) Stop() error {
	// TODO: stop server on graceful interrupt
	//s.router.
	return nil
}

package api

import (
	"math/rand"
	"net/http"

	"example.com/cipher"
	"example.com/db"

	"github.com/gin-gonic/gin"
)

type querySearch struct {
	Link     string `form:"link" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type queryValue struct {
	Message  string `form:"message" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Server struct {
	listenAddr     string
	serverClient   *gin.Engine
	databaseClient *db.RedisClient
}

func NewServer(listenAddr string, databaseClient *db.RedisClient) *Server {
	router := gin.Default()

	return &Server{
		listenAddr:     listenAddr,
		serverClient:   router,
		databaseClient: databaseClient,
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) Start() {
	s.serverClient.Use(corsMiddleware())

	s.serverClient.POST("read-db", s.handleGetQuery)
	s.serverClient.POST("add-value", s.handleAddValue)

	s.serverClient.Run(s.listenAddr)
}

func (s *Server) handleGetQuery(c *gin.Context) {
	var newQuery querySearch

	err := c.ShouldBind(&newQuery)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, example: {link: <link>, password: <password>}",
		})
		return
	}

	if len(newQuery.Password) > 32 {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Too long passowrd.",
		})
		return
	}

	hexEncrypted, err := s.databaseClient.GetKeyValue(newQuery.Link)
	if err != nil {
		c.SecureJSON(http.StatusNotFound, gin.H{
			"message": "Couldn't find link.",
		})
		return
	}

	plaintext, err := cipher.DecryptString(hexEncrypted, newQuery.Password)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Decryption failed.",
		})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{
		"message": plaintext,
	})
}

func randomLink(size int) string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, size)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func (s *Server) handleAddValue(c *gin.Context) {
	var newValue queryValue

	err := c.ShouldBind(&newValue)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"link": "Bad Request, example: {message: <message>, password: <password}",
		})
		return
	}

	if len(newValue.Password) > 32 {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"link": "Too long passowrd.",
		})
		return
	}

	hexEncrypted, err := cipher.EncryptString(newValue.Message, newValue.Password)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"link": "Encryption failed.",
		})
		return
	}

	link := randomLink(8)
	err = s.databaseClient.AddKeyValue(link, hexEncrypted)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"link": "Couldn't add message into database",
		})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{
		"link": link,
	})
}

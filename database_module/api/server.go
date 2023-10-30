package api

import (
	"math/rand"
	"net/http"

	"example.com/db"
	"github.com/gin-gonic/gin"
)

type querySearch struct {
	Link string `json:"link" binding:"required"`
}

type queryValue struct {
	Message string `json:"message" binding:"required"`
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

func (s *Server) Start() {
	s.serverClient.POST("read-db", s.handleGetQuery)
	s.serverClient.POST("add-value", s.handleAddValue)
	s.serverClient.Run(s.listenAddr)
}

func (s *Server) handleGetQuery(c *gin.Context) {
	var newQuery querySearch

	err := c.ShouldBindJSON(&newQuery)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request, example: {link: <link>}",
		})
		return
	}

	val, err := s.databaseClient.GetKeyValue(newQuery.Link)
	if err != nil {
		c.SecureJSON(http.StatusNotFound, gin.H{
			"error": "Couldn't find key value",
		})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{
		"message": val,
	})
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	random := make([]rune, n)

	for i := range random {
		random[i] = letters[rand.Intn(len(letters))]
	}
	return string(random)
}

func (s *Server) handleAddValue(c *gin.Context) {
	var newValue queryValue

	err := c.ShouldBindJSON(&newValue)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request, example: {message: <message>}",
		})
		return
	}

	randomLink := randomString(8)
	err = s.databaseClient.AddKeyValue(randomLink, newValue.Message)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't add message into database",
		})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{
		"link": randomLink,
	})
}

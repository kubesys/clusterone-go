/**
 * Copyright (2024, ) Institute of Software, Chinese Academy of Sciences
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * author: wuheng@iscas.ac.cn
 * since:  0.1
 */

package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	router *gin.Engine
}

func NewApiServer() *ApiServer {
	server := &ApiServer{
		router: gin.Default(),
	}
	server.setupRoutes()
	return server
}

func (s *ApiServer) setupRoutes() {
	apiGroup := s.router.Group("/api")
	apiGroup.Use(s.authenticateMiddleware()) // Bearer Token authentication middleware

	apiGroup.POST("/resource", s.Post)
	apiGroup.GET("/resource/:id", s.Get)
	apiGroup.PUT("/resource/:id", s.Put)
	apiGroup.DELETE("/resource/:id", s.Delete)
}

func (s *ApiServer) authenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform Bearer Token authentication logic here
		token := c.GetHeader("Authorization")
		// Validate and extract user information from the token

		// Example validation (you should replace this with your actual validation logic)
		if token != "Bearer valid_token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Authentication successful, proceed to the next middleware/handler
		c.Next()
	}
}

func (s *ApiServer) Post(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle POST logic with the JSON data

	c.JSON(http.StatusCreated, gin.H{"message": "Resource created successfully"})
}

func (s *ApiServer) Get(c *gin.Context) {
	// Extract resource ID from the URL parameters
	resourceID := c.Param("id")

	// Handle GET logic with the resource ID

	c.JSON(http.StatusOK, gin.H{"message": "Resource retrieved successfully", "id": resourceID})
}

func (s *ApiServer) Put(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract resource ID from the URL parameters
	resourceID := c.Param("id")

	// Handle PUT logic with the resource ID and JSON data

	c.JSON(http.StatusOK, gin.H{"message": "Resource updated successfully", "id": resourceID})
}

func (s *ApiServer) Delete(c *gin.Context) {
	// Extract resource ID from the URL parameters
	resourceID := c.Param("id")

	// Handle DELETE logic with the resource ID

	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully", "id": resourceID})
}

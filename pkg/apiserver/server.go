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
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	Router *gin.Engine
}

func NewApiServer() *ApiServer {
	server := &ApiServer{
		Router: gin.Default(),
	}
	server.setupRoutes()
	return server
}

func (s *ApiServer) setupRoutes() {
	apiGroup := s.Router.Group("/")
	//apiGroup.Use(s.authenticateMiddleware()) // Bearer Token authentication middleware
	// reg_exp = "/(api|apis)/([a-z0-9._-]{2,100})/([a-z0-9._-]{2,100})/namespaces/([a-z0-9._-]{2,100})/([a-z0-9._-]{2,100})"
	{
		apiGroup.POST("/*any", PostHandler)
		apiGroup.GET("/*any", GetHandler)
		apiGroup.PUT("/*any", PutHandler)
		apiGroup.DELETE("/*any", DeleteHandler)
	}
}

//func (s *ApiServer) authenticateMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Perform Bearer Token authentication logic here
//		token := c.GetHeader("Authorization")
//		// Validate and extract user information from the token
//
//		// Example validation (you should replace this with your actual validation logic)
//		if token != "Bearer valid_token" {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//
//		// Authentication successful, proceed to the next middleware/handler
//		c.Next()
//	}
//}

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
	"net/http"
	"regexp"
)

const post_list_regexp = `/(api|apis)/([a-z0-9._-]{2,100})/([a-z0-9._-]{2,100})/namespaces/([a-z0-9._-]{2,100})/([a-z0-9._-]{2,100})`
const get_put_delete_watch_regexp = `/(api|apis)/([a-z0-9._-]{2,100})/([a-z0-9._-]{2,100})/namespaces/([a-z0-9._-]{2,100})/([a-z0-9._-]{2,100})`

func GetHandler(c *gin.Context) {
	if valid(get_put_delete_watch_regexp, c.Request.RequestURI) == true {
		c.JSON(http.StatusOK, gin.H{
			"message": c.Request.RequestURI,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "the format is " + get_put_delete_watch_regexp,
		})
	}
}

func PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST request handled",
	})
}

func PutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT request handled",
	})
}

func DeleteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE request handled",
	})
}

func valid(pattern string, url string) bool {
	re := regexp.MustCompile(pattern)

	if re.MatchString(url) {
		return true
	}

	return false

}

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

package main

import (
	"clusterone-go/pkg/apiserver"
	"fmt"
)

func main() {
	startServer()
	//testPg()
}

func testPg() {
	pgOrm, err := apiserver.NewPostgresOrm("postgresql://postgres:password@ip:port/db")
	if err != nil {
		fmt.Println("Error creating PostgreSQL ORM:", err)
		return
	}

	pgOrm.CreateTable("test")
}

func startServer() {
	server := apiserver.NewApiServer()
	server.Router.Run(":8080")
}

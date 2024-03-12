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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOrm struct {
	Db *gorm.DB
}

// NewPostgresOrm postgresql://username:password@ip:port/db
func NewPostgresOrm(connection string) (*PostgresOrm, error) {
	gormDB, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresOrm{Db: gormDB}, nil
}

func (p *PostgresOrm) CreateTable(table string) error {
	createTableSQL := "CREATE TABLE IF NOT EXISTS " + table + " " +
		"(id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, age INT)"
	return p.Db.Exec(createTableSQL).Error
}

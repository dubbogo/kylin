/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

import (
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

import (
	"github.com/arana-db/arana/pkg/util/log"
)

var (
	ctx context.Context
)

func SetupMySQLContainer(password, database string) testcontainers.Container {
	log.Info("Setup MySQL Container")
	ctx = context.Background()

	seedDataPath, err := os.Getwd()
	if err != nil {
		log.Errorf("Error get working directory: %s", err)
		panic(fmt.Sprintf("%v", err))
	}

	mountPath := seedDataPath + "/../docker/scripts"

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": password,
			"MYSQL_DATABASE":      database,
		},
		BindMounts: map[string]string{
			"/docker-entrypoint-initdb.d": mountPath,
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		log.Errorf("Error Start MySQL container: %s", err)
		panic(fmt.Sprintf("%v", err))
	}
	return container
}

func CloseContainer(container testcontainers.Container) {
	log.Info("Start terminating MySQL container")
	err := container.Terminate(ctx)
	if err != nil {
		log.Errorf("Terminating MySQL container error: %s", err)
		panic(fmt.Sprintf("%v", err))
	}
}

func OpenDBConnection(username, password, database string, container testcontainers.Container) (*sql.DB, error) {
	host, _ := container.Host(ctx)
	p, _ := container.MappedPort(ctx, "3306/tcp")
	port := p.Int()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		username, password, host, port, database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Info("Error connect to MySQL: %+v\n", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Infof("Error ping to MySQL: %+v\n", err)
		return nil, err
	}

	return db, nil
}

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"encoding/json"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/uber/jaeger/examples/hotrod/pkg/log"
)

var (
	// 'frontend' service

	// WorkerPoolSize is the size of goroutine pool used to query for routes
	WorkerPoolSize = 3

	//FrontendFileLocation
	FrontendFileLocation = "services/frontend/web_assets/index.html"

	// 'customer' service

	// MySQLGetDelay is how long retrieving a customer record takes.
	// Using large value mostly because I cannot click the button fast enough to cause a queue.
	MySQLGetDelay = 300 * time.Millisecond

	// MySQLGetDelayStdDev is standard deviation
	MySQLGetDelayStdDev = MySQLGetDelay / 10

	// RouteWorkerPoolSize is the size of the worker pool used to query `route` service
	RouteWorkerPoolSize = 3

	// CustomerAPIURL defines API URL for our customer microservice
	CustomerAPIURL = "http://127.0.0.1:8081"

	// 'driver' service

	// RedisFindDelay is how long finding closest drivers takes
	RedisFindDelay = 20 * time.Millisecond

	// RedisFindDelayStdDev is standard deviation
	RedisFindDelayStdDev = RedisFindDelay / 4

	// RedisGetDelay is how long retrieving a driver record takes
	RedisGetDelay = 10 * time.Millisecond

	// RedisGetDelayStdDev is standard deEnvWorkerPoolSizeviation
	RedisGetDelayStdDev = RedisGetDelay / 4

	// DriverAPIURL defines API URL for drivers microservice
	DriverAPIURL = "127.0.0.1:8082"

	// 'route' service

	// RouteAPIURL defines API URL for routes microservice
	RouteAPIURL = "http://127.0.0.1:8083"

	// RouteCalcDelay is how long a route calculation takes
	RouteCalcDelay = 50 * time.Millisecond

	// RouteCalcDelayStdDev is standard deviation
	RouteCalcDelayStdDev = RouteCalcDelay / 4
)

type Configuration struct {
	WorkerPoolSize       int           `envconfig:"WORKER_POOL_SIZE"`
	MySQLGetDelay        time.Duration `envconfig:"MYSQL_GET_DELAY"`
	MySQLGetDelayStdDev  time.Duration `envconfig:"MYSQL_GET_DELAY_STD_DEV"`
	RouteWorkerPoolSize  int           `envconfig:"ROUTE_WORKER_POOL_SIZE"`
	RedisFindDelay       time.Duration `envconfig:"REDIS_FIND_DELAY"`
	RedisFindDelayStdDev time.Duration `envconfig:"REDIS_FIND_DELAY_STD_DEV"`
	RedisGetDelay        time.Duration `envconfig:"REDIS_GET_DELAY"`
	RedisGetDelayStdDev  time.Duration `envconfig:"REDIS_GET_DELAY_STD_DEV"`
	RouteCalcDelay       time.Duration `envconfig:"ROUTE_CALC_DELAY"`
	RouteCalcDelayStdDev time.Duration `envconfig:"ROUTE_CALC_DELAY_STD_DEV"`
	FrontendFileLocation string        `envconfig:"FRONTEND_FILE_LOCATION"`
	CustomerAPIURL       string        `envconfig:"CUSTOMER_API_LOCATION"`
	DriverAPIURL         string        `envconfig:"DRIVER_API_LOCATION"`
	RouteAPIURL          string        `envconfig:"ROUTE_API_LOCATION"`
}

// GetConfig sets default config for the application
func GetConfig(log log.Factory) Configuration {
	// TODO: Add cli flags for those values too
	config := Configuration{
		WorkerPoolSize:       WorkerPoolSize,
		MySQLGetDelay:        MySQLGetDelay,
		MySQLGetDelayStdDev:  MySQLGetDelayStdDev,
		RouteWorkerPoolSize:  RouteWorkerPoolSize,
		RedisFindDelay:       RedisFindDelay,
		RedisFindDelayStdDev: RedisFindDelayStdDev,
		RedisGetDelay:        RedisGetDelay,
		RedisGetDelayStdDev:  RedisGetDelayStdDev,
		RouteCalcDelay:       RouteCalcDelay,
		RouteCalcDelayStdDev: RedisFindDelayStdDev,
		FrontendFileLocation: FrontendFileLocation,
		CustomerAPIURL:       CustomerAPIURL,
		DriverAPIURL:         DriverAPIURL,
		RouteAPIURL:          RouteAPIURL,
	}
	err := envconfig.Process("hotrod", &config)
	if err != nil {
		log.Bg().Error(err.Error())
	}
	out, _ := json.Marshal(config)
	log.Bg().Info(string(out))

	return config
}

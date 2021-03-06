// Copyright 2018 Axel Springer SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cfg

import (
	"net/http"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// Config exports a reusable config
var Config *CmdConfig

const (
	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = log.WarnLevel

	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP

	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT

	// DefaultVerbose is the default verbosity.
	DefaultVerbose = false

	// DefaultTimeout is the default time to configure the runtime
	DefaultTimeout = 60

	// DefaultHTTPMethod is the default method to be used for url
	DefaultHTTPMethod = http.MethodGet

	// DefaultHTTPURL is the default url to be checked
	DefaultHTTPURL = "http://localhost:8080"
)

func init() {
	Config = &CmdConfig{
		Verbose:      DefaultVerbose,
		LogLevel:     DefaultLogLevel,
		ReloadSignal: DefaultReloadSignal,
		KillSignal:   DefaultKillSignal,
		Timeout:      DefaultTimeout,
		HTTPMethod:   DefaultHTTPMethod,
		HTTPURL:      DefaultHTTPURL,
	}
}

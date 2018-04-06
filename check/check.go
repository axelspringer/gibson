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

package check

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	config "github.com/axelspringer/gibson/cfg"
	log "github.com/sirupsen/logrus"
)

// Cmd exports the check command
var Cmd *cobra.Command

// config
var cfg = config.Config

// error
var (
	errNoURL     = errors.New("no valid url provided")
	errUnhealthy = errors.New("unhealthy")
)

// New creates a new command line interface.
func New() *Check {
	return &Check{}
}

// exports command by default
func init() {
	Cmd = &cobra.Command{
		Use:   "check",
		Short: "Checks health of an url or process",
		RunE:  runE,
	}

	// URL
	Cmd.Flags().StringVar(&cfg.URL, "url", "", "path in the parameter store")
}

func runE(cmd *cobra.Command, args []string) error {
	// new Run
	// var check = new(Check)

	// simple check for URL, the
	if cfg.URL == "" {
		return errNoURL
	}

	// make new request
	req, err := http.NewRequest("GET", cfg.URL, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// create new ctx
	ctx, cancel := context.WithTimeout(req.Context(), cfg.Timeout*time.Second)
	defer cancel()

	// use ctx
	req = req.WithContext(ctx)

	// create client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// check for status code
	if res.StatusCode != http.StatusOK {
		return errUnhealthy
	}

	return nil // noop
}

// func err(cmd *cobra.Command, err error) error {
// 	// configure fields
// 	logger := log.WithFields(log.Fields{
// 		"cmd":  Cmd.Name(),
// 		"args": Cmd.Args,
// 	})
// 	logger.Error(err) // log

// 	return nil // err
// }

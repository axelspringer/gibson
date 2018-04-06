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
	"fmt"

	"github.com/spf13/cobra"

	config "github.com/axelspringer/gibson/cfg"
)

// Cmd exports the check command
var Cmd *cobra.Command

// config
var cfg = config.Config

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
}

func runE(cmd *cobra.Command, args []string) error {
	// new Run
	var run = new(Check)

	fmt.Println(run.args)

	return nil // noop
}

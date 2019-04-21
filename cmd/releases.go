// Copyright © 2019 kPherox <admin@mail.kr-kp.com>
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

package cmd

import (
    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var releasesCmd = &cobra.Command{
    Use:   "releases",
    Short: "Fetch data of github releases.",
    Long:  `This command for fetch github releases.

You can get version & assets list.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        return nil
    },
}

func init() {
    rootCmd.AddCommand(releasesCmd)
}
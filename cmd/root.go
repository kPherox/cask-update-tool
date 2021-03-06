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
    "fmt"
    "strings"

    homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    flag "github.com/spf13/pflag"
)

// Flag alias map
type AliasFlags map[string]string

func (fa AliasFlags) Get(name string) string {
    v, ok := fa[name]
    if ok {
        return v
    } else {
        return name
    }
}

// Global flags
var (
    cfgFile string
    cfgPath string
    cfgName string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:     "cask-update-tool",
    Version: "0.0.1",
    Short:   "Command line tool for cask update, using GitHub Release API.",
    RunE: func(cmd *cobra.Command, args []string) error {
        return fmt.Errorf(`Need subcommand`)
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    rootCmd.Execute()
}

var aliasFlags AliasFlags = map[string]string{
    "config": "config-file",
}

func init() {
    cobra.OnInitialize(initConfig)

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", fmt.Sprintf("alias: --config, config file (default <config-path>/<config-name>.<%s>)", strings.Join(viper.SupportedExts, " | ")))
    rootCmd.PersistentFlags().StringVar(&cfgPath, "config-path", "$HOME", "config file path")
    rootCmd.PersistentFlags().StringVar(&cfgName, "config-name", ".cask-update-tool", "config file name")
    rootCmd.PersistentFlags().SortFlags = false
    rootCmd.Flags().SortFlags = false
    rootCmd.SetGlobalNormalizationFunc(func(_ *flag.FlagSet, name string)flag.NormalizedName{
        return flag.NormalizedName(aliasFlags.Get(name))
    })
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile != "" {
        // Use config file from the flag.
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory when default $HOME.
        if cfgPath == "$HOME" {
            home, err := homedir.Dir()
            if err != nil {
                fmt.Println(err)
                return
            }
            cfgPath = home
        }

        // Search config in home directory with name ".cask-update-tool" (without extension).
        viper.AddConfigPath(cfgPath)
        viper.SetConfigName(cfgName)
    }

    viper.AutomaticEnv() // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}

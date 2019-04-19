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
    "os"
    "strings"

    homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    flag "github.com/spf13/pflag"
)

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
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Usage()
        os.Exit(1)
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func flagAliases(name string) string {
    switch name {
    case "config":
        return "config-file"
    default:
        return name
    }
}

func init() {
    cobra.OnInitialize(initConfig)

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", fmt.Sprintf("alias: --config, config file (default <config-path>/<config-name>.<%s>)", strings.Join(viper.SupportedExts, " | ")))
    rootCmd.PersistentFlags().StringVar(&cfgPath, "config-path", "$HOME", "config file path")
    rootCmd.PersistentFlags().StringVar(&cfgName, "config-name", ".cask-update-tool", "config file name")
    rootCmd.PersistentFlags().SortFlags = false
    rootCmd.Flags().SortFlags = false
    rootCmd.SetGlobalNormalizationFunc(func(_ *flag.FlagSet, name string)flag.NormalizedName{
        return flag.NormalizedName(flagAliases(name))
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
                os.Exit(1)
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

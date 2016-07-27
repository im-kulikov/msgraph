// Copyright © 2016 Seth Wright <seth@crosse.org>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	debug        bool
	httpdebug    bool
	clientID     string
	clientSecret string
	tenantDomain string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "msgraph",
	Short: "A command-line interface to the Microsoft Graph API",
	Long:  `A command-line utility to interact with the Microsoft Graph API.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.msgraph.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
	RootCmd.PersistentFlags().BoolVarP(&httpdebug, "httpdebug", "", false, "Enable HTTP debug logging")
	RootCmd.PersistentFlags().StringVar(&tenantDomain, "tenantDomain", "", "Tenant domain")
	RootCmd.PersistentFlags().StringVar(&clientID, "id", "", "OAuth2 Client ID")
	RootCmd.PersistentFlags().StringVar(&clientSecret, "secret", "", "OAuth2 Client Secret")

	viper.BindPFlag("tenantDomain", RootCmd.PersistentFlags().Lookup("tenantDomain"))
	viper.BindPFlag("clientID", RootCmd.PersistentFlags().Lookup("id"))
	viper.BindPFlag("clientSecret", RootCmd.PersistentFlags().Lookup("secret"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("httpdebug", RootCmd.PersistentFlags().Lookup("httpdebug"))

	viper.SetDefault("debug", false)
	viper.SetDefault("httpdebug", false)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("msgraph") // name of config file (without extension)
	viper.AddConfigPath(".")       // add current directory as the first search path
	viper.AddConfigPath("$HOME")   // adding home directory as the next search path
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Fprintf(os.Stderr, "Using config file: %v\n", viper.ConfigFileUsed())
	}
}

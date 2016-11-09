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
	"sort"

	"github.com/Crosse/msgraph"
	"github.com/fatih/structs"
	"github.com/guregu/null"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userGetCmd represents the get command
var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user from the Graph API",
	Long:  `Get a user from the Graph API`,
	Run: func(cmd *cobra.Command, args []string) {
		getUsers(args[0], args[1:])
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)
}

func setupAPI() *msgraph.GraphAPI {
	api := msgraph.New(viper.GetString("tenantDomain"))
	api.SetDebug(viper.GetBool("debug"))
	api.SetHTTPDebug(viper.GetBool("httpdebug"))
	api.SetClientID(viper.GetString("clientID"))
	api.SetClientSecret(viper.GetString("clientSecret"))

	return api
}

func getUsers(user string, properties []string) {
	api := setupAPI()

	// Get the user from the Graph.
	result, err := api.GetUser(user, properties)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		return
	}

	// Convert the object to a map.
	m := structs.Map(result)

	// Get longest property name and sort all property names
	// alphabetically.
	var props []string
	longest := 0
	for k, _ := range m {
		props = append(props, k)
		if len(k) > longest {
			longest = len(k)
		}
	}
	sort.Strings(props)

	// Write out the properties.
	format := fmt.Sprintf("%%-%ds: %%v\n", longest)
	for _, prop := range props {
		fmt.Printf(format, prop, getValueAsString(m, prop))
	}
}

func getValueAsString(user map[string]interface{}, property string) string {
	var t interface{}
	t = user[property]

	switch t := t.(type) {
	case null.Bool:
		if t.Valid {
			return fmt.Sprintf("%s", t.NullBool.Bool)
		}
	case null.Time:
		if t.Valid {
			return fmt.Sprintf("%s", t.Time)
		}
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", t)
	}

	return ""
}

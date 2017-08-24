// <--
// Copyright © 2017 AppNexus Inc.
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
// -->

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reverseList bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the current alter chain",
	Long: `
List the current alters in order for how they are applied.
Note that a '*' indicates whether or not the alter has been
applied to the database in the current configs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolVarP(&reverseList, "reverse", "r", false,
		"List the contents of current alter chain in reverse order")
}

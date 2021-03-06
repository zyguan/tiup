// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import "github.com/spf13/cobra"

func newCompletionsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate tab-completion scripts for your shell",
		Long: `Enable tab completion for Bash, Zsh, The script is
output on stdout, allowing one to re-direct the output to the
file of their choosing. Where you place the file will depend
on which shell, and which operating system you are using. Your
particular configuration may also determine where these scripts
need to be placed.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

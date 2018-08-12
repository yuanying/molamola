/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package commands

import (
	"flag"

	"github.com/spf13/cobra"
)

// NewDefaultCommand returns the default (aka root) command for kustomize command.
func NewDefaultCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "molamola",
		Short: "kustomize manages declarative configuration of Kubernetes",
		Long: `
kustomize manages declarative configuration of Kubernetes.

See https://github.com/yuanying/molamola
`,
	}

	// c.AddCommand(
	// 	// TODO: Make consistent API for newCmd* functions.
	// 	newCmdBuild(stdOut, fsys),
	// 	newCmdDiff(stdOut, stdErr, fsys),
	// 	newCmdEdit(fsys),
	// 	newCmdVersion(stdOut),
	// )
	c.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	// Workaround for this issue:
	// https://github.com/kubernetes/kubernetes/issues/17162
	flag.CommandLine.Parse([]string{})
	return c
}

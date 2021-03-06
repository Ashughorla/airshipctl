/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package cmd

import (
	"io"

	"github.com/spf13/cobra"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"opendev.org/airship/airshipctl/cmd/baremetal"
	"opendev.org/airship/airshipctl/cmd/cluster"
	"opendev.org/airship/airshipctl/cmd/completion"
	"opendev.org/airship/airshipctl/cmd/config"
	"opendev.org/airship/airshipctl/cmd/document"
	"opendev.org/airship/airshipctl/cmd/image"
	"opendev.org/airship/airshipctl/cmd/phase"
	"opendev.org/airship/airshipctl/cmd/secret"
	"opendev.org/airship/airshipctl/pkg/environment"
	"opendev.org/airship/airshipctl/pkg/log"
)

// NewAirshipCTLCommand creates a root `airshipctl` command with the default commands attached
func NewAirshipCTLCommand(out io.Writer) *cobra.Command {
	rootCmd, settings := NewRootCommand(out)
	return AddDefaultAirshipCTLCommands(rootCmd, settings)
}

// NewRootCommand creates the root `airshipctl` command. All other commands are
// subcommands branching from this one
func NewRootCommand(out io.Writer) (*cobra.Command, *environment.AirshipCTLSettings) {
	var debug bool
	rootCmd := &cobra.Command{
		Use:           "airshipctl",
		Short:         "A unified entrypoint to various airship components",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.Init(debug, cmd.OutOrStdout())
		},
	}
	rootCmd.SetOut(out)
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable verbose output")

	return rootCmd, makeRootSettings(rootCmd)
}

// AddDefaultAirshipCTLCommands is a convenience function for adding all of the
// default commands to airshipctl
func AddDefaultAirshipCTLCommands(cmd *cobra.Command, settings *environment.AirshipCTLSettings) *cobra.Command {
	cmd.AddCommand(baremetal.NewBaremetalCommand(settings))
	cmd.AddCommand(cluster.NewClusterCommand(settings))
	cmd.AddCommand(completion.NewCompletionCommand())
	cmd.AddCommand(document.NewDocumentCommand(settings))
	cmd.AddCommand(config.NewConfigCommand(settings))
	cmd.AddCommand(image.NewImageCommand(settings))
	cmd.AddCommand(secret.NewSecretCommand())
	cmd.AddCommand(phase.NewPhaseCommand(settings))
	cmd.AddCommand(NewVersionCommand())

	return cmd
}

// makeRootSettings holds all actions about environment.AirshipCTLSettings
func makeRootSettings(cmd *cobra.Command) *environment.AirshipCTLSettings {
	settings := &environment.AirshipCTLSettings{}
	settings.InitFlags(cmd)
	return settings
}

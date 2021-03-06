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

package config

import (
	"github.com/spf13/cobra"

	"opendev.org/airship/airshipctl/pkg/config"
	"opendev.org/airship/airshipctl/pkg/environment"
)

// NewConfigCommand creates a command for interacting with the airshipctl configuration.
func NewConfigCommand(rootSettings *environment.AirshipCTLSettings) *cobra.Command {
	configRootCmd := &cobra.Command{
		Use:                   "config",
		DisableFlagsInUseLine: true,
		Short:                 "Manage the airshipctl config file",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if parentPreRun := cmd.Root().PersistentPreRun; parentPreRun != nil {
				parentPreRun(cmd.Root(), args)
			}

			if cmd.Use == "init" {
				rootSettings.Create = true
			}
			// Load or Initialize airship Config
			rootSettings.InitConfig()
		},
	}

	cfgFactory := config.CreateFactory(&rootSettings.AirshipConfigPath, &rootSettings.KubeConfigPath)

	configRootCmd.AddCommand(NewGetContextCommand(cfgFactory))
	configRootCmd.AddCommand(NewSetContextCommand(cfgFactory))

	configRootCmd.AddCommand(NewGetManagementConfigCommand(cfgFactory))
	configRootCmd.AddCommand(NewSetManagementConfigCommand(cfgFactory))

	configRootCmd.AddCommand(NewImportCommand(cfgFactory))
	configRootCmd.AddCommand(NewUseContextCommand(cfgFactory))

	configRootCmd.AddCommand(NewGetManifestCommand(cfgFactory))
	configRootCmd.AddCommand(NewSetManifestCommand(cfgFactory))

	// Init will have different factory
	configRootCmd.AddCommand(NewInitCommand(rootSettings))
	return configRootCmd
}

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

import (
	"github.com/c4pt0r/tiup/pkg/meta"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var all bool
	var nightly bool
	var force bool
	cmd := &cobra.Command{
		Use:   "update [component1] [component2..N]",
		Short: "Update tiup components to the latest version",
		RunE: func(cmd *cobra.Command, components []string) error {
			if (len(components) == 0 && !all && !force) || (len(components) > 0 && all) {
				return cmd.Help()
			}
			return updateComponents(components, nightly, force)
		},
	}
	cmd.Flags().BoolVar(&all, "all", false, "Update all components")
	cmd.Flags().BoolVar(&nightly, "nightly", false, "Update the components to nightly version")
	cmd.Flags().BoolVar(&force, "force", false, "Force update a component to the latest version")
	return cmd
}

func updateComponents(components []string, nightly, force bool) error {
	if len(components) == 0 {
		installed, err := profile.InstalledComponents()
		if err != nil {
			return err
		}
		components = installed
	}
	for _, component := range components {
		manifest, err := repository.ComponentVersions(component)
		if err != nil {
			return err
		}
		var latestVer meta.Version
		if nightly {
			latestVer = manifest.LatestNightly()
		} else {
			latestVer = manifest.LatestStable()
		}
		versions, err := profile.InstalledVersions(component)
		if err != nil {
			return err
		}

		if !force {
			var found bool
			for _, v := range versions {
				if meta.Version(v) == latestVer {
					found = true
					break
				}
			}
			if found {
				continue
			}
		}

		err = repository.DownloadComponent(profile.ComponentsDir(), component, latestVer)
		if err != nil {
			return err
		}
	}
	return nil
}

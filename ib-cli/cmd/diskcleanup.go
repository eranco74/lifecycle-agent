/*
Copyright 2023.

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

package cmd

import (
	"fmt"

	"github.com/openshift/assisted-installer/src/ops"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var diskCleanupCmd = &cobra.Command{
	Use:   "disk-cleanup",
	Short: "Clean the installation disk.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := diskCleanup(); err != nil {
			log.Fatalf("Error executing installationiso command: %v", err)
		}
	},
}

func init() {
	// Add installationiso command
	rootCmd.AddCommand(diskCleanupCmd)
	addFlags(diskCleanupCmd)
	diskCleanupCmd.MarkFlagRequired("installation-disk")
}

func diskCleanup() error {

	var err error
	log.Info("Cleaning Installation disk")

	diskCleanup = ops.NewDiskCleanup()

	isoCreator := diskCleanup. (log, op, workDir)
	if err = isoCreator.Create(seedImage, seedVersion, authFile, pullSecretFile, sshPublicKeyFile, lcaImage, rhcosLiveIso,
		installationDisk, extraPartitionStart, precacheBestEffort, precacheDisabled); err != nil {
		err = fmt.Errorf("failed to create installation ISO: %w", err)
		log.Errorf(err.Error())
		return err
	}

	log.Info("Installation ISO created successfully!")
	return nil
}

// Execute executes the root command.
func Execute() error {
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   noColor,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	fmt.Println(`ib-cli assists Image Based Install (IBI).

  Find more information at: https://github.com/openshift-kni/lifecycle-agent/blob/main/ib-cli/README.md
	`)
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("failed to run root cmd for iso: %w", err)
	}
	return nil
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)
const version = "0.0.3"
var rootCmd = &cobra.Command{
	Use:   "maggie",
	Short: "Maggie is love",
	Long: `Maggie does many things`,
}

var selfSubCmd = &cobra.Command{
	Use:   "self",
	Short: "Self",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Maggie",
	Long:  `All software has versions. This is Maggies's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Maggie v0.0.1 -- HEAD")
	},
}

var updateCmd = &cobra.Command{
	Use: "update",
	Short: "Update maggie, if needed",
	Long: `We'll check if a newer version exists, if so, we'll update Maggie'`,
	Run: func(mcd *cobra.Command, args []string) {

		v := semver.MustParse(version)
		latest, err := selfupdate.UpdateSelf(v, "nodeone/maggie")
		if err != nil {
			log.Println("Binary update failed:", err)
			return
		}
		if latest.Version.Equals(v) {
			// latest version is the same as current version. It means current binary is up to date.
			log.Println("Current binary is the latest version", version)
		} else {
			log.Println("Successfully updated to version", latest.Version)
			log.Println("Release note:\n", latest.ReleaseNotes)
		}
	},
}

func init() {
	selfSubCmd.AddCommand(versionCmd)
	selfSubCmd.AddCommand(updateCmd)

	rootCmd.AddCommand(selfSubCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
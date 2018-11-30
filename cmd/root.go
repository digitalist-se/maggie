package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

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

func init() {
	selfSubCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(selfSubCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
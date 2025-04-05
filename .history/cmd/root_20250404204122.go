package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var vers bool

func FullVersion() string {
	version := fmt.Sprintf("Version   : 1.0")
	return version
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-shell",
	Short: "go-shell",
	Long:  "go-shell",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(FullVersion())
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print go-shell version")
}

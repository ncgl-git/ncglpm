package cmd

import (
	"github.com/ncgl-git/ncglpm/internal"
	"github.com/spf13/cobra"
)

// shellhudCmd represents the shellhud command
var shellhudCmd = &cobra.Command{
	Use:   "shellhud",
	Short: "Trigger output for shell startup",
	Long: `For use in a shell profile or rc file, to display various information upon shell startup.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.DisplayNetTree()
	},
}

func init() {
	rootCmd.AddCommand(shellhudCmd)
}

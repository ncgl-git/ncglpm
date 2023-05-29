package cmd

import (
	"fmt"
	"runtime"
	"os/exec"
	"github.com/spf13/cobra"
)

// poweroffCmd represents the poweroff command
var poweroffCmd = &cobra.Command{
	Use:   "poweroff",
	Short: "Turn off the initiating computer. Meant for computers with a wake-on-rtc or wake-on-lan capability.",
	Long: `Turn off the initiating computer. Meant for computers with a wake-on-rtc or wake-on-lan capability. Intended to save electricity / hardware wear while being off during desired times.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch runtime.GOOS {
		case "windows":
			if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
				fmt.Println("Failed to initiate shutdown:", err)
			}
		case "linux", "freebsd", "netbsd", "openbsd": 
		if err := exec.Command("cmd", "shutdown").Run(); err != nil {
			fmt.Println("Failed to initiate shutdown:", err)
		}
		}

	},
}

func init() {
	rootCmd.AddCommand(poweroffCmd)

	poweroffCmd.Flags().String(
		"manual",
		"",
	    "Force execution of input command instead. Intended when user is running on an unsupported OS, but wants to maintain consistency and use this to manage behaviour.",
	)
}

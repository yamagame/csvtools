package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamagame/csvtools"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of joincsv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "joincsv",
	Short: "join csv",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var inputLeftReader io.Reader = cmd.InOrStdin()
		var inputRightReader io.Reader = cmd.InOrStdin()
		if len(args) > 0 && args[0] != "-" {
			file, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			inputLeftReader = file
		}
		if len(args) > 1 && args[1] != "-" {
			file, err := os.Open(args[1])
			if err != nil {
				panic(err)
			}
			inputRightReader = file
		}
		left, err := csvtools.Read(inputLeftReader)
		if err != nil {
			panic(err)
		}
		right, err := csvtools.Read(inputRightReader)
		if err != nil {
			panic(err)
		}
		result := csvtools.Join(left, right)
		csvtools.Dump(result)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

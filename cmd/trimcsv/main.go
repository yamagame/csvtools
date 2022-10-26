package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamagame/csvtools"
)

var colpos int
var rowpos int
var colsize int
var rowsize int

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of trimcsv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "trimcsv",
	Short: "trim csv",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var inputReader io.Reader = cmd.InOrStdin()
		if len(args) > 0 && args[0] != "-" {
			file, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			inputReader = file
		}
		csv, err := csvtools.Read(inputReader)
		if err != nil {
			panic(err)
		}
		csvtools.Dump(csvtools.Trim(csv, colpos, rowpos, colsize, rowsize))
	},
}

func init() {
	rootCmd.Flags().IntVarP(&colpos, "col", "c", 0, "col position")
	rootCmd.Flags().IntVarP(&rowpos, "row", "r", 0, "row position")
	rootCmd.Flags().IntVarP(&colsize, "size", "s", 1, "col size")
	rootCmd.Flags().IntVarP(&rowsize, "line", "l", 1, "row size")
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

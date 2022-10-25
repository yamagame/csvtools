package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamagame/csvtools"
)

var keypos int
var skipline int

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of sortcsv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "sortcsv",
	Short: "transpose csv",
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
		if skipline > 0 {
			header := csv[:skipline]
			targetrecords := csv[skipline:]
			sortedrecords := csvtools.Sort(targetrecords, keypos)
			result := append(header, sortedrecords...)
			csvtools.Dump(result)
		} else {
			csvtools.Dump(csvtools.Sort(csv, keypos))
		}
	},
}

func init() {
	rootCmd.Flags().IntVarP(&keypos, "key", "k", 0, "sort key position")
	rootCmd.Flags().IntVarP(&skipline, "skip", "s", 0, "skip line count")
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

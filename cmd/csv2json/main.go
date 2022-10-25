package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamagame/csvtools"
)

var rowFlag *bool
var colFlag *bool
var keypos int
var valuepos int

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of csv2json",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "csv2json",
	Short: "convert csv to hash",
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
		var hash string
		if !*rowFlag && !*colFlag {
			hash = csvtools.Json(csv)
		} else if *rowFlag {
			hash = csvtools.Json(csvtools.RowHash(csv, keypos, valuepos))
		} else if *colFlag {
			hash = csvtools.Json(csvtools.ColHash(csv, keypos, valuepos))
		}
		fmt.Printf("%s\n", hash)
	},
}

func init() {
	rowFlag = rootCmd.Flags().BoolP("row", "r", false, "row base hash")
	colFlag = rootCmd.Flags().BoolP("col", "c", false, "col base hash")
	rootCmd.Flags().IntVarP(&keypos, "key", "k", 0, "key position")
	rootCmd.Flags().IntVarP(&valuepos, "value", "v", 1, "value position")
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

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gencode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "gencode",
	Short: "generate code",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		className, err := cmd.PersistentFlags().GetString("class")
		if err != nil {
			return err
		}
		templateClassName, err := cmd.PersistentFlags().GetString("target")
		if err != nil {
			return err
		}
		var inputReader io.Reader = cmd.InOrStdin()
		if len(args) > 0 {
			file, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			inputReader = file
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(inputReader)
		params := map[string]string{}
		params["class"] = className
		src := buf.String()
		if templateClassName != "" {
			src = strings.ReplaceAll(src, templateClassName, "{{.class}}")
		}
		tmpl, err := template.New("gencode").Parse(src)
		if err != nil {
			return err
		}
		return tmpl.Execute(os.Stdout, params)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("class", "c", "Class", "Generate with class name")
	rootCmd.PersistentFlags().StringP("target", "t", "", "Replaced class name")
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

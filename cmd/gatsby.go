package cmd

import (
	"strings"

	"github.com/TheBoringDude/tailwify/worker"
	"github.com/spf13/cobra"
)

// gatsbyCmd represents the gatsby command
var gatsbyCmd = &cobra.Command{
	Use:   "gatsby",
	Short: "Gatsby.js App",
	Long:  `Configure and setup a Gatsby.js APP`,
	Run: func(cmd *cobra.Command, args []string) {
		generate := &worker.Worker{
			AppType:     "gatsby",
			ProjectName: strings.ToLower(projectName), // npm & yarn doesn't allow having caps in project names
			JsApp:       true,
		}

		// generate
		generate.Start()
	},
}

func init() {
	generateCmd.AddCommand(gatsbyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gatsbyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatsbyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
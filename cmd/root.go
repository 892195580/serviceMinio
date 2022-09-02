package cmd

import (
"fmt"
"os"
//"runtime"
//"strconv"

log "github.com/sirupsen/logrus"
"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mgo",
	Short: "The mgo command line interface.",
	Long:  `The mgo command line interface lets you create and manage MogDB clusters.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("Execute called")

	if err := RootCmd.Execute(); err != nil {
		log.Debug(err.Error())
		os.Exit(-1)
	}
}
func init() {
	log.Debug("init called")
}


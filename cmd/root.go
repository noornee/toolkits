/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "toolkits",
	Short: "A collection of personal utility tools to streamline development workflow",
	Long: `
Toolkits is a personal command-line application that bundles various utilities 
to streamline your development workflow. It provides convenient interfaces to 
common tools and services, allowing you to execute routine tasks efficiently
through simple commands. Designed to be extensible, you can easily add new
utilities as your needs evolve.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toolkits.yaml)")

	// setCmd.Flags().StringVarP(&key, "key", "k", "", "The Redis key where data will be stored")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

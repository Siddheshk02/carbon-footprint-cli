/*
Copyright © 2023 Siddhesh Khandagale
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Carbon-Footprint-CLI",
	Short: "A Carbon Footprint Calculator",
	Long:  `A Carbon Footprint Calculator CLI built using Golang.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		magenta := color.New(color.FgHiMagenta).SprintFunc()
		fmt.Print("\n")
		emoji.Println(red("Welcome to Carbon Footprint Calculator CLI! :hand:"))
		fmt.Print("\n")
		emoji.Println(green("Calculating one's carbon footprint is a crucial step in understanding how our daily actions \t\nand choices impact the environment. :shamrock:"))
		fmt.Println(yellow("\nUsing this tool we can raise awareness about the importance of reducing our greenhouse gas emissions \t\nand encourage people to take action towards a more sustainable lifestyle."))
		fmt.Println(magenta("\nMoreover, with the growing concern over climate change and its impact on our planet, \t\nthere is a need for more tools and resources that can help individuals \t\nand businesses understand and reduce their carbon footprint."))
		fmt.Print("\n")
		// content := "Calculating one's carbon footprint is a crucial step in understanding how our daily actions and choices impact the environment. Using this tool we can raise awareness about the importance of reducing our greenhouse gas emissions and encourage people to take action towards a more sustainable lifestyle. Moreover, with the growing concern over climate change and its impact on our planet, there is a need for more tools and resources that can help individuals and businesses understand and reduce their carbon footprint."
		// fmt.Printf("%s\t\n\n", content)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Carbon-Footprint-CLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2023 Siddhesh Khandagale
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CalculateCmd represents the Calculate command
var CalculateCmd = &cobra.Command{
	Use:   "Calculate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Calculate called")
		var opt int
		fmt.Println("Select the Option for which you want to calculate the Carbon Equivalent : ")

		fmt.Println("  1. Traditional Energy")
		fmt.Println("  2. Car Travel.")
		fmt.Println("  3. Flight.")
		fmt.Println("  4. MotorBike.")
		fmt.Println("  5. Public Transportation.")
		fmt.Println("  6. Clean Energy.")

		fmt.Println("\nEnter your Option : ")
		fmt.Scanln(&opt)
	},
}

func init() {
	rootCmd.AddCommand(CalculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CalculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CalculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

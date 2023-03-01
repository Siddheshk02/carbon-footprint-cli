/*
Copyright © 2023 Siddhesh Khandagale
*/
package cmd

import (
	"fmt"

	"github.com/Siddheshk02/Carbon-Footprint-CLI/lib"
	"github.com/spf13/cobra"
)

// type Body struct {
// 	first  int  `json:"carbonEquivalent"`
// 	second bool `json:"success"`
// }

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

		switch opt {
		case 1:
			var en int
			var loc string
			fmt.Print("Enter the Energy usage(KWH) :")
			fmt.Scanln(&en)
			fmt.Print("Enter the location (USA, Canada, UK, Europe, Africa, LatinAmerica, MiddleEast, OtherCountry):")
			fmt.Scanln(&loc)
			lib.Traditional(en, loc)

		case 2:
			var dis int
			var vehicle string
			fmt.Print("Enter the Distance Travelled (in KM) :")
			fmt.Scanln(&dis)
			fmt.Println("Select the Vehicle Type : ")
			fmt.Println("  SmallDieselCar ,   MediumDieselCar,")
			fmt.Println("  LargeDieselCar ,   MediumHybridCar,")
			fmt.Println("  LargeHybridCar ,   MediumLPGCar,")
			fmt.Println("  LargeLPGCar ,      MediumCNGCar,")
			fmt.Println("  LargeCNGCar ,      SmallPetrolVan,")
			fmt.Println("  LargePetrolVan ,   SmallDielselVan,")
			fmt.Println("  MediumDielselVan , LargeDielselVan,")
			fmt.Println("  LPGVan ,           CNGVan,")
			fmt.Println("  SmallPetrolCar ,   MediumPetrolCar , LargePetrolCar")
			fmt.Print("Enter the Vehicle type(Name) : ")
			fmt.Scanln(&vehicle)
			lib.CarTravel(dis, vehicle)

		}

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

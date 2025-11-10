/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"net/netip"
	"os"

	"github.com/heikkilavaste/go_subnet/modules/filer"
	"github.com/heikkilavaste/go_subnet/modules/ranger"
	local_types "github.com/heikkilavaste/go_subnet/modules/types"
	"github.com/spf13/cobra"
	"go4.org/netipx"
)

// breakdownCmd represents the breakdown command
var logger *slog.Logger
var b netipx.IPSetBuilder
var Plist []netip.Prefix
var Sets []local_types.AddressSet
var breakdownCmd = &cobra.Command{
	Use:   "breakdown",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("breakdown called")
		Nw, _ := cmd.Flags().GetStringSlice("network")
		Sn, _ := cmd.Flags().GetInt("size")
		Output, _ := cmd.Flags().GetString("output")
		Ofile, _ := cmd.Flags().GetString("file")
		Tab, _ := cmd.Flags().GetString("tab")
		Range := ranger.NewRange(Nw)
		out := Range.BreakDown(nil, uint8(Sn))
		for _, r := range out {
			Range2 := ranger.NewRange([]string{r})
			Sets = append(Sets, Range2.Parse())

		}
		switch Output {
		case "console":
			filer.WriteToConsole(Sets)
		case "yaml":
			filer.WriteToYaml(Sets, Ofile)
		case "json":
			fmt.Println(Output, " not yet implemented. ")
		case "csv":
			filer.WriteToCSV(Sets, Ofile, Tab)
		case "netbox":
			fmt.Println(Output, " not yet implemented. ")
		}

	},
}

func init() {
	rootCmd.AddCommand(breakdownCmd)
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	breakdownCmd.Flags().StringSliceP("network", "n", []string{}, "network prefix")
	breakdownCmd.Flags().IntP("size", "s", 0, "subnet size")
	breakdownCmd.Flags().StringP("output", "o", "console", "output")
	breakdownCmd.Flags().StringP("file", "f", "subnets", "write to")
	breakdownCmd.Flags().StringP("tab", "t", "tab1", "excel tab name")
	breakdownCmd.MarkFlagRequired("network")
}

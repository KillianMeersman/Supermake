package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
	"github.com/spf13/cobra"
)

var file string = "./Supermake"
var skipTargets []string = []string{}

var rootCmd = &cobra.Command{
	Use:   "supermake target [target...]",
	Short: "Supermake is a modern CI pipelining tool",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := parse.ParseSupermakeFileV2(file)
		if err != nil {
			log.Fatal(err)
		}

		for _, skip := range skipTargets {
			if target, ok := file.Targets[skip]; ok {
				target.Done()
			}
		}

		for _, target := range args {
			err = file.Run(context.Background(), target)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "./Supermake", "Supermake file to use")
	rootCmd.PersistentFlags().StringArrayVarP(&skipTargets, "outdated", "o", []string{}, "Targets to skip")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

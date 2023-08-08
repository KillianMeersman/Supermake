package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
	"github.com/spf13/cobra"
)

var file string = "./Supermake"
var cwd string = "."
var skipTargets []string = []string{}
var timeout time.Duration

var rootCmd = &cobra.Command{
	Use:   "supermake target [target...]",
	Short: "Supermake is a modern CI pipelining tool",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := parse.ParseSupermakeFileV2(cwd, file)
		if err != nil {
			log.Fatal(err)
		}

		for _, skip := range skipTargets {
			if target, ok := file.Targets[skip]; ok {
				target.Done()
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		for _, target := range args {
			err = file.Run(ctx, target, cwd)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "./Supermake", "Supermake file to use")
	rootCmd.PersistentFlags().StringVarP(&cwd, "cwd", "c", ".", "Working directory")
	rootCmd.PersistentFlags().StringArrayVarP(&skipTargets, "outdated", "o", []string{}, "Targets to skip")
	rootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 10*time.Minute, "Timeout, expressed as a string e.g. '5m'")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

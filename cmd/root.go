package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/KillianMeersman/Supermake/pkg/supermake"
	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
	"github.com/spf13/cobra"
)

var file string = "./Supermake"
var cwd string = "."
var skipTargets []string = []string{}
var timeout time.Duration

var root = &cobra.Command{
	Use:   "supermake",
	Short: "Supermake is a modern CI pipelining tool",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please provide a command, use --help for more information")
	},
}

var run = &cobra.Command{
	Use:   "run target [...target]",
	Short: "Run one or more targets",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			cancel()
		}()

		file, err := parse.ParseSupermakeFileV2(cwd, file)
		if err != nil {
			log.Fatal(err)
		}

		scheduler := supermake.NewLocalScheduler()

		for _, skip := range skipTargets {
			if target, ok := file.Targets[skip]; ok {
				err := scheduler.TargetDone(ctx, target)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		if !path.IsAbs(cwd) {
			cd, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			cwd = path.Join(cd, cwd)
		}

		if len(args) == 0 {
			fmt.Println("Please provide at least one target.")
			fmt.Print(file.Help())
			return
		}
		err = file.Run(ctx, scheduler, cwd, args...)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Get the Supermake version",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	root.AddCommand(run, version)
	run.PersistentFlags().StringVarP(&file, "file", "f", "./Supermake", "Supermake file to use")
	run.PersistentFlags().StringVarP(&cwd, "cwd", "c", ".", "Working directory")
	run.PersistentFlags().StringArrayVarP(&skipTargets, "omit", "o", []string{}, "Targets to omit")
	run.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 2*time.Hour, "Timeout, expressed as a string e.g. '5m'")
}

func Execute() {
	if err := run.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
	"github.com/spf13/cobra"
)

var file string = "./Supermake"

var rootCmd = &cobra.Command{
	Use:   "supermake target [target...]",
	Short: "Supermake is a modern CI pipelining tool",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := parse.ParseSupermakeFileV2(file)
		if err != nil {
			log.Fatal(err)
		}

		for _, target := range args {
			err = file.Run(target)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&file, "file", "./Supermake", "Supermake file to use")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

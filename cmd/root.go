package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "supermake",
	Short: "Supermake is a modern CI pipelining tool",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := parse.ParseSupermakeFileV2("./Supermake")
		if err != nil {
			log.Fatal(err)
		}
		err = file.Run(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

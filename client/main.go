package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
	"os"
	"github.com/miniledger/client/peer"
)

var rootCmd = &cobra.Command{
	Use: "peer",
	//Run:func(cmd *cobra.Command,args []string ){
	//	log.Println(cmd.Use, args)
	//	log.Println("bye!")
	//}
}

const envRoot = "core"

func main() {
	setupEnv()
	rootCmd.AddCommand(peer.Cmd)
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}

func setupEnv() {
	viper.SetEnvPrefix(envRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}

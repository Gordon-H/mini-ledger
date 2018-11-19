package channel

import (
	"fmt"
	"github.com/miniledger/common/tools/configtxgen/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var CreateChannelTxCmd = &cobra.Command{
	Use: "createChannelTx",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("there should not be any arg for this command")
		}
		return execute()
	},
}

var file, channel, output string

func init() {
	CreateChannelTxCmd.Flags().StringVarP(&file, "configFile", "f", "", "config file")
	CreateChannelTxCmd.Flags().StringVarP(&channel, "channel", "c", "", "channel")
	CreateChannelTxCmd.Flags().StringVarP(&output, "outputFile", "o", "", "output file")
}

func execute() error {
	txConfig, err := loadConfig()
	if err!=nil{
		log.Fatal(err)
	}
	channelOrgs  := txConfig.Channels[channel]
	if channelOrgs == nil{
		log.Fatal("the channel is not presented in the given config file")
	}



	return nil

}
func loadConfig() (*config.TxConfig, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	txConfig := config.TxConfig{}
	err = yaml.Unmarshal(bs, &txConfig)
	if err != nil {
		log.Fatal(err)
	}
	return &txConfig, nil
}

package channel

import (
	"fmt"
	"github.com/miniledger/common/tools/configtxgen/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"github.com/miniledger/protos/orderer"
	"github.com/miniledger/protos/utils"
	protos2 "github.com/miniledger/protos/peer"
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

var file, channelName, output string

func init() {
	CreateChannelTxCmd.Flags().StringVarP(&file, "configFile", "f", "", "config file")
	CreateChannelTxCmd.Flags().StringVarP(&channelName, "channel", "c", "", "channel")
	CreateChannelTxCmd.Flags().StringVarP(&output, "outputFile", "o", "", "output file")
}

func execute() error {
	txConfig, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
	channelOrgs := txConfig.Channels[channelName]
	if channelOrgs == nil {
		log.Fatal("the channel is not presented in the given config file")
	}
	createTx, err := createTxFromConfig(txConfig)
	proposal ,err := utils.CreateSignedProposal(protos2.HeaderType_CHANNEL_CREATE, createTx)
	ioutil.WriteFile(output, )

	return nil

}
func createTxFromConfig(txConfig *config.TxConfig) (protos.ChannelCreateTx, error) {
	var orgs []*protos.Organization
	for _, c := range txConfig.Channels[channelName] {
		orgs = append(orgs, &protos.Organization{OrgName: c.Name})
	}
	tx := protos.ChannelCreateTx{ChannelName: channelName, Organizations: orgs}
	return tx, nil
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

package config

type TxConfig struct {
	Genesis       *Genesis                   `yaml:"Genesis"`
	Organizations []*Organization            `yaml:"Organizations"`
	Channels      map[string][]*Organization `yaml:"Channels"`
}
type Genesis struct {
	OrdererType string   `yaml:"OrdererType"`
	Addresses   []string `yaml:"Addresses"`
}
type Organization struct {
	Name   string `yaml:"Name"`
	MSPDir string `yaml:"MSPDir"`
}

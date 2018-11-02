/*
skycoin daemon
*/
package main

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

import (
	"flag"
	_ "net/http/pprof"
	"os"

	"github.com/chazu/skycoin/src/readable"
	"github.com/chazu/skycoin/src/skycoin"
	"github.com/chazu/skycoin/src/util/logging"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.25.0-rc1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "df894f1466e44b5ab0cc0c220f5c751567a3f13e2f4091dd6a186892edb21e0a5d07149eb7ea2d3ebfe167f27cb45db7c32823e27dd61e94b90c45baa73932d901"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "23p53TdHiiF9dsUQGnS7HhjbScRxyjmjUXg"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "03fa74a5559c336d2a5901525a6c35d98d503bd9eb786a61186b9b616878cf927f"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = "c3cdf67a11558ebcebb7f98dee7496c7b8f544432f810cbb472ed473515d8e87"

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1541117729
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 100000000000000

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{}

	nodeConfig = skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "",
		Port:                6000,
		WebInterfacePort:    6420,
		DataDirectory:       "$HOME/.skycoin",
	})

	parseFlags = true
)

func init() {
	nodeConfig.RegisterFlags()
}

func main() {
	if parseFlags {
		flag.Parse()
	}

	// create a new fiber coin instance
	coin := skycoin.NewCoin(skycoin.Config{
		Node: nodeConfig,
		Build: readable.BuildInfo{
			Version: Version,
			Commit:  Commit,
			Branch:  Branch,
		},
	}, logger)

	// parse config values
	if err := coin.ParseConfig(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// run fiber coin node
	if err := coin.Run(); err != nil {
		os.Exit(1)
	}
}

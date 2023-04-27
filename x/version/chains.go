package version

const (
	MochaChainID          = "mocha"
	ArabicaChainID        = "arabica-4"
	BlockspaceraceChainID = "blockspacerace-0"
	MainnetChainID        = "celestia-1"
)

// StandardChainVersions returns a map of chain IDs to their respective
// ChainVersionConfig. Each hardfork should be added to this map.
func StandardChainVersions() map[string]ChainVersionConfig {
	version0Only := NewChainVersionConfig(map[uint64]int64{
		0: 0,
	})
	mainnetVersions := NewChainVersionConfig(map[uint64]int64{
		1: 0,
	})
	return map[string]ChainVersionConfig{
		MochaChainID:          version0Only,
		ArabicaChainID:        version0Only,
		BlockspaceraceChainID: version0Only,
		MainnetChainID:        mainnetVersions,
	}
}
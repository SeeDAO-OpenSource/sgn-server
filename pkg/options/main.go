package options

import (
	"github.com/google/wire"
)

var OptionsSet = wire.NewSet(
	wire.Struct(new(OptionsInitor), "*"),
)

type OptionsInitor struct {
}

func (initor *OptionsInitor) Init() error {
	// err := viper.UnmarshalKey("ApiServer", initor.ApiOptions)
	// err = viper.UnmarshalKey("Etherscan", initor.EtherscanOptions)
	return nil
}

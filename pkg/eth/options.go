package eth

type EthOptions struct {
	DailUrl string
}

func NewEthOptions() *EthOptions {
	return &EthOptions{
		DailUrl: "",
	}
}

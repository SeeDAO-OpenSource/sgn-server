package eth

type EthOptions struct {
	DailUrl string
}

func NewEthOptions() *EthOptions {
	return &EthOptions{
		DailUrl: "wss://mainnet.infura.io/ws/v3/9aa3d95b3bc440fa88ea12eaa4456161",
	}
}

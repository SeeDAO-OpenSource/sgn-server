package erc721

type EtherScanOptions struct {
	ApiKey  string
	BaseURL string
	Proxy   string
}

var EsOptions = &EtherScanOptions{
	BaseURL: "https://api.etherscan.io/api?",
}

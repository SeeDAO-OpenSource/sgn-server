package erc721

type EtherScanOptions struct {
	ApiKey  string
	BaseURL string
	Proxy   string
}

var EsOptions = &EtherScanOptions{
	BaseURL: "https://api.etherscan.io/api?",
	ApiKey:  "ZZKAMSFGQ6KEFD6ZWM3PSI5UAQ724KSY75",
	Proxy:   "http://localhost:4780",
}

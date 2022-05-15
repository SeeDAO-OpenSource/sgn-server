package apiserver

import (
	"io/ioutil"
	"strings"

	blobv1 "github.com/waite-lee/sgn/internal/apiserver/blob/v1"
	nftv1 "github.com/waite-lee/sgn/internal/apiserver/nft/v1"
	"github.com/waite-lee/sgn/pkg/server"
)

func AddApiServer(builder *server.ServerBuiler) {
	nftv1.AddNftV1(builder)
	blobv1.AddBlobV1(builder)
	builder.Configure(func(s *server.Server) error {
		relaceDemoAddress(s.Options)
		return nil
	})
}

func relaceDemoAddress(options *server.ServerOptions) {
	content, err := ioutil.ReadFile("app/demo/index.js")
	if err == nil {
		js := string(content)
		js = strings.ReplaceAll(js, "{ServerURL}", options.ServerURL)
		ioutil.WriteFile("app/demo/index.js", []byte(js), 0644)
	}
}

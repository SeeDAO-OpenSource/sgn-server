package apiserver

import (
	"io/ioutil"
	"strings"

	blobv1 "github.com/SeeDAO-OpenSource/sgn/internal/apiserver/blob/v1"
	idv1 "github.com/SeeDAO-OpenSource/sgn/internal/apiserver/identity/v1"
	sgnv1 "github.com/SeeDAO-OpenSource/sgn/internal/apiserver/sgn/v1"
	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/swagger"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
)

func AddApiServer(builder *server.ServerBuiler) {
	sgnv1.AddSgnV1(builder)
	blobv1.AddBlobV1(builder)
	idv1.AddIdentityV1(builder)
	swagger.AddSwagger(builder)
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

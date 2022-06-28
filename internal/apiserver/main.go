package apiserver

import (
	"io/ioutil"
	"strings"

	blob "github.com/SeeDAO-OpenSource/sgn/internal/apiserver/blob"
	memberapi "github.com/SeeDAO-OpenSource/sgn/internal/apiserver/member"
	sgn "github.com/SeeDAO-OpenSource/sgn/internal/apiserver/sgn"
	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/swagger"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
)

func AddApiServer(builder *server.ServerBuiler) {
	builder.App.ConfigureServices(func() error {
		utils.ViperBind("Server", builder.Options)
		services.AddValue(builder.Options)
		return nil
	})
	builder.Add(sgn.SgnApi).
		Add(memberapi.MemberApi).
		Add(blob.BlobStore).
		Add(swagger.SwaggerDoc).
		Configure(func(s *server.Server) error {
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

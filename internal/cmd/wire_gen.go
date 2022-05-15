// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/waite-lee/sgn/internal/common"
	"github.com/waite-lee/sgn/pkg/app"
)

// Injectors from wire_inject.go:

func BuildCommands(buider *app.AppBuilder) AppCommands {
	testCmd := NewTestCmd()
	apiServerCmd := NewApiServerCmd(buider)
	nftPullCmd := NewNftPullCmd(buider)
	configCmd := NewConfigCmd()
	httpClientOptions := _wireHttpClientOptionsValue
	updateCmd := NewUpdateCmd(httpClientOptions)
	appCommands := AppCommands{
		Test:      testCmd,
		ApiServer: apiServerCmd,
		NftPull:   nftPullCmd,
		Config:    configCmd,
		Update:    updateCmd,
	}
	return appCommands
}

var (
	_wireHttpClientOptionsValue = common.HttpOptions
)

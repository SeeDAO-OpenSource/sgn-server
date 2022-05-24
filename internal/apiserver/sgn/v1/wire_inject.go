//go:build wireinject
// +build wireinject

package sgnv1

import (
	"github.com/google/wire"
)

func BuildSgnServiceV1() (*SgnService, error) {
	wire.Build(
		SgnV1Set,
	)
	return &SgnService{}, nil
}

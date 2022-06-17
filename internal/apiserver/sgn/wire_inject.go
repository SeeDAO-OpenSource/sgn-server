//go:build wireinject
// +build wireinject

package sgn

import (
	"github.com/google/wire"
)

func BuildSgnServiceV1() (*SgnService, error) {
	wire.Build(
		sgnSet,
	)
	return &SgnService{}, nil
}

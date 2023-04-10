package appstats

import (
	"github.com/stretchr/testify/mock"
)

type StatsGeneratorMock struct {
	mock.Mock
}

func (statsGeneratorMock *StatsGeneratorMock) CollectStats(hostName string) []StatsDo {
	args := statsGeneratorMock.Called(hostName)
	return args.Get(0).([]StatsDo)
}

package reports

import (
	"github.com/stretchr/testify/mock"
)

type ReportGeneratorMock struct {
	mock.Mock
}

func (reportGeneratorMock *ReportGeneratorMock) GenerateReport(hosts []string) ReportDo {
	args := reportGeneratorMock.Called(hosts)
	return args.Get(0).(ReportDo)
}

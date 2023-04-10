package fileops

import (
	"io/fs"

	"github.com/stretchr/testify/mock"
)

type OsFileOpsMock struct {
	mock.Mock
}

func (osFileOpsMock *OsFileOpsMock) ReadFile(path string) ([]byte, error) {
	args := osFileOpsMock.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func (osFileOpsMock *OsFileOpsMock) WriteFile(path string, data []byte, permission fs.FileMode) error {
	args := osFileOpsMock.Called(path, data, permission)
	return args.Error(0)
}

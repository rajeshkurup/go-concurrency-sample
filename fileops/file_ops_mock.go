package fileops

import (
	"github.com/stretchr/testify/mock"
)

type FileOpsMock struct {
	mock.Mock
}

func (fileOpsMock *FileOpsMock) ReadFile(path string) string {
	args := fileOpsMock.Called(path)
	return args.String(0)
}

func (fileOpsMock *FileOpsMock) WriteFile(path string, data []byte) error {
	args := fileOpsMock.Called(path, data)
	return args.Error(0)
}

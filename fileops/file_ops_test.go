package fileops

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileSuccess(test *testing.T) {
	osFileOpsMock := OsFileOpsMock{}
	fileOps := FileOps{&osFileOpsMock}

	osFileOpsMock.On("ReadFile", "hosts.txt").Return([]byte("sddsdasd"), nil).Once()

	data := fileOps.ReadFile("hosts.txt")

	osFileOpsMock.AssertExpectations(test)

	assert.Equal(test, "sddsdasd", data, "TestReadFileSuccess Failed: Content doesn't match with expected value")
}

func TestReadFileFailed(test *testing.T) {
	osFileOpsMock := OsFileOpsMock{}
	fileOps := FileOps{&osFileOpsMock}

	osFileOpsMock.On("ReadFile", "hosts.txt").Return([]byte("sddsdasd"), errors.New("File not found")).Once()

	data := fileOps.ReadFile("hosts.txt")

	osFileOpsMock.AssertExpectations(test)

	assert.Equal(test, "", data, "TestReadFileFailed Failed: Content doesn't match with expected value")
}

func TestWriteFileSuccess(test *testing.T) {
	osFileOpsMock := OsFileOpsMock{}
	fileOps := FileOps{&osFileOpsMock}

	osFileOpsMock.On("WriteFile", "hosts.txt", []byte("sddsdasd"), fs.FileMode(0664)).Return(nil).Once()

	err := fileOps.WriteFile("hosts.txt", []byte("sddsdasd"))

	osFileOpsMock.AssertExpectations(test)

	assert.NoError(test, err, "TestWriteFileSuccess Failed")
}

func TestWriteFileFailed(test *testing.T) {
	osFileOpsMock := OsFileOpsMock{}
	fileOps := FileOps{&osFileOpsMock}

	osFileOpsMock.On("WriteFile", "hosts.txt", []byte("sddsdasd"), fs.FileMode(0664)).Return(errors.New("Failed to write hosts.txt")).Once()

	err := fileOps.WriteFile("hosts.txt", []byte("sddsdasd"))

	osFileOpsMock.AssertExpectations(test)

	assert.Equal(test, "Failed to write hosts.txt", err.Error(), "TestWriteFileFailed Failed")
}

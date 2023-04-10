package fileops

import (
	"io/fs"
	"os"
)

/**
 * @brief Interface for OS File Operations
 */
type OsFileOpsInterface interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte, permission fs.FileMode) error
}

/**
 * @brief Helper for performing OS File Operations
 */
type OsFileOps struct {
	// Empty
}

/**
 * @brief Constructor for OsFileOps
 * @return Instance of OsFileOps
 */
func MakeOsFileOps() OsFileOps {
	return OsFileOps{}
}

/**
 * @brief A public method to read content of a file
 * @param path Fully qualified path to the file
 * @return Content of the file as byte array if succeeded
 * @return Instance of error if failed
 */
func (osFileOps *OsFileOps) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

/**
 * @brief A public method to write data to a file
 * @param path Fully qualified path to the file
 * @param data Data to be written as byte array
 * @param permission Permission to be set on the file
 * @return Instance of error if failed otherwise none
 */
func (osFileOps *OsFileOps) WriteFile(path string, data []byte, permission fs.FileMode) error {
	return os.WriteFile(path, data, permission)
}

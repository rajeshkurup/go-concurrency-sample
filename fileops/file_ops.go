package fileops

import (
	"fmt"
	"log"
	"strings"
)

/**
 * @brief Interface for File Operations
 */
type FileOpsInterface interface {
	ReadFile(path string) string
	WriteFile(path string, data []byte) error
}

/**
 * @brief Helper for performing File Operations
 */
type FileOps struct {
	osFileOps OsFileOpsInterface
}

/**
 * @brief Constructor for FileOps
 * @return An object of FileOps
 */
func MakeFileOps() FileOps {
	osFileOps := MakeOsFileOps()
	return FileOps{&osFileOps}
}

/**
 * @brief A public method to read file from local file system
 * @param path Fully qualified path to file
 * @return Content of the file a string object
 */
func (fileOps *FileOps) ReadFile(path string) string {
	hosts := ""
	data, err := fileOps.osFileOps.ReadFile(path)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to read file %s - error=%s", path, err.Error()))
	} else {
		hosts = strings.TrimSpace(string(data))
	}

	return hosts
}

/**
 * @brief A public method to write file into local file system.
 *			File permission would be set as Read/Write for Owner and Group, Read for others.
 * @param path Fully qualified path to file
 * @param data Data to written into file as byte array
 * @return Instance of error if failed otherwise none
 */
func (fileOps *FileOps) WriteFile(path string, data []byte) error {
	err := fileOps.osFileOps.WriteFile(path, data, 0664)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to write file %s - error=%s", path, err.Error()))
	}

	return err
}

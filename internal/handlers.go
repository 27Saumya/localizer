package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandlePath(sourcePath string) error {
	if sourceInfo, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("❌ path '%s' does not exist", sourcePath)
	} else {
		if sourceInfo.IsDir() {
			dirName, err := directoryName(sourcePath)
			if err != nil {
				return err
			}
			_, ok := Languages[dirName[:len(dirName)-6]]
			if !strings.HasSuffix(dirName, ".lproj") || !ok {
				return fmt.Errorf("❌ invalid directory name '%s'", dirName)
			}
			joinedPath := filepath.Join(sourcePath, "Localizable.strings")
			if _, err := os.Stat(joinedPath); os.IsNotExist(err) {
				return fmt.Errorf("❌ couldn't locate 'Localizable.strings' in '%s'", sourcePath)
			}
		} else {
			return fmt.Errorf("❌ expected a '.lproj' directory, found a file")
		}
	}
	return nil
}
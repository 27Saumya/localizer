package internal

import (
	"fmt"
	"path/filepath"
)

func directoryName(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("❌ failed to get the path of '%s'", absPath)
	}
	return filepath.Base(absPath), nil
}

func switchLanguage(lang *string) {
	lanuages := map[string]string{
		"zh-hans": "zh-cn",
		"zh-hant": "zh-tw",
		"zh-cn":   "zh-Hans",
		"zh-tw":   "zh-Hant",
	}
	if val, ok := lanuages[*lang]; ok {
		*lang = val
	}
}
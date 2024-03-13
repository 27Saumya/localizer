package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bregydoc/gtranslate"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func Localize(path string, args []string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	dirName, err := directoryName(path)
	if err != nil {
		panic(err)
	}

	from := dirName[:len(dirName)-6]

	filePath := filepath.Join(path, "Localizable.strings")
	read, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(read), "\n")

	for _, to := range args[1:] {
		to = strings.ToLower(to)
		language, ok := Languages[to]
		if !ok {
			fmt.Printf(color.YellowString("Ignoring '%s' - Invalid ISO code\n"), color.BlueString(to))
			continue
		}
		fmt.Printf(color.HiMagentaString("Localizing: '%s' ('%s')\n"), color.BlueString(language), color.BlueString(to))
		switchLanguage(&to)
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Color("magenta", "bold")
		s.Suffix = color.MagentaString(" Translating...")
		s.Start()
		translatedLines := translate(s, lines, from, to)
		switchLanguage(&to)
		saveDir, err := filepath.Abs(filepath.Dir(absPath))
		if err != nil {
			s.FinalMSG = fmt.Sprintf("❌ Failed to get absolute path of '%s'\n", absPath)
		}
		saveDir = filepath.Join(saveDir, fmt.Sprintf("%s.lproj", to))
		os.MkdirAll(saveDir, 0755)
		saveFile := filepath.Join(saveDir, "Localizable.strings")
		output := strings.Join(translatedLines, "\n")
		err = os.WriteFile(saveFile, []byte(output), 0644)
		if err != nil {
			s.FinalMSG = fmt.Sprintf("❌ Failed to save '%s'\n", saveFile)
		}
		s.FinalMSG = fmt.Sprintf(color.HiGreenString("✔ Successfully localized '%s' ('%s')\n"), color.BlueString(language), color.BlueString(to))
		s.Stop()
	}
}

func translate(s *spinner.Spinner, lines []string, from, to string) []string {
	translatedLines := make([]string, len(lines))
	for i, line := range lines {
		if strings.Contains(line, "=") && strings.Contains(line, ";") {
			keyValuePair := strings.Split(line, "=")
			key, value := strings.TrimSpace(keyValuePair[0]), strings.TrimSpace(keyValuePair[1])
			key, value = key[1:len(key)-1], value[1:len(value)-2]

			s.Suffix = fmt.Sprintf(color.MagentaString(" Translating: '%s'..."), color.BlueString(key))

			translated, err := gtranslate.TranslateWithParams(
				value,
				gtranslate.TranslationParams{
					From: from,
					To:   to,
				},
			)

			if err != nil {
				fmt.Printf("\n❌ Failed to translate '%s' ('%s')\n", key, err)
			}
			translatedLines[i] = fmt.Sprintf("'%s' = '%s';", key, translated)
		}
	}
	return translatedLines
}
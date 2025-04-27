package gen

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func writeStringToFile(fileName, text string) (err error) {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	return
}

func readStringFromFile(fileName string) (string, error) {

	text, err := os.ReadFile(fileName)
	return string(text), err
}

func projectName() (name string, err error) {

	goMod := "./go.mod"

	if !fileExists(goMod) {
		err = errors.New("go.mod is not exists")
		return
	}

	f, err := os.Open(goMod)
	if err != nil {
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		text, _, err := reader.ReadLine()
		if err == io.EOF {
			return "", err
		}

		textArr := strings.Split(string(text), " ")

		return textArr[len(textArr)-1], err
	}

}

func strFirstToUpper(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func strFirstToLower(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}

// get framework path
func FrameworkRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return dir
}

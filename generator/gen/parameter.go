package gen

import (
	"fmt"
	"log"
	"strings"
)

type Parameter struct {
	Name     string
	TempPath string
	ctlName  string
	filename string
}

func (m *Parameter) Make() {
	m.parseName()

	if fileExists(m.filename) {
		log.Printf("[%s] already exists", m.filename)
		return
	}

	tempFile := fmt.Sprintf("%s/parameter", m.TempPath)

	temp, err := readStringFromFile(tempFile)
	if err != nil {
		log.Printf("read [%s] template failed, err: %s", tempFile, err)
		return
	}

	// replace
	fileText := strings.Replace(temp, CtlTempReplaceStr, m.ctlName, -1)

	err = writeStringToFile(m.filename, fileText)
	if err != nil {
		log.Printf("create parameter [%s] template failed, err: %s", m.filename, err)
	}
}

func (m *Parameter) parseName() {

	nameArr := strings.Split(m.Name, "/")
	if len(nameArr) <= 0 {
		log.Printf("[%s] is empty", m.Name)
		return
	}
	if len(nameArr) == 1 {
		m.ctlName = strFirstToUpper(m.Name)
		m.filename = fmt.Sprintf("./%s.go", m.Name)
		return
	}

	m.ctlName = strFirstToUpper(nameArr[len(nameArr)-1])
	m.filename = fmt.Sprintf("./%s/parameter/%s.go", strings.Join(nameArr[:len(nameArr)-1], "/"), nameArr[len(nameArr)-1])
}

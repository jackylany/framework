package gen

import (
	"fmt"
	"log"
	"strings"
)

type Middleware struct {
	Name     string
	TempPath string
	ctlName  string
	filename string
}

func (m *Middleware) Make() {
	m.parseName()

	if fileExists(m.filename) {
		log.Printf("[%s] already exists", m.filename)
		return
	}

	tempFile := fmt.Sprintf("%s/middleware", m.TempPath)

	temp, err := readStringFromFile(tempFile)
	if err != nil {
		log.Printf("read [%s] template failed, err: %s", tempFile, err)
		return
	}

	// replace
	fileText := strings.Replace(temp, CtlTempReplaceStr, m.ctlName, -1)

	err = writeStringToFile(m.filename, fileText)
	if err != nil {
		log.Printf("create middleware [%s] template failed, err: %s", m.filename, err)
	}
}

func (m *Middleware) parseName() {

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
	m.filename = fmt.Sprintf("./%s/middleware/%s.go", strings.Join(nameArr[:len(nameArr)-1], "/"), nameArr[len(nameArr)-1])
}

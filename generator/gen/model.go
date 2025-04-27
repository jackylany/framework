package gen

import (
	"fmt"
	"log"
	"strings"
)

type Model struct {
	Name      string
	TempPath  string
	ctlName   string
	tableName string
	filename  string
}

func (m *Model) Make() {
	m.parseName()

	if fileExists(m.filename) {
		log.Printf("[%s] already exists", m.filename)
		return
	}

	tempFile := fmt.Sprintf("%s/model", m.TempPath)

	temp, err := readStringFromFile(tempFile)
	if err != nil {
		log.Printf("read [%s] template failed, err: %s", tempFile, err)
		return
	}

	// replace
	fileText := strings.Replace(temp, CtlTempReplaceStr, m.ctlName, -1)
	fileText = strings.Replace(fileText, TableTempReplaceStr, m.tableName, -1)

	err = writeStringToFile(m.filename, fileText)
	if err != nil {
		log.Printf("create model [%s] template failed, err: %s", m.filename, err)
	}
}

func (m *Model) parseName() {

	nameArr := strings.Split(m.Name, "/")
	if len(nameArr) <= 0 {
		log.Printf("[%s] is empty", m.Name)
		return
	}

	if len(nameArr) == 1 {
		m.ctlName = strFirstToUpper(m.Name)
		m.tableName = strFirstToLower(m.Name)
		m.filename = fmt.Sprintf("./%s.go", m.Name)
		return
	}

	moduleName := strings.Join(nameArr[:len(nameArr)-1], "/")

	m.tableName = strFirstToLower(nameArr[len(nameArr)-1])
	m.ctlName = strFirstToUpper(m.tableName)
	m.filename = fmt.Sprintf("./%s/model/%s.go", moduleName, m.tableName)

}

package gen

import (
	"fmt"
	"log"
	"strings"
)

type Transform struct {
	Name        string
	TempPath    string
	projectName string
	moduleName  string
	ctlName     string
	modelName   string
	filename    string
}

func (m *Transform) Make() {
	m.parseName()

	if fileExists(m.filename) {
		log.Printf("[%s] already exists", m.filename)
		return
	}

	tempFile := fmt.Sprintf("%s/transformer", m.TempPath)

	temp, err := readStringFromFile(tempFile)
	if err != nil {
		log.Printf("read [%s] template failed, err: %s", tempFile, err)
		return
	}

	// replace
	fileText := strings.Replace(temp, ProjectTempReplaceStr, m.projectName+"/", -1)
	fileText = strings.Replace(fileText, ModuleTempReplaceStr, m.moduleName+"/", -1)
	fileText = strings.Replace(fileText, CtlTempReplaceStr, m.ctlName, -1)
	fileText = strings.Replace(fileText, ModelTempReplaceStr, m.modelName, -1)

	err = writeStringToFile(m.filename, fileText)
	if err != nil {
		log.Printf("create transformer [%s] template failed, err: %s", m.filename, err)
	}
}

func (m *Transform) parseName() {

	pName, err := projectName()
	if err != nil {
		log.Printf("get project name failed err: %s", err)
		return
	}
	m.projectName = pName

	nameArr := strings.Split(m.Name, "/")
	if len(nameArr) <= 0 {
		log.Printf("[%s] is empty", m.Name)
		return
	}
	if len(nameArr) == 1 {
		m.projectName = ""
		m.moduleName = ""
		m.ctlName = strFirstToUpper(m.Name)
		m.modelName = fmt.Sprintf("%ss", m.ctlName)
		m.filename = fmt.Sprintf("./%s.go", m.Name)
		return
	}

	m.moduleName = strings.Join(nameArr[:len(nameArr)-1], "/")
	m.ctlName = strFirstToUpper(nameArr[len(nameArr)-1])
	m.modelName = fmt.Sprintf("%ss", m.ctlName)
	m.filename = fmt.Sprintf("./%s/transformer/%s.go", m.moduleName, nameArr[len(nameArr)-1])
}

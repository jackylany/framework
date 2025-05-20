package gen

import (
	"fmt"
	"log"
	"strings"
)

type Controller struct {
	Name        string
	TempPath    string
	projectName string
	moduleName  string
	ctlName     string
	modelName   string
	filename    string
}

func (m *Controller) Make() {

	m.parseName()

	if fileExists(m.filename) {
		log.Printf("[%s] already exists", m.filename)
		return
	}

	tempFile := fmt.Sprintf("%s/controller", m.TempPath)

	temp, err := readStringFromFile(tempFile)
	if err != nil {
		log.Printf("read [%s] template failed, err: %s", tempFile, err)
		return
	}

	m.projectName += "/"
	if m.moduleName != "" {
		m.moduleName += "/"
	}

	// replace
	fileText := strings.Replace(temp, ProjectTempReplaceStr, m.projectName, -1)
	fileText = strings.Replace(fileText, ModuleTempReplaceStr, m.moduleName, -1)
	fileText = strings.Replace(fileText, CtlTempReplaceStr, m.ctlName, -1)
	fileText = strings.Replace(fileText, ModelTempReplaceStr, m.modelName, -1)

	err = writeStringToFile(m.filename, fileText)
	if err != nil {
		log.Printf("create controller [%s] template failed, err: %s", m.filename, err)
	}
}

func (m *Controller) parseName() {

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

	m.moduleName = strings.Join(nameArr[:len(nameArr)-1], "/")
	m.ctlName = strFirstToUpper(nameArr[len(nameArr)-1])
	m.modelName = fmt.Sprintf("%ss", m.ctlName)
	m.filename = fmt.Sprintf("./%s/controller/%s.go", m.moduleName, nameArr[len(nameArr)-1])
}

package gen

import (
	"fmt"
	"log"
	"strings"
)

type Validator struct {
	Name     string
	TempPath string
	ctlName  string
	filename string
}

func (m *Validator) Make() {

}

func (m *Validator) parseName() {

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

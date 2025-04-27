package generator

type Generator interface {
	Make()
}

func Make(m Generator) {
	m.Make()
}

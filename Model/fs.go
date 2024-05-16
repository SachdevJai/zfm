package Model

import "os"

func (m *Model) OpenSelectedDir() {
	os.Chdir(m.fs[m.Pointer].Name())
	f, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m.fs = nil

	m.fs, err = f.Readdir(-1)
	if err != nil {
		panic(err)
	}

	m.Pointer = 0

	m.CurrentDirectory, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

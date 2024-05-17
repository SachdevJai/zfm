package internal

import "os"

func (m *Model) OpenSelectedDir() {
	os.Chdir(m.fs[m.pointer].Name())
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

	m.pointer = 0

	m.currentDirectory, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func (m *Model) OpenParentDir() {
	os.Chdir("..")
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

	m.pointer = 0

	m.currentDirectory, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

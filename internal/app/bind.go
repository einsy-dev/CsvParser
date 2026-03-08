package app

import "fmt"

type index struct{}

func (b *index) Console(val ...any) {
	fmt.Println(val...)
}

func (b *index) Parse(files []string) {
	fmt.Println(files)
}

func (b *index) Copy() {}

var Bind = &index{}

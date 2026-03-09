package app

import (
	"fmt"

	"golang.design/x/clipboard"
)

type index struct{}

func (b *index) Console(val ...any) {
	fmt.Println(val...)
}

func (b *index) Parse(files []string) {
	fmt.Println(files)
}

func (b *index) Copy(data string) {
	clipboard.Write(clipboard.FmtText, []byte(data))
}

var Bind = &index{}

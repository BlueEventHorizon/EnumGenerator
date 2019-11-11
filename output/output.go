package output

import (
	"fmt"
	"os"

	"../env"
)

type Output struct{}

var fd *os.File
var err error

func (t Output) Open(path string) {

	fd, err = os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	fd.Seek(0, 0)

}

func (t Output) Print(text string) {
	if env.OUTPUT_FILE {
		fd.WriteString(text)
	} else {
		fmt.Print(text)
	}
}

func (t Output) Close() {
	fd.Close()
}

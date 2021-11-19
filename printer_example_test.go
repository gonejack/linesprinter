package linesprinter_test

import (
	"encoding/base64"
	"os"

	"github.com/gonejack/linesprinter"
)

func ExampleNewLinesPrinter() {
	p := linesprinter.NewLinesPrinter(os.Stdout, 76, []byte("\n"))
	defer p.Close()

	w := base64.NewEncoder(base64.StdEncoding, p)
	defer w.Close()

	t := "this is some random string this is some random string this is some random string"
	w.Write([]byte(t))

	// Output:
	// dGhpcyBpcyBzb21lIHJhbmRvbSBzdHJpbmcgdGhpcyBpcyBzb21lIHJhbmRvbSBzdHJpbmcgdGhp
	// cyBpcyBzb21lIHJhbmRvbSBzdHJpbmc=
}

package main

import (
	"io"
	"os"
	"path"
	"strings"
)

// Reads all .html/css/js files in the current folder
// and encodes them as strings literals in webFiles.go
func main() {
	fs, _ := os.ReadDir("./web")
	out, _ := os.Create("webFiles.go")
	out.Write([]byte("package main \n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".html") ||
			strings.HasSuffix(f.Name(), ".css") ||
			strings.HasSuffix(f.Name(), ".js") {
			out.Write([]byte(strings.TrimSuffix(f.Name(), path.Ext(f.Name())) + " = `"))
			c, _ := os.Open("./web/" + f.Name())
			io.Copy(out, c)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}

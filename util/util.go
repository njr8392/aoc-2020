//package of helper functions reading and manipulating data
package util

import(
	"os"
	"log"
	"io"
	)

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, info.Size()-1)
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return buf
}

package main

import (
	"log"
	"os"

	"github.com/xyproto/icostring"
)

func WriteImage(filename, icoString string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := icostring.WriteImage(f, icoString); err != nil {
		return err
	}
	return f.Close()
}

func main() {
	// "aaaaaaaa aaaqqaaa ffqqqqff ffqqqqff aaqqqqaa aaqqqqaa pppqqppp pppppppp #08f" is also a valid icostring
	if err := WriteImage("favicon.ico", "pppppppppppppppp"); err != nil {
		log.Fatalln(err)
	}
}

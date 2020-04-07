package iterm2

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const BEL = "\a"
const ESC = "\033"

func PrintOSC() {
	if os.Getenv("TERM") == "screen" {
		fmt.Printf(ESC + `Ptmux;` + ESC + ESC + `]`)
	} else {
		fmt.Printf(ESC + "]")
	}
}

func PrintST() {
	if os.Getenv("TERM") == "screen" {
		fmt.Printf(BEL + ESC + `\`)
	} else {
		fmt.Printf(BEL)
	}
}

func PrintImage(filename string) error {
	options := make(map[string]string)

	filebuf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	options["name"] = filename
	options["size"] = strconv.Itoa(len(filebuf))
	options["width"] = "auto"
	options["height"] = "auto"
	options["preserveAspectRatio"] = "1"
	options["inline"] = "1"

	PrintOSC()
	fmt.Printf("1337;File=")

	for k, v := range options {
		fmt.Printf("%s=%s;", k, v)
	}

	fmt.Printf(":%s", base64.StdEncoding.EncodeToString(filebuf))

	PrintST()

	return nil
}

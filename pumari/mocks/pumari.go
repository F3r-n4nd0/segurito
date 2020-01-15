package main

import (
	"fmt"
	"os"
	"time"
)

//pumari.exe -u JROCA101 -e ENTRADA -d 10/10/2019 12:23:00

func main() {

	strCmd := ""
	for _, element := range os.Args {
		strCmd += " " + element
	}
	time.Sleep(15 * time.Second)
	fmt.Println(strCmd)

}

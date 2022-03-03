package main

import (
	"fmt"
	"github.com/jcabrera/portainer/validations"
	"io/ioutil"
	"os"
)

func main() {
	file := os.Args[1]

	bts, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}

	valid := validations.ContainerDefinition(string(bts))

	if valid == true {
		fmt.Printf("valid definition file:%s",file)
	} else {
		fmt.Printf("invalid definition file:%s",file)
	}
}

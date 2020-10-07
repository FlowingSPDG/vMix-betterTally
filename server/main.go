package main

import (
	"github.com/sirupsen/logrus"

	vmix "github.com/FlowingSPDG/vmix-go-TCP"
)

var (
	vm = *&vmix.Vmix{}
)

func init() {
	logrus.Debugln("START...")

	var err error // Declare error to avoid "vm" var's shadowing
	vm, err = vmix.New()
	if err != nil {
		panic(err)
	}
}

func main() {

}

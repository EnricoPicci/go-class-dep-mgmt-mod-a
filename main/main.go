package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class-dep-mgmt-mod-b/pkgmodb2"
)

func main() {
	fmt.Println("I am the main program in Module A that calls a function of Module B that calls a function of Module C")
	fmt.Println("The function of Module B calls a function from the external module Module C")
	pkgmodb2.ApiThatCallsModuleCApiThatCallRemainsInteral()
}

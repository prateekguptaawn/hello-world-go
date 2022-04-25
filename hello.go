package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

type HelloWrd struct {
	Msg  string
	Name string
}

func main() {
	helloWrd := HelloWrd{
		Msg:  "Hello",
		Name: "World",
	}
	noVariablePresent := true
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if len(os.Args) > 1 {
		noVariablePresent = false
		helloWrd.Name = os.Args[1]
		sugar.Debugf("Command line argument found : %s", os.Args[1])
	}

	name, present := os.LookupEnv("HELLO_NAME")
	if present {
		noVariablePresent = false
		helloWrd.Name = name
		sugar.Debugf("Enviorment variable HELLO_NAME: %s", name)
	}

	msg, present := os.LookupEnv("HELLO_MSG")
	if present {
		noVariablePresent = false
		helloWrd.Msg = msg
		sugar.Debugf("Enviorment variable HELLO_MSG: %s", msg)
	}

	if noVariablePresent {
		sugar.Debugw("No command line argument or environment variable found!")
	}

	fmt.Printf("%s %s!", helloWrd.Msg, helloWrd.Name)
}

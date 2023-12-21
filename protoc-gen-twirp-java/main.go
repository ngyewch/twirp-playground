package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
)

func main() {
	var flags flag.FlagSet
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			for _, service := range f.Services {
				_, _ = fmt.Fprintf(os.Stderr, "%s\n", service.Desc.Name())
				for _, method := range service.Methods {
					_, _ = fmt.Fprintf(os.Stderr, "    %s\n", method.Desc.Name())
				}
				// TODO
			}
		}
		return nil
	})
}

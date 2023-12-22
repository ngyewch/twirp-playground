package main

import (
	"flag"
	"github.com/ngyewch/protoc-gen-twirp-java/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	g, err := generator.New()
	if err != nil {
		panic(err)
	}

	var flags flag.FlagSet
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(g.Generate)
}

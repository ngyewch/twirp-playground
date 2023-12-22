package generator

import (
	"os"
	"strings"
)

type JavaPackage string

func (jp JavaPackage) Resolve(cn JavaClassName) JavaClassName {
	if jp != "" {
		return JavaClassName(string(jp) + "." + string(cn))
	} else {
		return cn
	}
}

func (jp JavaPackage) Path() string {
	return strings.ReplaceAll(string(jp), ".", string(os.PathSeparator))
}

type JavaClassName string

func (cn JavaClassName) Path() string {
	return strings.ReplaceAll(string(cn), ".", string(os.PathSeparator)) + ".java"
}

package generator

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"path/filepath"
	"strings"
)

type JavaProtoFile struct {
	Package        JavaPackage
	OuterClassName JavaClassName
	MultipleFiles  bool
}

func NewJavaProtoFile(fd protoreflect.FileDescriptor) JavaProtoFile {
	javaPackage := ""
	javaOuterClassName := filenameToJavaClassName(fd.Path())
	javaMultipleFiles := false
	options, ok := fd.Options().(*descriptorpb.FileOptions)
	if ok {
		javaPackage = options.GetJavaPackage()
		javaOuterClassName1 := options.GetJavaOuterClassname()
		if javaOuterClassName1 != "" {
			javaOuterClassName = javaOuterClassName1
		}
		javaMultipleFiles = options.GetJavaMultipleFiles()
	}

	return JavaProtoFile{
		Package:        JavaPackage(javaPackage),
		OuterClassName: JavaClassName(javaOuterClassName),
		MultipleFiles:  javaMultipleFiles,
	}
}

func (jpf JavaProtoFile) FullOuterClassName() JavaClassName {
	return jpf.Package.Resolve(jpf.OuterClassName)
}

func filenameToJavaClassName(path string) string {
	baseName := filepath.Base(path)
	if strings.HasSuffix(baseName, ".proto") {
		return strcase.ToCamel(baseName[0 : len(baseName)-6])
	} else {
		return strcase.ToCamel(baseName)
	}
}

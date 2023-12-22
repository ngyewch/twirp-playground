package generator

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"path/filepath"
	"strings"
)

type FileDescriptorWrapper struct {
	fd                 protoreflect.FileDescriptor
	javaPackage        JavaPackage
	javaOuterClassName JavaClassName
	javaMultipleFiles  bool
}

func WrapFileDescriptor(fd protoreflect.FileDescriptor) FileDescriptorWrapper {
	javaPackage := ""
	javaOuterClassName := filenameToJavaClassName(fd.Path())
	javaMultipleFiles := false
	options, ok := fd.Options().(*descriptorpb.FileOptions)
	if ok {
		javaPackage = options.GetJavaPackage()
		javaOuterClassName1 := options.GetJavaOuterClassname()
		if javaOuterClassName1 != "" {
			javaOuterClassName = JavaClassName(javaOuterClassName1)
		}
		javaMultipleFiles = options.GetJavaMultipleFiles()
	}

	return FileDescriptorWrapper{
		fd:                 fd,
		javaPackage:        JavaPackage(javaPackage),
		javaOuterClassName: javaOuterClassName,
		javaMultipleFiles:  javaMultipleFiles,
	}
}

func (w FileDescriptorWrapper) JavaPackage() JavaPackage {
	return w.javaPackage
}

func (w FileDescriptorWrapper) JavaOuterClassName() JavaClassName {
	return w.javaOuterClassName
}

func (w FileDescriptorWrapper) JavaFullOuterClassName() JavaClassName {
	return w.javaPackage.Resolve(w.javaOuterClassName)
}

func filenameToJavaClassName(path string) JavaClassName {
	baseName := filepath.Base(path)
	if strings.HasSuffix(baseName, ".proto") {
		return JavaClassName(strcase.ToCamel(baseName[0 : len(baseName)-6]))
	} else {
		return JavaClassName(strcase.ToCamel(baseName))
	}
}

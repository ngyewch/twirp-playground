package generator

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
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
			javaOuterClassName = JavaClassName(javaOuterClassName1)
		}
		javaMultipleFiles = options.GetJavaMultipleFiles()
	}

	return JavaProtoFile{
		Package:        JavaPackage(javaPackage),
		OuterClassName: javaOuterClassName,
		MultipleFiles:  javaMultipleFiles,
	}
}

func (jpf JavaProtoFile) FullOuterClassName() JavaClassName {
	return jpf.Package.Resolve(jpf.OuterClassName)
}

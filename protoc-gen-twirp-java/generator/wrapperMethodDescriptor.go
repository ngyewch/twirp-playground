package generator

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MethodDescriptorWrapper struct {
	md         protoreflect.MethodDescriptor
	javaMethod JavaMethod
}

func WrapMethodDescriptor(md protoreflect.MethodDescriptor) MethodDescriptorWrapper {
	return MethodDescriptorWrapper{
		md: md,
		javaMethod: JavaMethod{
			Name:       strcase.ToLowerCamel(string(md.Name())),
			ReturnType: messageDescriptorToJavaType(md.Output()),
			Arguments: []JavaMethodArgument{
				{
					Type: messageDescriptorToJavaType(md.Input()),
					Name: "input",
				},
			},
		},
	}
}

func (w MethodDescriptorWrapper) Name() string {
	return string(w.md.Name())
}

func (w MethodDescriptorWrapper) JavaMethod() JavaMethod {
	return w.javaMethod
}

func messageDescriptorToJavaType(md protoreflect.MessageDescriptor) string {
	fdw := WrapFileDescriptor(md.ParentFile(), false)
	jc := JavaClassName(strcase.ToCamel(string(md.Name())))
	return string(fdw.JavaPackage().Resolve(jc))
}

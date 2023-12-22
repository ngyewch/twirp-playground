package generator

import (
	"bytes"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"log"
	"text/template"
)

type Generator struct {
	templates *template.Template
}

type TemplateData struct {
	ProtoFile          ProtoFileWrapper
	JavaOuterClassName string
	Services           []ServiceDescriptor
}

type ServiceDescriptor struct {
	Name          string
	JavaInterface JavaInterface
	Methods       []MethodDescriptor
}

type MethodDescriptor struct {
	Name       string
	JavaMethod JavaMethod
}

type JavaInterface struct {
	Name    string
	Methods []JavaMethod
}

type JavaMethod struct {
	Name       string
	ReturnType string
	Arguments  []JavaMethodArgument
	Throws     []string
}

type JavaMethodArgument struct {
	Name string
	Type string
}

func New() (*Generator, error) {
	templates, err := template.ParseFS(templateFS, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	return &Generator{
		templates: templates,
	}, nil
}

func (g *Generator) Generate(gen *protogen.Plugin) error {
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

		log.Printf("processing %s", f.Desc.Path())

		if len(f.Services) == 0 {
			continue
		}

		pfw := NewProtoFileWrapper(f.Desc)

		templateData := TemplateData{
			ProtoFile: pfw,
		}
		templateData.JavaOuterClassName = string(pfw.JavaOuterClassName()) + "Twirp"
		for _, service := range f.Services {
			templateData.Services = append(templateData.Services, toServiceDescriptor(service))
		}

		outputBuffer := bytes.NewBuffer(nil)
		err := g.templates.ExecuteTemplate(outputBuffer, "Twirp.java.tmpl", templateData)
		if err != nil {
			return err
		}

		fullOuterClassName := JavaClassName(string(pfw.JavaFullOuterClassName()) + "Twirp")
		generatedFile := gen.NewGeneratedFile(fullOuterClassName.Path(), f.GoImportPath)
		_, err = generatedFile.Write(outputBuffer.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func toServiceDescriptor(service *protogen.Service) ServiceDescriptor {
	serviceDescriptor := ServiceDescriptor{
		Name: string(service.Desc.Name()),
		JavaInterface: JavaInterface{
			Name: strcase.ToCamel(string(service.Desc.Name())),
		},
	}
	for _, method := range service.Methods {
		serviceDescriptor.Methods = append(serviceDescriptor.Methods, MethodDescriptor{
			Name: string(method.Desc.Name()),
			JavaMethod: JavaMethod{
				ReturnType: convertMessageToJavaType(method.Output),
				Name:       strcase.ToLowerCamel(string(method.Desc.Name())),
				Arguments: []JavaMethodArgument{
					{
						Type: convertMessageToJavaType(method.Input),
						Name: "input",
					},
				},
			},
		})
	}
	return serviceDescriptor
}

func convertMessageToJavaType(message *protogen.Message) string {
	pfw := NewProtoFileWrapper(message.Desc.ParentFile())
	jc := JavaClassName(strcase.ToCamel(string(message.Desc.Name())))
	return string(pfw.JavaPackage().Resolve(jc))
}

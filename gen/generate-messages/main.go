package main

import (
	"flag"
	"fmt"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
	"os"
	"strconv"
	"strings"
)

var (
	pkg     string
	fixSpec *datadictionary.DataDictionary
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-messages [flags] <path to data dictionary>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func initPackage() {
	pkg = strings.ToLower(fixSpec.FIXType) + strconv.Itoa(fixSpec.Major) + strconv.Itoa(fixSpec.Minor)

	if fixSpec.ServicePack != 0 {
		pkg += "sp" + strconv.Itoa(fixSpec.ServicePack)
	}
}

func genMessages() {
	for _, m := range fixSpec.Messages {
		genMessage(m)
	}
}

func genCracker() {
	fileOut := fmt.Sprintf("package %v\n", pkg)
	fileOut += buildCrackerImports()
	fileOut += buildCrack()
	fileOut += buildMessageRouter()
	fileOut += buildMessageCracker()

	filePath := pkg + "/message_cracker.go"
	gen.WriteFile(filePath, fileOut)
}

func buildCrackerImports() string {
	return `
import(
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/message"
	"github.com/quickfixgo/quickfix/fix/field"
)
`
}

func buildCrack() (out string) {
	out += "func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) message.MessageReject {\n"
	out += `
  msgType:=new(field.MsgType)
switch msg.Header.Get(msgType); msgType.Value {
`

	for msgType, m := range fixSpec.Messages {
		out += fmt.Sprintf("case \"%v\":\n", msgType)
		out += fmt.Sprintf("return router.On%v%v(%v{msg},sessionID)\n", strings.ToUpper(pkg), m.Name, m.Name)
	}
	out += "}\n"
	out += "return message.NewInvalidMessageType(msg)\n"
	out += "}\n"

	return
}

func buildMessageRouter() (out string) {
	out += "type MessageRouter interface {\n"

	for _, m := range fixSpec.Messages {
		out += fmt.Sprintf("On%v%v(msg %v, sessionID quickfix.SessionID) message.MessageReject\n", strings.ToUpper(pkg), m.Name, m.Name)
	}

	out += "}\n"

	return
}

func buildMessageCracker() (out string) {
	out += fmt.Sprintf("type %vMessageCracker struct {}\n", strings.ToUpper(pkg))

	for _, m := range fixSpec.Messages {
		out += fmt.Sprintf("func (c * %vMessageCracker) On%v%v(msg %v, sessionId quickfix.SessionID) message.MessageReject {\n", strings.ToUpper(pkg), strings.ToUpper(pkg), m.Name, m.Name)
		out += "return message.NewUnsupportedMessageType(msg.Message)\n}\n"
	}

	return
}

func genMessage(msg *datadictionary.MessageDef) {
	fileOut := fmt.Sprintf("package %v\n", pkg)
	fileOut += `
import( 
  "github.com/quickfixgo/quickfix/message"
  "github.com/quickfixgo/quickfix/fix/field"
)
`
	fileOut += fmt.Sprintf("//%v msg type = %v.\n", msg.Name, msg.MsgType)
	fileOut += fmt.Sprintf("type %v struct {\n message.Message}\n", msg.Name)

	fileOut += fmt.Sprintf("//%vBuilder builds %v messages.\n", msg.Name, msg.Name)
	fileOut += fmt.Sprintf("type %vBuilder struct {\n message.MessageBuilder}\n", msg.Name)

	requiredFields := make([]*datadictionary.FieldDef, 0, len(msg.FieldsInDeclarationOrder))
	for _, field := range msg.FieldsInDeclarationOrder {
		if field.Required {
			requiredFields = append(requiredFields, field)
		}
	}

	fileOut += fmt.Sprintf("//New%vBuilder returns an initialized %vBuilder with specified required fields.\n", msg.Name, msg.Name)
	fileOut += fmt.Sprintf("func New%vBuilder(\n", msg.Name)
	builderArgs := make([]string, len(requiredFields))
	for i, field := range requiredFields {
		builderArgs[i] = fmt.Sprintf("%v field.%v", strings.ToLower(field.Name), field.Name)
	}
	fileOut += strings.Join(builderArgs, ",\n")
	fileOut += fmt.Sprintf(") *%vBuilder {\n", msg.Name)
	fileOut += fmt.Sprintf("builder:=new(%vBuilder)\n", msg.Name)

	for _, field := range requiredFields {
		fileOut += fmt.Sprintf("builder.Body.Set(%v)\n", strings.ToLower(field.Name))
	}

	fileOut += "return builder\n"
	fileOut += "}\n"

	for _, field := range msg.FieldsInDeclarationOrder {
		if field.Required {
			fileOut += fmt.Sprintf("//%v is a required field for %v.\n", field.Name, msg.Name)
		} else {
			fileOut += fmt.Sprintf("//%v is a non-required field for %v.\n", field.Name, msg.Name)
		}
		fileOut += fmt.Sprintf("func (m * %v) %v() (*field.%v, error) {\n", msg.Name, field.Name, field.Name)
		fileOut += fmt.Sprintf("f:=new(field.%v)\n", field.Name)
		fileOut += "err:=m.Body.Get(f)\n"
		fileOut += "return f, err\n}\n"
	}

	filePath := pkg + "/" + msg.Name + ".go"
	gen.WriteFile(filePath, fileOut)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
	}

	dataDict := flag.Arg(0)

	if spec, err := datadictionary.Parse(dataDict); err != nil {
		panic(err)
	} else {
		fixSpec = spec
	}

	initPackage()

	if fi, err := os.Stat(pkg); os.IsNotExist(err) {
		if err := os.Mkdir(pkg, os.ModePerm); err != nil {
			panic(err)
		}
	} else if !fi.IsDir() {
		panic(pkg + "/ is not a directory")
	}

	genCracker()
	genMessages()
}

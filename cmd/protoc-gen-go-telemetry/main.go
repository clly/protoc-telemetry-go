package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protodesc"

	"github.com/clly/proto-telemetry/cmd/pkg/generators"
	fields "github.com/clly/proto-telemetry/cmd/pkg/generators"
	"github.com/clly/proto-telemetry/cmd/pkg/logger"
	"github.com/clly/proto-telemetry/cmd/pkg/options"
)

type config struct {
	includeMap       bool
	logLevel         int
	telemetryBackend string
}

func main() {
	var flags flag.FlagSet
	includeMap := flags.Bool("include-map", false, "include map key/values in trace span")
	logLevel := flags.Int("loglevel", 0, "Set the log level. Higher numbers add more logging, Tops out at 3")
	telemetryBackend := flags.String("telemetry-backend", "opentelemetry", "Telemetry implementation to use. Supports opentelemetry or opencensus")
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}

	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		flags.PrintDefaults()
		os.Exit(0)
	}

	opts.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if !f.Generate {
				continue
			}
			cfg := config{
				includeMap:       *includeMap,
				logLevel:         *logLevel,
				telemetryBackend: *telemetryBackend,
			}

			excludeFile := options.GetTelemetryExcludeFile(protodesc.ToFileDescriptorProto(f.Desc))
			if excludeFile {
				fmt.Fprintln(os.Stderr, "skipping file", f.Desc.Path())
				continue
			}
			generateFile(p, f, cfg)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, f *protogen.File, cfg config) {
	filename := f.GeneratedFilenamePrefix + "_telemetry.pb.go"

	g := gen.NewGeneratedFile(filename, f.GoImportPath)
	g.P("// Code generated by protoc-gen-go-telemetry version:", version(), ". DO NOT EDIT.")
	g.P()
	g.P("package ", f.GoPackageName)
	g.P()

	setLogger(cfg.logLevel)

	var fileGen *fields.FileGenerator
	switch cfg.telemetryBackend {
	case "opentelemetry":
		fileGen = fields.NewFileGenerator(g, &generators.OpentelemetryGenerator{})
	case "opencensus":
		fileGen = fields.NewFileGenerator(g, &generators.OpencensusGenerator{})

	}

	fileGen.Generate()

	msgs := collectMessages(f.Messages, fileGen.Telemetry)

	for _, msgGenerator := range msgs {
		msg := msgGenerator.Proto()
		if msg.GoIdent.GoImportPath != f.GoImportPath {
			debugLog(msg.GoIdent.String(), "is unsupported. GoImportPath does not match")
			continue
		}
		if msg.Desc.IsMapEntry() {
			debugLog("Skipping map entry", msg.GoIdent.GoName)
			// we use fmt to format the map keys so if we have a map import it
			continue
		}
		debugLog("generating fields for messages", msg.GoIdent.GoName)
		msgGenerator.Generate(fileGen, false)

		msgGenerator.Tail(g)

		msgGenerator.Generate(fileGen, true)
		msgGenerator.Tail(g)
	}
}

func collectMessages(msgs []*protogen.Message, t fields.TelemetryBackend) []fields.Message {
	messages := make([]fields.Message, 0, len(msgs))
	for _, m := range msgs {
		fieldMessage := fields.MessageGenerator(m, t)
		messages = append(messages, fieldMessage)
		debugLog("found message", m.GoIdent.GoName, "for generation")

		for _, child := range fieldMessage.Children() {
			messages = append(messages, child)
			debugLog("found child message", m.GoIdent.GoName, "for generation")
		}
	}
	return messages
}

func setLogger(i int) {
	var l logger.LogLevel
	if i < 0 {
		l = logger.Error
	} else if i > 3 {
		l = logger.Debug
	}
	switch i {
	case 0:
		l = logger.Error
	case 1:
		l = logger.Warn
	case 2:
		l = logger.Info
	case 3:
		l = logger.Debug
	default:
	}

	logger.SetLevel(l)
}

// version retrieves the current tag from the golang vcs info on build
func version() string {
	v := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		v = info.Main.Version
	}
	return v
}

func debugLog(s ...string) {
	logger.Default().Debug(s...)
}

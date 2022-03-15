package config

import (
	"bytes"
	"io"
	"os"

	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"gopkg.in/yaml.v3"
)

var globalAllowedTypes = []string{
	".tar",
	".tar.gz",
	".tgz",
	".zip",
}

// FileOptions represents the options set in the yaml file.
//
// TODO: this type is badly named, it should be command config or something.
type FileOptions interface {

	// Commands returns the command configured in the config.yml file.
	Commands() Command

	// FileTypes returns the configured allowed file types for the handler
	// command.
	FileTypes() []string

	// ServiceName returns the name of the configured service.
	ServiceName() string
}

func ParseFileOptions(path string) (FileOptions, error) {
	log.Logger().Trace("config.ParseFileOptions")
	file, err := os.Open(path)
	if err != nil {
		log.Logger().Debug("failed to open config file")
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if b, err := io.Copy(buf, file); err != nil {
		log.Logger().Debug("failed to read config file")
		return nil, err
	} else {
		log.Logger().Debug("copied", b, "bytes into buffer")
	}

	return ParseOptionsReader(buf)
}

func ParseOptionsReader(reader io.Reader) (FileOptions, error) {
	log.Logger().Trace("config.ParseOptionsReader")

	out := new(fileOptions)
	dec := yaml.NewDecoder(reader)
	if err := dec.Decode(out); err != nil {
		log.Logger().Debug("failed to decode byte buffer")
		return nil, err
	}

	out.Extensions = appendDefaultFileTypes(out.Extensions)

	return out, nil
}

func appendDefaultFileTypes(types []string) []string {
	log.Logger().Trace("config.appendDefaultFileTypes")
	return append(types, globalAllowedTypes...)
}

type fileOptions struct {
	Command    Command  `yaml:"command"`
	Extensions []string `yaml:"file-types"`
	SvcName    string   `yaml:"service-name"`
}

func (f *fileOptions) Commands() Command {
	return f.Command
}

func (f *fileOptions) FileTypes() []string {
	return f.Extensions
}

func (f *fileOptions) ServiceName() string {
	return f.SvcName
}

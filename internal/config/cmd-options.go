package config

import (
	"bytes"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

var globalAllowedTypes = []string{
	".tar",
	".tar.gz",
	".tgz",
	".zip",
}

type FileOptions interface {
	Commands() Command
	FileTypes() []string
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
	if _, err := io.Copy(file, buf); err != nil {
		log.Logger().Debug("failed to read config file")
		return nil, err
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

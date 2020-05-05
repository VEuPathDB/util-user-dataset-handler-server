package command

import (
	"encoding/json"
	"errors"
	"github.com/VEuPathDB/util-exporter-server/internal/dataset"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	dsFilePerm = 0666
	datasetFile = "dataset.json"
	metaFile = "meta.json"
	errWriteMeta = "Failed to write dataset meta json file: "
)

func (r *runner) generateMetaFiles(files []string) ([]string, error) {
	meta := r.getMeta()
	size, err := r.getFileSize(files)

	if err != nil {
		return nil, err
	}

	totalSize := util.SumU64(size...)

	ds := dataset.Info{
		Owner:        strconv.FormatUint(uint64(r.details.UserID), 10),
		Projects:     r.details.Projects,
		Type:         meta.Type,
		Dependencies: meta.Dependencies,
		Created:      uint64(time.Now().UTC().Unix()),
		Size:         totalSize,
		DataFiles:    makeFileEntries(files, size),
	}

	sum := dataset.Metadata{
		Name:        meta.Name,
		Description: meta.Description,
		Summary:     meta.Summary,
	}

	// Write dataset.json
	if err := writeMetaFile(r.details.WorkingDir, datasetFile, &ds); err != nil {
		return nil, err
	}

	// Write meta.json
	if err := writeMetaFile(r.details.WorkingDir, metaFile, &sum); err != nil {
		return nil, err
	}

	r.details.Size = uint(totalSize)
	r.storeDetails()

	return append(files, datasetFile, metaFile), nil
}

func writeMetaFile(dir, file string, data interface{}) error {
	open, err := openMetaFile(path.Join(dir, file))
	if err != nil {
		return err
	}
	return writeMetaJson(open, data)
}

func writeMetaJson(file *os.File, data interface{}) error {
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err := enc.Encode(data)
	if err != nil {
		return errors.New(errWriteMeta + err.Error())
	}
	return nil
}

func openMetaFile(file string) (*os.File, error) {
	o, e := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, dsFilePerm)
	if e != nil {
		return nil, errors.New(errWriteMeta + e.Error())
	}
	return o, nil
}

func makeFileEntries(files []string, sizes []uint64) []dataset.File {
	out := make([]dataset.File, len(files))
	for i := range files {
		out[i] = dataset.File{
			File: files[i],
			Name: files[i],
			Size: sizes[i],
		}
	}
	return out
}
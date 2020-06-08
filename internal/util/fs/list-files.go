package fs

import "io/ioutil"

func ListFiles(dir string) ([]string, error) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	out := make([]string, len(entries))

	for i, ent := range entries {
		out[i] = ent.Name()
	}

	return out, nil
}

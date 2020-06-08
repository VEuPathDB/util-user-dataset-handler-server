package slice

func ContainsString(slc []string, tgt string) bool {
	for i := range slc {
		if slc[i] == tgt {
			return true
		}
	}

	return false
}

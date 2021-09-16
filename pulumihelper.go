package pulumihelper

func Mapgetkeys(strings map[string]string) []string {
	keys := make([]string, 0, len(strings))

	for k := range strings {
		keys = append(keys, k)
	}
	return keys
}

func Mapgetvalues(strings map[string]string) []string {

	values := make([]string, 0, len(strings))

	for _, v := range strings {
		values = append(values, v)
	}
	return values
}

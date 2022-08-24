package helpers

func GetKeys(m map[string]interface{}) []string {
	/*
		Get keys from a map
	*/

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
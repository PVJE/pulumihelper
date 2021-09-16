package pulumihelper

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func GetMapKeys(strings map[string]string) []string {
	keys := make([]string, 0, len(strings))

	for k := range strings {
		keys = append(keys, k)
	}
	return keys
}

func GetMapValues(strings map[string]string) []string {

	values := make([]string, 0, len(strings))

	for _, v := range strings {
		values = append(values, v)
	}
	return values
}

func ToPulumiStringMap(tags_key map[string]string, tags_values map[string]string) pulumi.StringMap {
	tagsMap := make(pulumi.StringMap)
	for k := range tags_key {
		tagsMap[tags_key[k]] = pulumi.String(tags_values[k])
	}
	return tagsMap
}

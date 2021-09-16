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

func ToPulumiStringMap(maps ...map[string]string) pulumi.StringMap {
	tagsMap := make(pulumi.StringMap)

	for k := range maps[0] {
		tagsMap[maps[1][k]] = pulumi.String(maps[0][k])
	}
	return tagsMap
}

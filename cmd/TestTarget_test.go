package cmd

import "fmt"

type testTarget struct {
	loadParam [][]interface{}
}

func (target *testTarget) Load(path1, path2 string) string {
	target.loadParam = append(target.loadParam, []interface{}{path1, path2})

	return fmt.Sprintf(`Load("%s", "%s")`, path1, path2)
}

package cmd

import "fmt"

type testTarget struct {
	loadParam [][]interface{}
	infoParam [][]interface{}
}

func (target *testTarget) Load(path1, path2 string) string {
	target.loadParam = append(target.loadParam, []interface{}{path1, path2})

	return fmt.Sprintf(`Load("%s", "%s")`, path1, path2)
}

func (target *testTarget) Info() string {
	target.infoParam = append(target.infoParam, []interface{}{})

	return fmt.Sprintf(`Info()`)
}
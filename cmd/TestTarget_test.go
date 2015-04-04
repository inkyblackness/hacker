package cmd

import "fmt"

type testTarget struct {
	loadParam [][]interface{}
	infoParam [][]interface{}
	cdParam   [][]interface{}
	dumpParam [][]interface{}
}

func (target *testTarget) Load(path1, path2 string) string {
	target.loadParam = append(target.loadParam, []interface{}{path1, path2})

	return fmt.Sprintf(`Load("%s", "%s")`, path1, path2)
}

func (target *testTarget) Info() string {
	target.infoParam = append(target.infoParam, []interface{}{})

	return fmt.Sprintf(`Info()`)
}

func (target *testTarget) ChangeDirectory(path string) string {
	target.cdParam = append(target.cdParam, []interface{}{path})

	return fmt.Sprintf(`Cd(%s)`, path)
}

func (target *testTarget) Dump() string {
	target.dumpParam = append(target.dumpParam, []interface{}{})

	return fmt.Sprintf(`Dump()`)
}

func (target *testTarget) Diff(source string) string {

	return fmt.Sprintf(`Diff()`)
}

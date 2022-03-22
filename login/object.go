
package main

import "os"
import "strings"

type Object struct {
	Index Index
	ObjectName string
}

func (o Object) FullName() (string) {
	return o.Index.Root + "/" + o.Index.IndexName + "/" + o.ObjectName
}

func (o Object) WriteMap(m map[string]string) (error) {
	os.MkdirAll(o.FullName(), 0666)

	for filename, content := range m {
		fileError := os.WriteFile(o.FullName() + "/" + filename, []byte(content), 0777)
		if fileError != nil { return fileError }
	}

	return nil
}

func (o Object) WriteArrayMap(m map[string][]string) (error) {
	n := make(map[string]string, len(m))

	for key := range m {
		n[key] = string(strings.Join(m[key], "\n"))
	}

	return o.WriteMap(n)
}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/runtimes"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
)

func main() {
	var log Logger
	var rt runtimes.NervRuntime
	rt.Logger = &log
	rt.Init()
	testdata, _ := json.Marshal(2)
	result, err := rt.RunFunction("./testdata/js/testfn.js", schema.RuntimeTypeJS, testdata)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

package functional

import (
	"time"

	// v8 "github.com/Nerv-Core-Developers/nerv-provider-golang/deplib/v8/golang"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	v8 "github.com/augustoroman/v8"
)

func JSExec(funcScript string, data string) (*shared.RuntimeOutput, time.Duration, error) {
	start := time.Now()
	ctx := v8.NewIsolate().NewContext()
	val, err := ctx.Create(map[string]interface{}{
		"data": data,
	})
	if err != nil {
		return nil, time.Since(start), err
	}
	err = ctx.Global().Set("fnData", val)
	if err != nil {
		return nil, time.Since(start), err
	}
	res, err := ctx.Eval(funcScript, `fn.js`)
	if err != nil {
		return nil, time.Since(start), err
	}
	fOutput, err := readJSOutput(res)

	if err != nil {
		return nil, time.Since(start), err
	}
	return fOutput, time.Since(start), nil
}

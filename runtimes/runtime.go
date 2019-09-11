package runtimes

import (
	"errors"
	"io/ioutil"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/runtimes/functional"
)

func (rt *NervRuntime) Init() error {
	rt.Logger.Init("runtimes")
	rt.Logger.Log("Init runtimes...")
	return nil
}

func (rt *NervRuntime) RunFunction(fnLink string, fnType schema.RuntimeType, data []byte) (*shared.RuntimeOutput, error) {
	// fnInfo, err := rt.Params.Controller.GetFunctionInfo(fnHash)
	// if err != nil {
	// 	return nil, err
	// }
	switch fnType {
	case schema.RuntimeTypeJS: // js
		dat, err := ioutil.ReadFile(fnLink)
		if err != nil {
			return nil, errors.New("Error when reading function. Detail: " + err.Error())
		}
		output, _, err := functional.JSExec(string(dat), string(data))
		if err != nil {
			return nil, errors.New("Error when executing function. Detail: " + err.Error())
		}
		return output, nil
	case schema.RuntimeTypeWebASM: // webasm
		return nil, errors.New("Not supported yet")
	case schema.RuntimeTypeContainer: // container
		return nil, errors.New("Not supported yet")
	default:
		return nil, errors.New("Unknown function type")
	}
}

func (rt *NervRuntime) RunSequence() error {
	return nil
}

func (rt *NervRuntime) StartService() error {
	return nil
}

func (rt *NervRuntime) Version() int {
	return Version
}

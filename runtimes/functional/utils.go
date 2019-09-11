package functional

import (
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
	v8 "github.com/augustoroman/v8"
)

func readOutput(out []byte) (*schema.RuntimeOutput, error) {
	var output *schema.RuntimeOutput
	// in := bufio.NewScanner(bytes.NewReader(out))
	// for in.Scan() {
	// 	outLine := strings.Split(in.Text(), " | ")
	// 	switch outLine[0] {
	// 	case "LOG":
	// 		output.LogData = append(output.LogData, outLine[1])
	// 	case "OUTPUT":
	// 		output.Status = 200
	// 		output.Result = outLine[1]
	// 		return output, nil
	// 	}
	// }
	// if err := in.Err(); err != nil {
	// 	return output, err
	// }
	return output, nil
}

func readJSOutput(out *v8.Value) (*shared.RuntimeOutput, error) {
	var output *shared.RuntimeOutput
	// logs, err := out.Get("logs")
	// if err != nil {
	// 	return nil, err
	// }
	result, err := out.Get("result")
	if err != nil {
		return nil, err
	}
	status, err := out.Get("status")
	if err != nil {
		return nil, err
	}
	// builder := flatbuffers.NewBuilder(0)
	// resultBytes := builder.CreateByteVector(result.Bytes())
	// logsBytes := builder.CreateByteVector(logs.Bytes())

	// schema.RuntimeOutputStart(builder)
	// schema.RuntimeOutputAddResult(builder, resultBytes)
	// schema.RuntimeOutputAddLog(builder, logsBytes)
	// schema.RuntimeOutputAddStatus(builder, status.Bool())
	// outputBytes := schema.RuntimeOutputEnd(builder)

	// builder.Finish(outputBytes)
	// output := builder.FinishedBytes()
	output.Status = status.Bool()
	output.Result = result.String()
	return output, nil
}

package functional

import (
	"github.com/naokichau/nerv-provider-golang/shared/schema"
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

func readJSOutput(out string) (*schema.RuntimeOutput, error) {
	var output *schema.RuntimeOutput
	return output, nil
}

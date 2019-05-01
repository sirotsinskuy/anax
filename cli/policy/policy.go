package policy

import (
	"encoding/json"
	"fmt"
	"github.com/open-horizon/anax/cli/cliconfig"
	"github.com/open-horizon/anax/cli/cliutils"
	"github.com/open-horizon/anax/externalpolicy"
	"net/http"
)

func List() {
	// Get the node policy info
	nodePolicy := externalpolicy.ExternalPolicy{}
	cliutils.HorizonGet("node/policy", []int{200}, &nodePolicy, false)

	// Output the combined info
	jsonBytes, err := json.MarshalIndent(nodePolicy, "", cliutils.JSON_INDENT)
	if err != nil {
		cliutils.Fatal(cliutils.JSON_PARSING_ERROR, "failed to marshal 'hzn policy list' output: %v", err)
	}
	fmt.Printf("%s\n", jsonBytes)
}

func Update(fileName string) {

	ep := new(externalpolicy.ExternalPolicy)
	readInputFile(fileName, ep)

	_, _ = cliutils.HorizonPutPost(http.MethodPost, "node/policy", []int{201, 200}, ep)

	fmt.Println("Horizon node policy updated.")

}

func readInputFile(filePath string, inputFileStruct *externalpolicy.ExternalPolicy) {
	newBytes := cliconfig.ReadJsonFileWithLocalConfig(filePath)
	err := json.Unmarshal(newBytes, inputFileStruct)
	if err != nil {
		cliutils.Fatal(cliutils.JSON_PARSING_ERROR, "failed to unmarshal json input file %s: %v", filePath, err)
	}
}
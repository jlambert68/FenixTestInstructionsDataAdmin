package LocalExecutionMethods

import (
	"FenixTestInstructionsDataAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
)

// MethodsForLocalExecutionsStruct
// Struct for holding all information of how to execute a TestInstruction
// Even when there are no information about the methods this struct must exist
type MethodsForLocalExecutionsStruct struct {
	FangEngineClassesMethodsAttributes *FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct `json:"FangEngineClassesMethodsAttributes"`
}

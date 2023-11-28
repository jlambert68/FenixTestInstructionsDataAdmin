package TestInstructionAndTestInstuctionContainerTypes

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"time"
)

// MessageTypes used when defining the TestInstructionsMap and TestInstructionContainersMap themselves

// TestInstructionStruct
// Struct for holding all data for a TestInstruction
type TestInstructionStruct struct {
	TestInstruction                    *TypeAndStructs.TestInstructionStruct                                 `json:"TestInstruction"`
	BasicTestInstructionInformation    *TypeAndStructs.BasicTestInstructionInformationStruct                 `json:"BasicTestInstructionInformation"`
	ImmatureTestInstructionInformation []*TypeAndStructs.ImmatureTestInstructionInformationStruct            `json:"ImmatureTestInstructionInformation"`
	TestInstructionAttribute           []*TypeAndStructs.TestInstructionAttributeStruct                      `json:"TestInstructionAttribute"`
	ImmatureElementModel               []*TypeAndStructs.ImmatureElementModelMessageStruct                   `json:"ImmatureElementModel"`
	FangEngineClassesMethodsAttributes *FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct `json:"FangEngineClassesMethodsAttributes"`
}

// TestInstructionContainerStruct
// Struct for holding all data for a TestInstructionContainer
type TestInstructionContainerStruct struct {
	TestInstructionContainer                 *TypeAndStructs.TestInstructionContainerStruct                  `json:"TestInstructionContainer"`
	BasicTestInstructionContainerInformation *TypeAndStructs.BasicTestInstructionContainerInformationStruct  `json:"BasicTestInstructionContainerInformation"`
	ImmatureTestInstructionContainer         []*TypeAndStructs.ImmatureTestInstructionContainerMessageStruct `json:"ImmatureTestInstructionContainer"`
	ImmatureElementModel                     []*TypeAndStructs.ImmatureElementModelMessageStruct             `json:"ImmatureElementModel"`
}

// MessageTypes used when sending available TestInstructionsMap and TestInstructionContainersMap to Fenix backend

// TestInstructionInstanceVersionStruct
// Struct for one TestInstruction, to be sent over gRPC to Fenix backend
type TestInstructionInstanceVersionStruct struct {
	TestInstructionInstance             *TestInstructionStruct `json:"TestInstructionInstance"`
	TestInstructionInstanceMajorVersion int                    `json:"TestInstructionInstanceMajorVersion"`
	TestInstructionInstanceMinorVersion int                    `json:"TestInstructionInstanceMinorVersion"`
	Deprecated                          bool                   `json:"Deprecated"`
	Enabled                             bool                   `json:"Enabled"`
	TestInstructionInstanceVersionHash  string                 `json:"TestInstructionInstanceVersionHash"`
}

// TestInstructionInstanceVersionsStruct
// Struct for all versions of one TestInstruction, to be sent over gRPC to Fenix backend
type TestInstructionInstanceVersionsStruct struct {
	TestInstructionVersions     []*TestInstructionInstanceVersionStruct `json:"TestInstructionVersions"`     // Last version is first in slice
	TestInstructionVersionsHash string                                  `json:"TestInstructionVersionsHash"` // SHA256 of all TestInstructionVersions.TestInstructionInstanceVersionHash using Fenix standard way of hashing values together
}

// TestInstructionsStruct
// Struct for all TestInstructionsMap, to be sent over gPRC to Fenix backend
type TestInstructionsStruct struct {
	TestInstructionsMap  map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionInstanceVersionsStruct `json:"TestInstructionsMap"`
	TestInstructionsHash string                                                                            `json:"TestInstructionsHash"` // SHA256 of all TestInstructionsMap.TestInstructionVersionsHash using Fenix standard way of hashing values together
}

// TestInstructionContainerInstanceVersionStruct
// Struct for one TestInstructionContainer, to be sent over gRPC to Fenix backend
type TestInstructionContainerInstanceVersionStruct struct {
	TestInstructionContainerInstance             *TestInstructionContainerStruct `json:"TestInstructionContainerInstance"`
	TestInstructionContainerInstanceMajorVersion int                             `json:"TestInstructionContainerInstanceMajorVersion"`
	TestInstructionContainerInstanceMinorVersion int                             `json:"TestInstructionContainerInstanceMinorVersion"`
	Deprecated                                   bool                            `json:"Deprecated"`
	Enabled                                      bool                            `json:"Enabled"`
	TestInstructionContainerInstanceHash         string                          `json:"TestInstructionContainerInstanceHash"`
}

// TestInstructionContainerInstanceVersionsStruct
// Struct for all versions of one TestInstructionContainer, to be sent over gRPC to Fenix backend
type TestInstructionContainerInstanceVersionsStruct struct {
	TestInstructionContainerVersions     []*TestInstructionContainerInstanceVersionStruct `json:"TestInstructionContainerVersions"`     // Last version is first in slice
	TestInstructionContainerVersionsHash string                                           `json:"TestInstructionContainerVersionsHash"` // SHA256 of all TestInstructionContainerVersions.TestInstructionContainerInstanceHash using Fenix standard way of hashing values together
}

// TestInstructionContainersStruct
// Struct for all TestInstructionContainersMap, to be sent over gPRC to Fenix backend
type TestInstructionContainersStruct struct {
	TestInstructionContainersMap  map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionContainerInstanceVersionsStruct `json:"TestInstructionContainersMap"`
	TestInstructionContainersHash string                                                                                     `json:"xxxTestInstructionContainersHashxxx"` // SHA256 of all TestInstructionContainersMap.TestInstructionContainerVersionsHash using Fenix standard way of hashing values together
}

// TestInstructionsAndTestInstructionsContainersStruct
// Struct for all TestInstructions and TestInstructionsContainers from a "System" that should be sent to Fenix backen
type TestInstructionsAndTestInstructionsContainersStruct struct {
	TestInstructions                                         *TestInstructionsStruct          `json:"TestInstructions"`
	TestInstructionContainers                                *TestInstructionContainersStruct `json:"TestInstructionContainers"`
	MessageCreationTimeStamp                                 time.Time                        `json:"MessageCreationTimeStamp"`
	TestInstructionsAndTestInstructionsContainersMessageHash string                           `json:"testinstructionsandtestinstructionscontainersmessagehash"` // SHA256(TestInstructionsHash concat TestInstructionContainersHash)
}

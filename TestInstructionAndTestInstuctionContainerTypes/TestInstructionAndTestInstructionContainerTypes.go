package TestInstructionAndTestInstuctionContainerTypes

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"time"
)

// MessageTypes used when defining the TestInstructions and TestInstructionContainers themselves

// TestInstructionStruct
// Struct for holding all data for a TestInstruction
type TestInstructionStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}

// TestInstructionContainerStruct
// Struct for holding all data for a TestInstructionContainer
type TestInstructionContainerStruct struct {
	TestInstructionContainer                 TypeAndStructs.TestInstructionContainerStruct
	BasicTestInstructionContainerInformation TypeAndStructs.BasicTestInstructionContainerInformationStruct
	ImmatureTestInstructionContainer         []TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
	ImmatureElementModel                     []TypeAndStructs.ImmatureElementModelMessageStruct
}

// MessageTypes used when sending available TestInstructions and TestInstructionContainers to Fenix backend

// TestInstructionInstanceVersionStruct
// Struct for one TestInstruction, to be sent over gRPC to Fenix backend
type TestInstructionInstanceVersionStruct struct {
	TestInstructionInstance             TestInstructionStruct
	TestInstructionInstanceMajorVersion int
	TestInstructionInstanceMinorVersion int
	Deprecated                          bool
	Enabled                             bool
	TestInstructionInstanceHash         string
}

// TestInstructionInstanceVersionsStruct
// Struct for all versions of one TestInstruction, to be sent over gRPC to Fenix backend
type TestInstructionInstanceVersionsStruct struct {
	TestInstructionVersions     []TestInstructionInstanceVersionStruct
	TestInstructionVersionsHash string // SHA256 of all TestInstructionVersions.TestInstructionInstanceHash using Fenix standard way of hashing values together
}

// TestInstructionsStruct
// Struct for all TestInstructions, to be sent over gPRC to Fenix backend
type TestInstructionsStruct struct {
	TestInstructions     []TestInstructionInstanceVersionsStruct
	TestInstructionsHash string // SHA256 of all TestInstructions.TestInstructionVersionsHash using Fenix standard way of hashing values together
}

// TestInstructionContainerInstanceVersionStruct
// Struct for one TestInstructionContainer, to be sent over gRPC to Fenix backend
type TestInstructionContainerInstanceVersionStruct struct {
	TestInstructionContainerInstance             TestInstructionContainerStruct
	TestInstructionContainerInstanceMajorVersion int
	TestInstructionContainerInstanceMinorVersion int
	Deprecated                                   bool
	Enabled                                      bool
	TestInstructionContainerInstanceHash         string
}

// TestInstructionContainerInstanceVersionsStruct
// Struct for all versions of one TestInstructionContainer, to be sent over gRPC to Fenix backend
type TestInstructionContainerInstanceVersionsStruct struct {
	TestInstructionContainerVersions     []TestInstructionContainerInstanceVersionStruct
	TestInstructionContainerVersionsHash string // SHA256 of all TestInstructionContainerVersions.TestInstructionContainerInstanceHash using Fenix standard way of hashing values together
}

// TestInstructionContainersStruct
// Struct for all TestInstructionContainers, to be sent over gPRC to Fenix backend
type TestInstructionContainersStruct struct {
	TestInstructionContainers     []TestInstructionContainerInstanceVersionsStruct
	TestInstructionContainersHash string // SHA256 of all TestInstructionContainers.TestInstructionContainerVersionsHash using Fenix standard way of hashing values together
}

// TestInstructionsAndTestInstructionsContainersStruct
// Struct for all TestInstructions and TestInstructionsContainers from a "System" that should be sent to Fenix backen
type TestInstructionsAndTestInstructionsContainersStruct struct {
	TestInstructions                                         TestInstructionsStruct
	TestInstructionContainers                                TestInstructionContainersStruct
	MessageCreationTimeStamp                                 time.Time
	TestInstructionsAndTestInstructionsContainersMessageHash string // SHA256(TestInstructionsHash concat TestInstructionContainersHash)
}

/*
// AllTestInstructions_SC_TestCaseSetUpStruct
// Struct for holding all data for all TestInstructions
type AllTestInstructions_SC_TestCaseSetUpStruct struct {
	TestInstruction                    []TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    []TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes []FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}


*/

package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
)

// TestInstruction_SC_TestCaseSetUpStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_SC_TestCaseSetUpStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}

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

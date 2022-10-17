package TestInstructions

import (
	"FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** DeleteSelectedInstructedParties
	TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties               TypeAndStructs.OriginalElementUUIDType = "e25e07ef-473b-42be-8a2e-5e1eb3c76cc1"
	TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties               TypeAndStructs.TestInstructionNameType = "DeleteSelectedInstructedParties"
	TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties                                                  = TestInstructionTypeUUID_CA_SettlementAgreementManagement
	TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedInstructedParties                                                  = TestInstructionTypeName_CA_SettlementAgreementManagement
	TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties        string                                 = "DeleteSelectedInstructedParties for an existing Settlement Agreement"
	TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties      string                                 = "DeleteSelectedInstructedParties for an existing Settlement Agreement"
	TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedInstructedParties         bool                                   = false
	TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedInstructedParties            bool                                   = true
	TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties int                                    = 1
	TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties int                                    = 0
	TestInstructionColor_CA_SettlementAgreement_DeleteSelectedInstructedParties              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_SettlementAgreement_DeleteSelectedInstructedParties                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_SettlementAgreement_DeleteSelectedInstructedParties                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"
)

// TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedPartiesStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedPartiesStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedPartiesStruct

// Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties() {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties = TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedPartiesStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			FangEngineClassNameUUID:  FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_SettlementAgreement,
			FangEngineClassNameNAME:  FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_SettlementAgreement,
			FangEngineMethodNameUUID: FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
			FangEngineMethodNameNAME: FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_SettlementAgreement_DeleteSelectedInstructedParties,
			Attributes:               make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - SettlementAgreement_DeleteSelectedInstructedParties
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - SettlementAgreement_DeleteSelectedInstructedParties
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TCRuleDeletion:               TCRuleDeletion_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TCRuleSwap:                   TCRuleSwap_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedInstructedParties,
	}

	// Dropzone
	// No Dropzone needed

	// TestInstruction
	// No Attributes needed

}

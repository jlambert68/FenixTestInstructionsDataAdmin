package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** AddSelectedInstructedParties
	TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties               TypeAndStructs.OriginalElementUUIDType = "1f59bf39-6201-4841-948b-109c04e1557a"
	TestInstructionName_CA_SettlementAgreement_AddSelectedInstructedParties               TypeAndStructs.TestInstructionNameType = "AddSelectedInstructedParties"
	TestInstructionTypeUUID_CA_SettlementAgreement_AddSelectedInstructedParties                                                  = TestInstructionTypeUUID_CA_SettlementAgreementManagement
	TestInstructionTypeName_CA_SettlementAgreement_AddSelectedInstructedParties                                                  = TestInstructionTypeName_CA_SettlementAgreementManagement
	TestInstructionDescription_CA_SettlementAgreement_AddSelectedInstructedParties        string                                 = "AddSelectedInstructedParties for an existing Settlement Agreement"
	TestInstructionMouseOverText_CA_SettlementAgreement_AddSelectedInstructedParties      string                                 = "AddSelectedInstructedParties for an existing Settlement Agreement"
	TestInstructionDeprecated_CA_SettlementAgreement_AddSelectedInstructedParties         bool                                   = false
	TestInstructionEnabled_CA_SettlementAgreement_AddSelectedInstructedParties            bool                                   = true
	TestInstructionMajorVersionNumber_CA_SettlementAgreement_AddSelectedInstructedParties int                                    = 0
	TestInstructionMinorVersionNumber_CA_SettlementAgreement_AddSelectedInstructedParties int                                    = 1
	TestInstructionColor_CA_SettlementAgreement_AddSelectedInstructedParties              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_SettlementAgreement_AddSelectedInstructedParties                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_SettlementAgreement_AddSelectedInstructedParties                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// Attribute - 'bicAddress'
	TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress          TypeAndStructs.TestInstructionAttributeUUIDType = "c8c0899b-c48a-448d-a964-c2cde5a3db77"
	TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress          TypeAndStructs.TestInstructionAttributeNameType = "bicAddress"
	TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress   string                                          = "bicAddress is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress string                                          = "bicAddress is xxxxx"
)

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties TestInstruction_CA_TestCaseSetUpStruct

// Initate_TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties() TestInstruction_CA_TestCaseSetUpStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			FangEngineClassNameUUID:  FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_SettlementAgreement,
			FangEngineClassNameNAME:  FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_SettlementAgreement,
			FangEngineMethodNameUUID: FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_SettlementAgreement_AddSelectedInstructedParties,
			FangEngineMethodNameNAME: FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_SettlementAgreement_AddSelectedInstructedParties,
			Attributes:               make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - SettlementAgreement_AddSelectedInstructedParties
	TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_AddSelectedInstructedParties,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_AddSelectedInstructedParties,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_AddSelectedInstructedParties,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_AddSelectedInstructedParties,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_AddSelectedInstructedParties,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - SettlementAgreement_AddSelectedInstructedParties
	TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_AddSelectedInstructedParties,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_AddSelectedInstructedParties,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_AddSelectedInstructedParties,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_AddSelectedInstructedParties,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_SettlementAgreement_AddSelectedInstructedParties,
		TCRuleDeletion:               TCRuleDeletion_CA_SettlementAgreement_AddSelectedInstructedParties,
		TCRuleSwap:                   TCRuleSwap_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_AddSelectedInstructedParties,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_AddSelectedInstructedParties,
	}

	// Dropzone Add
	// No Dropzone needed

	// TestInstruction Attribute - 'bicAddress'
	var TestInstructionAttribute_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
		TestInstructionAttributeTypeUUID:              TestInstructionAttributeTypeUUID_CA_Standard,
		TestInstructionAttributeTypeName:              TestInstructionAttributeTypeName_CA_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress)

	// Add FangEngine relation for Attribute - 'bicAddress'
	var tempFangEngineAttributebicAddress FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributebicAddress = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress] = tempFangEngineAttributebicAddress

	// ImmatureElementModel - 'SettlementAgreement' - MethodName: 'AddSelectedInstructedParties'
	var TestInstructionImmatureElementModel_CA_SettlementAgreement_AddSelectedInstructedParties TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_SettlementAgreement_AddSelectedInstructedParties = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_SettlementAgreement_AddSelectedInstructedParties),
		PreviousElementUUID:      TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		NextElementUUID:          TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		FirstChildElementUUID:    TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		ParentElementUUID:        TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		TopImmatureElementUUID:   TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties,
		IsTopElement:             true,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.ImmatureElementModel = append(
		TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_SettlementAgreement_AddSelectedInstructedParties)

	return TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties
}

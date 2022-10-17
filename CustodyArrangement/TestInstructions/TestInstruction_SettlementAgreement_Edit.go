package TestInstructions

import (
	"FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** Edit
	TestInstructionUUID_CA_SettlementAgreement_Edit               TypeAndStructs.OriginalElementUUIDType = "19c2a8b9-c48b-4f5e-9f41-a9b571f77ebb"
	TestInstructionName_CA_SettlementAgreement_Edit               TypeAndStructs.TestInstructionNameType = "Edit"
	TestInstructionTypeUUID_CA_SettlementAgreement_Edit                                                  = TestInstructionTypeUUID_CA_SettlementAgreementManagement
	TestInstructionTypeName_CA_SettlementAgreement_Edit                                                  = TestInstructionTypeName_CA_SettlementAgreementManagement
	TestInstructionDescription_CA_SettlementAgreement_Edit        string                                 = "Edit an existing Settlement Agreement"
	TestInstructionMouseOverText_CA_SettlementAgreement_Edit      string                                 = "Edit an existing Settlement Agreement"
	TestInstructionDeprecated_CA_SettlementAgreement_Edit         bool                                   = false
	TestInstructionEnabled_CA_SettlementAgreement_Edit            bool                                   = true
	TestInstructionMajorVersionNumber_CA_SettlementAgreement_Edit int                                    = 1
	TestInstructionMinorVersionNumber_CA_SettlementAgreement_Edit int                                    = 0
	TestInstructionColor_CA_SettlementAgreement_Edit              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_SettlementAgreement_Edit                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_SettlementAgreement_Edit                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// Attribute - 'TestDataDuplicateCheck'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck          TypeAndStructs.TestInstructionAttributeUUIDType = "48763abc-b792-4bff-9132-258ec747f7ea"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDuplicateCheck          TypeAndStructs.TestInstructionAttributeNameType = "TestDataDuplicateCheck"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDuplicateCheck          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDuplicateCheck   string                                          = "TestDataDuplicateCheck is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDuplicateCheck string                                          = "TestDataDuplicateCheck is xxxxx"

	// Attribute - 'TestDataContractualRule'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule          TypeAndStructs.TestInstructionAttributeUUIDType = "ef315036-f4bc-4255-a090-110a4e9ec836"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataContractualRule          TypeAndStructs.TestInstructionAttributeNameType = "TestDataContractualRule"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataContractualRule          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataContractualRule   string                                          = "TestDataContractualRule is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataContractualRule string                                          = "TestDataContractualRule is xxxxx"

	// Attribute - 'TestDataMatchAffirmEligible'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible          TypeAndStructs.TestInstructionAttributeUUIDType = "7894ee2d-c6d5-47a3-97f3-5b75cabe1abb"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible          TypeAndStructs.TestInstructionAttributeNameType = "TestDataMatchAffirmEligible"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible   string                                          = "TestDataMatchAffirmEligible is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible string                                          = "TestDataMatchAffirmEligible is xxxxx"

	// Attribute - 'TestDataDirectAffirmInd'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd          TypeAndStructs.TestInstructionAttributeUUIDType = "0dae0c87-c659-46e1-a03b-15231f91d567"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd          TypeAndStructs.TestInstructionAttributeNameType = "TestDataDirectAffirmInd"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd   string                                          = "TestDataDirectAffirmInd is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd string                                          = "TestDataDirectAffirmInd is xxxxx"

	// Attribute - 'TestDataAffirmDublicateCheck'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck          TypeAndStructs.TestInstructionAttributeUUIDType = "0dae0c87-c659-46e1-a03b-15231f91d567"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck          TypeAndStructs.TestInstructionAttributeNameType = "TestDataAffirmDublicateCheck"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck   string                                          = "TestDataAffirmDublicateCheck is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck string                                          = "TestDataAffirmDublicateCheck is xxxxx"

	// Attribute - 'TestDataAffirmCancelledInd'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd          TypeAndStructs.TestInstructionAttributeUUIDType = "36ebedb6-3cdb-437a-84e5-5fa18bd7b383"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd          TypeAndStructs.TestInstructionAttributeNameType = "TestDataAffirmCancelledInd"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd   string                                          = "TestDataAffirmCancelledInd is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd string                                          = "TestDataAffirmCancelledInd is xxxxx"

	// Attribute - 'TestDataNearMatch'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch          TypeAndStructs.TestInstructionAttributeUUIDType = "c08d107f-91f4-42a5-aff7-4e50e66bbfc5"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataNearMatch          TypeAndStructs.TestInstructionAttributeNameType = "TestDataNearMatch"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataNearMatch          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataNearMatch   string                                          = "TestDataNearMatch is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataNearMatch string                                          = "TestDataNearMatch is xxxxx"

	// Attribute - 'TestDataReturnDeliveryService'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService          TypeAndStructs.TestInstructionAttributeUUIDType = "b4addbd9-95d4-46d8-a044-946d9704afc5"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService          TypeAndStructs.TestInstructionAttributeNameType = "TestDataReturnDeliveryService"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService   string                                          = "TestDataReturnDeliveryService is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService string                                          = "TestDataReturnDeliveryService is xxxxx"
)

// TestInstruction_CA_SettlementAgreement_EditStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_CA_SettlementAgreement_EditStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_SettlementAgreement_Edit TestInstruction_CA_SettlementAgreement_EditStruct

// Initate_TestInstruction_CA_SettlementAgreement_Edit
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_SettlementAgreement_Edit() {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_SettlementAgreement_Edit = TestInstruction_CA_SettlementAgreement_EditStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			FangEngineClassNameUUID:  FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_SettlementAgreement,
			FangEngineClassNameNAME:  FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_SettlementAgreement,
			FangEngineMethodNameUUID: FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_SettlementAgreement_Edit,
			FangEngineMethodNameNAME: FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_SettlementAgreement_Edit,
			Attributes:               make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - SettlementAgreement_Edit
	TestInstruction_CA_SettlementAgreement_Edit.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_Edit,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_Edit,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_Edit,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_Edit,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_Edit,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_Edit,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_Edit,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_Edit,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - SettlementAgreement_Edit
	TestInstruction_CA_SettlementAgreement_Edit.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_Edit,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_Edit,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_Edit,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_Edit,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_Edit,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_SettlementAgreement_Edit,
		TCRuleDeletion:               TCRuleDeletion_CA_SettlementAgreement_Edit,
		TCRuleSwap:                   TCRuleSwap_CA_SettlementAgreement_Edit,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_Edit,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_Edit,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_Edit,
	}

	// Dropzone Add
	// No Dropzone needed

	// TestInstruction Attribute - 'TestDataDuplicateCheck'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDuplicateCheck TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDuplicateCheck = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDuplicateCheck)

	// Add FangEngine relation for Attribute - 'TestDataDuplicateCheck'
	var tempFangEngineAttributeTestDataDuplicateCheck FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataDuplicateCheck = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck] = tempFangEngineAttributeTestDataDuplicateCheck

	// TestInstruction Attribute - 'TestDataContractualRule'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataContractualRule TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataContractualRule = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataContractualRule,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataContractualRule,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataContractualRule)

	// Add FangEngine relation for Attribute - 'TestDataContractualRule'
	var tempFangEngineAttributeTestDataContractualRule FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataContractualRule = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataContractualRule,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataContractualRule,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataContractualRule,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule] = tempFangEngineAttributeTestDataContractualRule

	// TestInstruction Attribute - 'TestDataMatchAffirmEligible'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible)

	// Add FangEngine relation for Attribute - 'TestDataMatchAffirmEligible'
	var tempFangEngineAttributeTestDataMatchAffirmEligible FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataMatchAffirmEligible = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible] = tempFangEngineAttributeTestDataMatchAffirmEligible

	// TestInstruction Attribute - 'TestDataDirectAffirmInd'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd)

	// Add FangEngine relation for Attribute - 'TestDataDirectAffirmInd'
	var tempFangEngineAttributeTestDataDirectAffirmInd FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataDirectAffirmInd = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd] = tempFangEngineAttributeTestDataDirectAffirmInd

	// TestInstruction Attribute - 'TestDataAffirmDublicateCheck'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck)

	// Add FangEngine relation for Attribute - 'TestDataAffirmDublicateCheck'
	var tempFangEngineAttributeTestDataAffirmDublicateCheck FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataAffirmDublicateCheck = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck] = tempFangEngineAttributeTestDataAffirmDublicateCheck

	// TestInstruction Attribute - 'TestDataAffirmCancelledInd'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd)

	// Add FangEngine relation for Attribute - 'TestDataAffirmCancelledInd'
	var tempFangEngineAttributeTestDataAffirmCancelledInd FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataAffirmCancelledInd = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd] = tempFangEngineAttributeTestDataAffirmCancelledInd

	// TestInstruction Attribute - 'TestDataNearMatch'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataNearMatch TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataNearMatch = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataNearMatch,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataNearMatch,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataNearMatch)

	// Add FangEngine relation for Attribute - 'TestDataNearMatch'
	var tempFangEngineAttributeTestDataNearMatch FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataNearMatch = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataNearMatch,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataNearMatch,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataNearMatch,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch] = tempFangEngineAttributeTestDataNearMatch

	// TestInstruction Attribute - 'TestDataReturnDeliveryService'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService)

	// Add FangEngine relation for Attribute - 'TestDataReturnDeliveryService'
	var tempFangEngineAttributeTestDataReturnDeliveryService FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataReturnDeliveryService = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService] = tempFangEngineAttributeTestDataReturnDeliveryService

}

package TestInstructions

import (
	"FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
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
			FangEngineClassNameUUID:  FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_CustodyAccount,
			FangEngineClassNameNAME:  FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_CustodyAccount,
			FangEngineMethodNameUUID: FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_CustodyAccount_Search,
			FangEngineMethodNameNAME: FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_CustodyAccount_Search,
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

	// TestInstruction Attribute - 'TestDataInterval'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataInterval TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataInterval = TypeAndStructs.TestInstructionAttributeStruct{
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
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataInterval)

	// TestInstruction Attribute - 'TestDataType'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataType TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataType = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataType,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataType,
		TestInstructionAttributeDescription:           "The channel datatype",
		TestInstructionAttributeMouseOver:             "The channel datatype",
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataType,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataType)

	// TestInstruction Attribute - 'TestDataBIC'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataBIC TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataBIC = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataBIC,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataBIC,
		TestInstructionAttributeDescription:           "The BIC",
		TestInstructionAttributeMouseOver:             "The BIC",
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataBIC,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataBIC)

	// TestInstruction Attribute - 'TestDataChannel'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataChannel TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataChannel = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataChannel,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataChannel,
		TestInstructionAttributeDescription:           "The channel",
		TestInstructionAttributeMouseOver:             "The channel",
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataChannel,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataChannel)

	// ImmatureElementModel - SettlementAgreement_Edit
	var TestInstructionImmatureElementModel_CA_SettlementAgreement_Edit TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_SettlementAgreement_Edit = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_SettlementAgreement_Edit,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_SettlementAgreement_Edit),
		PreviousElementUUID:      TestInstructionUUID_CA_SettlementAgreement_Edit,
		NextElementUUID:          TestInstructionUUID_CA_SettlementAgreement_Edit,
		FirstChildElementUUID:    TestInstructionUUID_CA_SettlementAgreement_Edit,
		ParentElementUUID:        TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_SettlementAgreement_Edit,
		TopImmatureElementUUID:   TestInstructionUUID_CA_SettlementAgreement_Edit,
		IsTopElement:             true,
	}
	TestInstruction_CA_SettlementAgreement_Edit.ImmatureElementModel = append(
		TestInstruction_CA_SettlementAgreement_Edit.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_SettlementAgreement_Edit)
}

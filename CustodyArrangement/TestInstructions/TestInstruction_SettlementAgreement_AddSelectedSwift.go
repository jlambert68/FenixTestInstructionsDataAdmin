package TestInstructions

import (
	"FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** AddSelectedSwift
	TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift               TypeAndStructs.OriginalElementUUIDType = "b81cb0e0-589e-4209-beab-7b745dd349e3"
	TestInstructionName_CA_SettlementAgreement_AddSelectedSwift               TypeAndStructs.TestInstructionNameType = "AddSelectedSwift"
	TestInstructionTypeUUID_CA_SettlementAgreement_AddSelectedSwift                                                  = TestInstructionTypeUUID_CA_SettlementAgreementManagement
	TestInstructionTypeName_CA_SettlementAgreement_AddSelectedSwift                                                  = TestInstructionTypeName_CA_SettlementAgreementManagement
	TestInstructionDescription_CA_SettlementAgreement_AddSelectedSwift        string                                 = "AddSelectedSwift to an existing Settlement Agreement"
	TestInstructionMouseOverText_CA_SettlementAgreement_AddSelectedSwift      string                                 = "AddSelectedSwift to an existing Settlement Agreement"
	TestInstructionDeprecated_CA_SettlementAgreement_AddSelectedSwift         bool                                   = false
	TestInstructionEnabled_CA_SettlementAgreement_AddSelectedSwift            bool                                   = true
	TestInstructionMajorVersionNumber_CA_SettlementAgreement_AddSelectedSwift int                                    = 1
	TestInstructionMinorVersionNumber_CA_SettlementAgreement_AddSelectedSwift int                                    = 0
	TestInstructionColor_CA_SettlementAgreement_AddSelectedSwift              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_SettlementAgreement_AddSelectedSwift                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_SettlementAgreement_AddSelectedSwift                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// Attribute - 'TestDataType'
	TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataType          TypeAndStructs.TestInstructionAttributeUUIDType = "bc3cf0ae-8938-4ae5-bf0a-de51ed90e8c8"
	TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataType          TypeAndStructs.TestInstructionAttributeNameType = "TestDataType"
	TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataType          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataType   string                                          = "TestDataType is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataType string                                          = "TestDataType is xxxxx"

	// Attribute - 'TestDataChannel'
	TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel          TypeAndStructs.TestInstructionAttributeUUIDType = "e95c142c-be96-4635-a55c-a5175adb7390"
	TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel          TypeAndStructs.TestInstructionAttributeNameType = "TestDataChannel"
	TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel   string                                          = "TestDataChannel is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel string                                          = "TestDataChannel is xxxxx"

	// Attribute - 'TestDataBic'
	TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataBic          TypeAndStructs.TestInstructionAttributeUUIDType = "b73e9141-8682-4358-a9d5-40381e83d55b"
	TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataBic          TypeAndStructs.TestInstructionAttributeNameType = "TestDataBic"
	TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataBic          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataBic   string                                          = "TestDataBic is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataBic string                                          = "TestDataBic is xxxxx"

	// Attribute - 'TestDataInterval'
	TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval          TypeAndStructs.TestInstructionAttributeUUIDType = "0dae0c87-c659-46e1-a03b-15231f91d567"
	TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval          TypeAndStructs.TestInstructionAttributeNameType = "TestDataInterval"
	TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval   string                                          = "TestDataInterval is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval string                                          = "TestDataInterval is xxxxx"
)

// TestInstruction_CA_SettlementAgreement_AddSelectedSwiftStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_CA_SettlementAgreement_AddSelectedSwiftStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_SettlementAgreement_AddSelectedSwift TestInstruction_CA_SettlementAgreement_AddSelectedSwiftStruct

// Initate_TestInstruction_CA_SettlementAgreement_AddSelectedSwift
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_SettlementAgreement_AddSelectedSwift() {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift = TestInstruction_CA_SettlementAgreement_AddSelectedSwiftStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			FangEngineClassNameUUID:  FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_SettlementAgreement,
			FangEngineClassNameNAME:  FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_SettlementAgreement,
			FangEngineMethodNameUUID: FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_SettlementAgreement_AddSelectedSwift,
			FangEngineMethodNameNAME: FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_SettlementAgreement_AddSelectedSwift,
			Attributes:               make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - SettlementAgreement_AddSelectedSwift
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_AddSelectedSwift,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_AddSelectedSwift,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_AddSelectedSwift,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_AddSelectedSwift,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_AddSelectedSwift,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - SettlementAgreement_AddSelectedSwift
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_AddSelectedSwift,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_AddSelectedSwift,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_AddSelectedSwift,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_AddSelectedSwift,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_SettlementAgreement_AddSelectedSwift,
		TCRuleDeletion:               TCRuleDeletion_CA_SettlementAgreement_AddSelectedSwift,
		TCRuleSwap:                   TCRuleSwap_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_AddSelectedSwift,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_AddSelectedSwift,
	}

	// Dropzone Add
	// No Dropzone needed

	// TestInstruction Attribute - 'TestDataType'
	var TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataType TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataType = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataType)

	// Add FangEngine relation for Attribute - 'TestDataType'
	var tempFangEngineAttributeTestDataType FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataType = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataType,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataType] = tempFangEngineAttributeTestDataType

	// TestInstruction Attribute - 'TestDataChannel'
	var TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel)

	// Add FangEngine relation for Attribute - 'TestDataChannel'
	var tempFangEngineAttributeTestDataChannel FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataChannel = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel] = tempFangEngineAttributeTestDataChannel

	// TestInstruction Attribute - 'TestDataBic'
	var TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataBic TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataBic = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataBic)

	// Add FangEngine relation for Attribute - 'TestDataBic'
	var tempFangEngineAttributeTestDataBic FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataBic = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataBic,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataBic] = tempFangEngineAttributeTestDataBic

	// TestInstruction Attribute - 'TestDataInterval'
	var TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_AddSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval)

	// Add FangEngine relation for Attribute - 'TestDataInterval'
	var tempFangEngineAttributeTestDataInterval FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataInterval = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval,
	}
	TestInstruction_CA_SettlementAgreement_AddSelectedSwift.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval] = tempFangEngineAttributeTestDataInterval

}
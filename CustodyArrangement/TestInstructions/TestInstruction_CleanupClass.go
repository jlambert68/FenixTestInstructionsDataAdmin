package TestInstructions

import (
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *************************************
	// *** TestInstruction *** CleanupClass
	TestInstructionUUID_CA_CleanupClass               TypeAndStructs.OriginalElementUUIDType = "5501bf8a-0512-476e-a8bd-75d0be3e2bad"
	TestInstructionName_CA_CleanupClass               TypeAndStructs.TestInstructionNameType = "CleanupClass"
	TestInstructionTypeUUID_CA_CleanupClass                                                  = TestInstructionTypeUUID_CA_CustodyAccount
	TestInstructionTypeName_CA_CleanupClass                                                  = TestInstructionTypeName_CA_CustodyAccount
	TestInstructionDescription_CA_CleanupClass        string                                 = "A TestInstruction that's cool"
	TestInstructionMouseOverText_CA_CleanupClass      string                                 = "This will be shown when hovering above this TestInstruction"
	TestInstructionDeprecated_CA_CleanupClass         bool                                   = false
	TestInstructionEnabled_CA_CleanupClass            bool                                   = true
	TestInstructionMajorVersionNumber_CA_CleanupClass int                                    = 1
	TestInstructionMinorVersionNumber_CA_CleanupClass int                                    = 0
	TestInstructionColor_CA_CleanupClass              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_CA_CleanupClass                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_CA_CleanupClass                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"

	// *** DropZone *** DeleteLocalMarketAndConfirm
	TestInstructionDropZoneUUID_CA_CleanupClass_DeleteLocalMarketAndConfirm        TypeAndStructs.DropZoneUUIDType = "0692a87b-e111-4d6c-9ec3-57f7891ee957"
	TestInstructionDropZoneName_CA_CleanupClass_DeleteLocalMarketAndConfirm        TypeAndStructs.DropZoneNameType = "DeleteLocalMarketAndConfirm"
	TestInstructionDropZoneDescription_CA_CleanupClass_DeleteLocalMarketAndConfirm string                          = "Delete local market and confirm"
	TestInstructionDropZoneMouseOver_CA_CleanupClass_DeleteLocalMarketAndConfirm   string                          = "Delete local market and confirm"
	TestInstructionDropZoneColor_CA_CleanupClass_DeleteLocalMarketAndConfirm       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'CustodyAccountId'
	TestInstructionAttributeUUID_CA_CleanupClass_CustodyAccountId          TypeAndStructs.TestInstructionAttributeUUIDType = "831b9553-f05e-4912-a832-69f785da9256"
	TestInstructionAttributeName_CA_CleanupClass_CustodyAccountId          TypeAndStructs.TestInstructionAttributeNameType = "CustodyAccountId"
	TestInstructionAttributeType_CA_CleanupClass_CustodyAccountId          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeActionCommand_CA_CleanupClass_CustodyAccountId                                                 = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE

	// Attribute - 'MarketName'
	TestInstructionAttributeUUID_CA_CleanupClass_MarketName          TypeAndStructs.TestInstructionAttributeUUIDType = "0918d9c1-da59-4d62-99ee-0e38cd2a078f"
	TestInstructionAttributeName_CA_CleanupClass_MarketName          TypeAndStructs.TestInstructionAttributeNameType = "MarketName"
	TestInstructionAttributeType_CA_CleanupClass_MarketName          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeActionCommand_CA_CleanupClass_MarketName                                                 = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
)

// TestInstruction_CA_CleanupClassStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_CA_CleanupClassStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
}

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_CleanupClass TestInstruction_CA_CleanupClassStruct

// Initate_TestInstruction_CA_CleanupClass
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_CleanupClass() {

	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - CleanupClass
	TestInstruction_CA_CleanupClass.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:          TestInstructionName_CA_CleanupClass,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CleanupClass,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CleanupClass,
		TestInstructionDescription:   TestInstructionDescription_CA_CleanupClass,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_CleanupClass,
		Deprecated:                   TestInstructionDeprecated_CA_CleanupClass,
		Enabled:                      TestInstructionEnabled_CA_CleanupClass,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_CleanupClass,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_CleanupClass,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - CleanupClass
	TestInstruction_CA_CleanupClass.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:          TestInstructionName_CA_CleanupClass,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CleanupClass,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CleanupClass,
		Deprecated:                   TestInstructionDeprecated_CA_CleanupClass,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_CleanupClass,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_CleanupClass,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_CleanupClass,
		TCRuleDeletion:               TCRuleDeletion_CA_CleanupClass,
		TCRuleSwap:                   TCRuleSwap_CA_CleanupClass,
		TestInstructionDescription:   TestInstructionDescription_CA_CleanupClass,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_CleanupClass,
		Enabled:                      TestInstructionEnabled_CA_CleanupClass,
	}

	// ImmatureTestInstructionInformation  - DropZone: DeleteLocalMarketAndConfirm, Attr: CustodyAccountId
	var TestInstruction_CA_CleanupClass_CustodyAccountId TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_CleanupClass_CustodyAccountId = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:          TestInstructionName_CA_CleanupClass,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneName:                 TestInstructionDropZoneName_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneColor:                TestInstructionDropZoneColor_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_CleanupClass_CustodyAccountId,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_CleanupClass_CustodyAccountId,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_CleanupClass_CustodyAccountId,
		AttributeValueAsString:       "This is a value From DropZone for Custody Account Id",
		AttributeValueUUID:           "0fc988f2-7add-4c54-88c4-2c1acb788af0",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_CleanupClass,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_CleanupClass_CustodyAccountId,
	}
	TestInstruction_CA_CleanupClass.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_CleanupClass.ImmatureTestInstructionInformation,
		TestInstruction_CA_CleanupClass_CustodyAccountId)

	// DropZone: DeleteLocalMarketAndConfirm
	// ImmatureTestInstructionInformation  - DropZone: DeleteLocalMarketAndConfirm, Attr: MarketName
	var TestInstruction_CA_CleanupClass_DeleteLocalMarketAndConfirm_MarketName TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_CleanupClass_DeleteLocalMarketAndConfirm_MarketName = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:          TestInstructionName_CA_CleanupClass,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneName:                 TestInstructionDropZoneName_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneColor:                TestInstructionColor_CA_CleanupClass,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_CleanupClass_MarketName,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_CleanupClass_MarketName,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_CleanupClass_MarketName,
		AttributeValueAsString:       "DropZone value for MarketName",
		AttributeValueUUID:           "4e8f5be7-975c-48ff-9b21-b27f818bd23d",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_CleanupClass,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_CleanupClass_MarketName,
	}
	TestInstruction_CA_CleanupClass.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_CleanupClass.ImmatureTestInstructionInformation,
		TestInstruction_CA_CleanupClass_DeleteLocalMarketAndConfirm_MarketName)

	// TestInstruction Attribute - 'CustodyAccountId'
	var TestInstructionAttribute_CA_CleanupClass_CustodyAccountId TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_CleanupClass_CustodyAccountId = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:                           TestInstructionName_CA_CleanupClass,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_CleanupClass_CustodyAccountId,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_CleanupClass_CustodyAccountId,
		TestInstructionAttributeDescription:           "Custody Account that should be cleaned",
		TestInstructionAttributeMouseOver:             "Custody Account that should be cleaned",
		TestInstructionAttributeTypeUUID:              TestInstructionAttributeTypeUUID_CA_Standard,
		TestInstructionAttributeTypeName:              TestInstructionAttributeTypeName_CA_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValueNO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUIDNo_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_CleanupClass_CustodyAccountId,
	}
	TestInstruction_CA_CleanupClass.TestInstructionAttribute = append(
		TestInstruction_CA_CleanupClass.TestInstructionAttribute,
		TestInstructionAttribute_CA_CleanupClass_CustodyAccountId)

	// TestInstruction Attribute - 'MarketName'
	var TestInstructionAttribute_CA_CleanupClass_MarketName TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_CleanupClass_MarketName = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:                           TestInstructionName_CA_CleanupClass,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_CleanupClass_MarketName,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_CleanupClass_MarketName,
		TestInstructionAttributeDescription:           "Custody Account that should be cleaned",
		TestInstructionAttributeMouseOver:             "Custody Account that should be cleaned",
		TestInstructionAttributeTypeUUID:              TestInstructionAttributeTypeUUID_CA_Standard,
		TestInstructionAttributeTypeName:              TestInstructionAttributeTypeName_CA_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValueNO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUIDNo_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_CleanupClass_MarketName,
	}
	TestInstruction_CA_CleanupClass.TestInstructionAttribute = append(
		TestInstruction_CA_CleanupClass.TestInstructionAttribute,
		TestInstructionAttribute_CA_CleanupClass_MarketName)

}

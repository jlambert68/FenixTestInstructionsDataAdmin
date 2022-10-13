package CustodyArrangement

import (
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *************************************
	// *** TestInstruction *** - *** TestInstruction ***
	TestInstructionUUID_CA_CleanupClass TypeAndStructs.OriginalElementUUIDType = "5501bf8a-0512-476e-a8bd-75d0be3e2bad"
	TestInstructionName_CA_CleanupClass TypeAndStructs.TestInstructionNameType = "CleanupClass"

	// *** DropZone ***
	TestInstructionDropZoneUUID_CA_CleanupClass_DeleteLocalMarketAndConfirm                 TypeAndStructs.DropZoneUUIDType = "0692a87b-e111-4d6c-9ec3-57f7891ee957"
	TestInstructionDropZoneName_CA_CleanupClass_DeleteLocalMarketAndConfirmDropZoneUUIDType TypeAndStructs.DropZoneNameType = "DeleteLocalMarketAndConfirm"
	TestInstructionDropZoneDescription                                                      string                          = "Delete local market and confirm"
	TestInstructionDropZoneMouseOver                                                        string                          = "Delete local market and confirm"
	TestInstructionDropZoneColor                                                            string                          = "#00000000"

	// Attribute
	TestInstructionAttributeUUID_CA_CleanupClass_CustodyAccountId TypeAndStructs.AttributeValueUUIDType           = "831b9553-f05e-4912-a832-69f785da9256"
	TestInstructionAttributeNAme_CA_CleanupClass_CustodyAccountId TypeAndStructs.TestInstructionAttributeNameType = "CustodyAccountId"

	TestInstructionAttributeUUID_CA_CleanupClass_MarketName TypeAndStructs.AttributeValueUUIDType           = "0918d9c1-da59-4d62-99ee-0e38cd2a078f"
	TestInstructionAttributeNAme_CA_CleanupClass_MarketName TypeAndStructs.TestInstructionAttributeNameType = "MarketName"
)

const (
	TestInstructionDescription        string                            = "A TestInstruction that's cool"
	TestInstructionMouseOverText      string                            = "This will be shown when hovering above this TestInstruction"
	TestInstructionDeprecated         bool                              = false
	TestInstructionEnabled            bool                              = true
	TestInstructionMajorVersionNumber int                               = 1
	TestInstructionMinorVersionNumber int                               = 0
	TestInstructionColor              TypeAndStructs.ColorType          = "#00ff00AA"
	TCRuleDeletion                    TypeAndStructs.TCRuleDeletionType = "TCRuleDeletion020"
	TCRuleSwap                        TypeAndStructs.TCRuleSwapType     = "TCRuleSwap020"
)

type TestInstruction_CA_CleanupClassStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           TypeAndStructs.TestInstructionAttributeStruct
}

var TestInstruction_CA_CleanupClass TestInstruction_CA_CleanupClassStruct

func Initate_TestInstruction_CA_CleanupClass() {

	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	TestInstruction_CA_CleanupClass.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:          TestInstructionName_CA_CleanupClass,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CustodyAccount,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CustodyAccount,
		TestInstructionDescription:   TestInstructionDescription,
		TestInstructionMouseOverText: TestInstructionMouseOverText,
		Deprecated:                   TestInstructionDeprecated,
		Enabled:                      TestInstructionEnabled,
		MajorVersionNumber:           TestInstructionMajorVersionNumber,
		MinorVersionNumber:           TestInstructionMinorVersionNumber,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	TestInstruction_CA_CleanupClass.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:          TestInstructionName_CA_CleanupClass,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CustodyAccount,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CustodyAccount,
		Deprecated:                   TestInstructionDeprecated,
		MajorVersionNumber:           TestInstructionMajorVersionNumber,
		MinorVersionNumber:           TestInstructionMinorVersionNumber,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor,
		TCRuleDeletion:               TCRuleDeletion,
		TCRuleSwap:                   TCRuleSwap,
		TestInstructionDescription:   TestInstructionDescription,
		TestInstructionMouseOverText: TestInstructionMouseOverText,
		Enabled:                      TestInstructionEnabled,
	}

	TestInstruction_CA_CleanupClass.ImmatureTestInstructionInformation = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:          TestInstructionName_CA_CleanupClass,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_CleanupClass_DeleteLocalMarketAndConfirm,
		DropZoneName:                 TestInstructionDropZoneName_CA_CleanupClass_DeleteLocalMarketAndConfirmDropZoneUUIDType,
		DropZoneDescription:          TestInstructionDropZoneDescription,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver,
		DropZoneColor:                TestInstructionColor,
		TestInstructionAttributeType: "",
		TestInstructionAttributeUUID: "",
		TestInstructionAttributeName: "",
		AttributeValueAsString:       "",
		AttributeValueUUID:           "",
		FirstImmatureElementUUID:     "",
		AttributeActionCommand:       0,
	}

	TestInstruction_CA_CleanupClass.TestInstructionAttribute = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_CleanupClass,
		TestInstructionName:                           TestInstructionName_CA_CleanupClass,
		TestInstructionAttributeUUID:                  "",
		TestInstructionAttributeName:                  "",
		TestInstructionAttributeDescription:           "",
		TestInstructionAttributeMouseOver:             "",
		TestInstructionAttributeTypeUUID:              "",
		TestInstructionAttributeTypeName:              "",
		TestInstructionAttributeValueAsString:         "",
		TestInstructionAttributeValueUUID:             "",
		TestInstructionAttributeVisible:               false,
		TestInstructionAttributeEnabled:               false,
		TestInstructionAttributeMandatory:             false,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             "",
		TestInstructionAttributeType:                  "",
	}

}

package CustodyArrangement

import (
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** AddOrDeleteSelectedSwift
	TestInstructionUUID_CA_AddOrDeleteSelectedSwift               TypeAndStructs.OriginalElementUUIDType = "5501bf8a-0512-476e-a8bd-75d0be3e2bad"
	TestInstructionName_CA_AddOrDeleteSelectedSwift               TypeAndStructs.TestInstructionNameType = "AddOrDeleteSelectedSwift"
	TestInstructionTypeUUID_CA_AddOrDeleteSelectedSwift                                                  = TestInstructionTypeUUID_CA_FundExecutionAgreementManagement
	TestInstructionTypeName_CA_AddOrDeleteSelectedSwift                                                  = TestInstructionTypeName_CA_FundExecutionAgreementManagement
	TestInstructionDescription_CA_AddOrDeleteSelectedSwift        string                                 = "xxxxxxx"
	TestInstructionMouseOverText_CA_AddOrDeleteSelectedSwift      string                                 = "xxxxxxxx"
	TestInstructionDeprecated_CA_AddOrDeleteSelectedSwift         bool                                   = false
	TestInstructionEnabled_CA_AddOrDeleteSelectedSwift            bool                                   = true
	TestInstructionMajorVersionNumber_CA_AddOrDeleteSelectedSwift int                                    = 1
	TestInstructionMinorVersionNumber_CA_AddOrDeleteSelectedSwift int                                    = 0
	TestInstructionColor_CA_AddOrDeleteSelectedSwift              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_AddOrDeleteSelectedSwift                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_CA_AddOrDeleteSelectedSwift                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"

	// *** DropZone *** Add
	TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Add        TypeAndStructs.DropZoneUUIDType = "60c864ef-8c12-4028-9f58-f372416dbb87"
	TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Add        TypeAndStructs.DropZoneNameType = "Add"
	TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Add string                          = "Add Selected Swift"
	TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Add   string                          = "Add Selected Swift"
	TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Add       TypeAndStructs.ColorType        = "#F0FEF0FF"

	// *** DropZone *** Delete
	TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Delete        TypeAndStructs.DropZoneUUIDType = "d204851f-2d4c-4cfa-bee9-a48ac806ada1"
	TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Delete        TypeAndStructs.DropZoneNameType = "Delete"
	TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Delete string                          = "Delete Selected Swift"
	TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Delete   string                          = "Delete Selected Swift"
	TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Delete       TypeAndStructs.ColorType        = "#F3F0F0FF"

	// Attribute - 'TestDataInterval'
	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataInterval          TypeAndStructs.TestInstructionAttributeUUIDType = "d45806fa-37d8-4720-9f39-2e78edfcfc26"
	TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataInterval          TypeAndStructs.TestInstructionAttributeNameType = "TestDataInterval"
	TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataInterval          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataInterval                                                 = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE

	// Attribute - 'TestDataType'
	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataType          TypeAndStructs.TestInstructionAttributeUUIDType = "e3b8dcab-76a7-4059-a732-94b9aafe7650"
	TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataType          TypeAndStructs.TestInstructionAttributeNameType = "TestDataType"
	TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataType          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataType                                                 = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE

	// Attribute - 'TestDataBIC'
	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataBIC          TypeAndStructs.TestInstructionAttributeUUIDType = "fbd294a9-b6e2-43a9-86ae-a49c54d71021"
	TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataBIC          TypeAndStructs.TestInstructionAttributeNameType = "TestDataBIC"
	TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataBIC          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataBIC                                                 = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE

	// Attribute - 'TestDataChannel'
	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataChannel          TypeAndStructs.TestInstructionAttributeUUIDType = "4b728987-f371-4e4a-b9c6-8e4df1125312"
	TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataChannel          TypeAndStructs.TestInstructionAttributeNameType = "TestDataChannel"
	TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataChannel          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataChannel                                                 = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
)

// TestInstruction_CA_AddOrDeleteSelectedSwiftStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_CA_AddOrDeleteSelectedSwiftStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
}

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_AddOrDeleteSelectedSwift TestInstruction_CA_AddOrDeleteSelectedSwiftStruct

// Initate_TestInstruction_CA_AddOrDeleteSelectedSwift
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_AddOrDeleteSelectedSwift() {

	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - AddOrDeleteSelectedSwift
	TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionTypeName:      TestInstructionTypeName_CA_AddOrDeleteSelectedSwift,
		TestInstructionDescription:   TestInstructionDescription_CA_AddOrDeleteSelectedSwift,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_AddOrDeleteSelectedSwift,
		Deprecated:                   TestInstructionDeprecated_CA_AddOrDeleteSelectedSwift,
		Enabled:                      TestInstructionEnabled_CA_AddOrDeleteSelectedSwift,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_AddOrDeleteSelectedSwift,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_AddOrDeleteSelectedSwift,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - AddOrDeleteSelectedSwift
	TestInstruction_CA_AddOrDeleteSelectedSwift.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionTypeName:      TestInstructionTypeName_CA_AddOrDeleteSelectedSwift,
		Deprecated:                   TestInstructionDeprecated_CA_AddOrDeleteSelectedSwift,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_AddOrDeleteSelectedSwift,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_AddOrDeleteSelectedSwift,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_AddOrDeleteSelectedSwift,
		TCRuleDeletion:               TCRuleDeletion_CA_AddOrDeleteSelectedSwift,
		TCRuleSwap:                   TCRuleSwap_CA_AddOrDeleteSelectedSwift,
		TestInstructionDescription:   TestInstructionDescription_CA_AddOrDeleteSelectedSwift,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_AddOrDeleteSelectedSwift,
		Enabled:                      TestInstructionEnabled_CA_AddOrDeleteSelectedSwift,
	}

	// Dropzone Add
	// ImmatureTestInstructionInformation  - DropZone: Add, Attr: TestDataInterval
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataInterval TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataInterval = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Add,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		AttributeValueAsString:       "DZ for TestDataInterval",
		AttributeValueUUID:           "979b1f1a-9628-43eb-ab25-7c5572412be4",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataInterval,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataInterval)

	// ImmatureTestInstructionInformation  - DropZone: Add, Attr: TestDataType
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataType TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataType = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Add,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataType,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataType,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataType,
		AttributeValueAsString:       "DZ for TestDataType",
		AttributeValueUUID:           "a5a80f94-f2d9-432f-8c69-437d8cef3f51",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataType,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataType)

	// ImmatureTestInstructionInformation  - DropZone: Add, Attr: TestDataBIC
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataBIC TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataBIC = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Add,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		AttributeValueAsString:       "DZ for TestDataBIC",
		AttributeValueUUID:           "6cfcddb3-6213-4595-ab1f-4af8368c30f7",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataBIC,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataBIC)

	// ImmatureTestInstructionInformation  - DropZone: Add, Attr: TestDataChannel
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataChannel TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataChannel = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Add,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Add,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		AttributeValueAsString:       "DZ for TestDataChannel",
		AttributeValueUUID:           "cb2597d6-5d76-4631-93b2-b729e764e2a7",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataChannel,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Add_TestDataChannel)

	// DropZone: Delete
	// ImmatureTestInstructionInformation  - DropZone: Delete, Attr: TestDataInterval
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataInterval TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataInterval = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Delete,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		AttributeValueAsString:       "DZ for TestDataInterval",
		AttributeValueUUID:           "979b1f1a-9628-43eb-ab25-7c5572412be4",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataInterval,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataInterval)

	// ImmatureTestInstructionInformation  - DropZone: Delete, Attr: TestDataType
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataType TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataType = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Delete,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataType,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataType,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataType,
		AttributeValueAsString:       "DZ for TestDataType",
		AttributeValueUUID:           "a5a80f94-f2d9-432f-8c69-437d8cef3f51",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataType,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataType)

	// ImmatureTestInstructionInformation  - DropZone: Delete, Attr: TestDataBIC
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataBIC TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataBIC = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Delete,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		AttributeValueAsString:       "DZ for TestDataBIC",
		AttributeValueUUID:           "6cfcddb3-6213-4595-ab1f-4af8368c30f7",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataBIC,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataBIC)

	// ImmatureTestInstructionInformation  - DropZone: Delete, Attr: TestDataChannel
	var TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataChannel TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataChannel = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_AddOrDeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneName:                 TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_AddOrDeleteSelectedSwift_Delete,
		DropZoneColor:                TestInstructionDropZoneColor_CA_AddOrDeleteSelectedSwift_Delete,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		AttributeValueAsString:       "DZ for TestDataChannel",
		AttributeValueUUID:           "cb2597d6-5d76-4631-93b2-b729e764e2a7",
		FirstImmatureElementUUID:     TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_AddOrDeleteSelectedSwift_TestDataChannel,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_AddOrDeleteSelectedSwift_Delete_TestDataChannel)

	// TestInstruction Attribute - 'TestDataInterval'
	var TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataInterval TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataInterval = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_AddOrDeleteSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataInterval,
		TestInstructionAttributeDescription:           "The Interval",
		TestInstructionAttributeMouseOver:             "The Interval",
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataInterval,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataInterval)

	// TestInstruction Attribute - 'TestDataType'
	var TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataType TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataType = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_AddOrDeleteSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataType,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataType,
		TestInstructionAttributeDescription:           "The channel datatype",
		TestInstructionAttributeMouseOver:             "The channel datatype",
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataType,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataType)

	// TestInstruction Attribute - 'TestDataBIC'
	var TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataBIC TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataBIC = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_AddOrDeleteSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataBIC,
		TestInstructionAttributeDescription:           "The BIC",
		TestInstructionAttributeMouseOver:             "The BIC",
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataBIC,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataBIC)

	// TestInstruction Attribute - 'TestDataChannel'
	var TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataChannel TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataChannel = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_AddOrDeleteSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_AddOrDeleteSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_AddOrDeleteSelectedSwift_TestDataChannel,
		TestInstructionAttributeDescription:           "The channel",
		TestInstructionAttributeMouseOver:             "The channel",
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_AddOrDeleteSelectedSwift_TestDataChannel,
	}
	TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_AddOrDeleteSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_AddOrDeleteSelectedSwift_TestDataChannel)

}

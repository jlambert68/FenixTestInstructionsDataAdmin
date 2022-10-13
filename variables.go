package main

type DomainUUIDType string
type DomainNameType string
type OriginalElementUUIDType string
type OriginalElementNameType string
type TestInstructionNameType string
type TestInstructionTypeUUIDType string
type TestInstructionTypeNameType string
type UpdatedTimeStampType string
type colorType string
type TCRuleDeletionType string
type TCRuleSwapType string
type DropZoneUUIDType string
type DropZoneNameType string
type TestInstructionAttributeTypeType string
type TestInstructionAttributeUUIDType string
type TestInstructionAttributeNameType string
type AttributeValueAsStringType string
type AttributeValueUUIDType string
type AttributeActionCommandType int
type TestInstructionAttributeInputMaskType string
type TestInstructionAttributeTypeUUIDType string
type TestInstructionAttributeTypeNameType string

type TestInstructionContainerNameType string
type TestInstructionContainerTypeUUIDType string
type TestInstructionContainerTypeNameType string
type TestInstructionContainerExecutionTypeType string //'SERIAL_PROCESSED', 'PARALLELLED_PROCESSED'

type TestCaseModelElementTypeType string

// ** Domains ***
const (
	DomainUUID_CA DomainUUIDType = "78a97c41-a098-4122-88d2-01ed4b6c4844"
	DomainName_CA DomainNameType = "Custody Arrangement"
)

// *** TestInstructions - Custody Arrangement ***
const (

	// *************************************
	// *** TestInstructionType ***
	TestInstructionTypeUUID_CA_FundExecutionAgreementManagement TestInstructionTypeUUIDType = "ceca6136-66df-4dcc-930f-44d729142334"
	TestInstructionTypeName_CA_FundExecutionAgreementManagement TestInstructionTypeUUIDType = "Fund Execution Agreement management"

	TestInstructionTypeUUID_CA_CustodyAccount TestInstructionTypeUUIDType = "5005928a-f253-4828-9c2f-023571760759"
	TestInstructionTypeName_CA_CustodyAccount TestInstructionTypeUUIDType = "Custody Account"

	// *************************************
	// *** TestInstruction *** - *** TestInstruction ***
	TestInstructionUUID_CA_CleanupClass OriginalElementUUIDType = "5501bf8a-0512-476e-a8bd-75d0be3e2bad"
	TestInstructionName_CA_CleanupClass TestInstructionNameType = "CleanupClass"

	// *** DropZone ***
	TestInstructionDropZoneUUID_CA_CleanupClass_DeleteLocalMarketAndConfirm                 DropZoneUUIDType = "0692a87b-e111-4d6c-9ec3-57f7891ee957"
	TestInstructionDropZoneName_CA_CleanupClass_DeleteLocalMarketAndConfirmDropZoneUUIDType DropZoneNameType = "DeleteLocalMarketAndConfirm"

	// Attribute
	TestInstructionAttributeUUID_CA_CleanupClass_CustodyAccountId AttributeValueUUIDType           = "831b9553-f05e-4912-a832-69f785da9256"
	TestInstructionAttributeNAme_CA_CleanupClass_CustodyAccountId TestInstructionAttributeNameType = "CustodyAccountId"

	TestInstructionAttributeUUID_CA_CleanupClass_MarketName AttributeValueUUIDType           = "0918d9c1-da59-4d62-99ee-0e38cd2a078f"
	TestInstructionAttributeNAme_CA_CleanupClass_MarketName TestInstructionAttributeNameType = "MarketName"

	// *** TestInstruction *** - *** TestInstruction ***
	TestInstructionUUID_CA_AddOrDeleteSelectedSwift OriginalElementUUIDType = "ceb23b06-c1b3-4981-9806-0a53bbe42a67"
	TestInstructionName_CA_AddOrDeleteSelectedSwift TestInstructionNameType = "Add or Delete Selected Swift"

	// *** DropZone ***
	TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Add DropZoneUUIDType = "60c864ef-8c12-4028-9f58-f372416dbb87"
	TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Add DropZoneNameType = "Add"

	TestInstructionDropZoneUUID_CA_AddOrDeleteSelectedSwift_Delete DropZoneUUIDType = "d204851f-2d4c-4cfa-bee9-a48ac806ada1"
	TestInstructionDropZoneName_CA_AddOrDeleteSelectedSwift_Delete DropZoneNameType = "Delete"

	// Attribute
	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataInterval AttributeValueUUIDType           = "d45806fa-37d8-4720-9f39-2e78edfcfc26"
	TestInstructionAttributeNAme_CA_AddOrDeleteSelectedSwift_TestDataInterval TestInstructionAttributeNameType = "TestDataInterval"

	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataType AttributeValueUUIDType           = "e3b8dcab-76a7-4059-a732-94b9aafe7650"
	TestInstructionAttributeNAme_CA_AddOrDeleteSelectedSwift_TestDataType TestInstructionAttributeNameType = "TestDataType"

	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataBIC AttributeValueUUIDType           = "fbd294a9-b6e2-43a9-86ae-a49c54d71021"
	TestInstructionAttributeNAme_CA_AddOrDeleteSelectedSwift_TestDataBIC TestInstructionAttributeNameType = "TestDataBIC"

	TestInstructionAttributeUUID_CA_AddOrDeleteSelectedSwift_TestDataChannel AttributeValueUUIDType           = "4b728987-f371-4e4a-b9c6-8e4df1125312"
	TestInstructionAttributeNAme_CA_AddOrDeleteSelectedSwift_TestDataChannel TestInstructionAttributeNameType = "TestDataChannel"
)

// *** TestInstructionContainers - Custody Arrangement ***
const (

	// *************************************
	// *** TestInstructionContainerType ***
	TestInstructionContainerTypeUUID_CA_BaseContainers     TestInstructionContainerTypeUUIDType = "aa1b9734-5dce-43c9-8d77-3368940cf126"
	TestInstructionContainerTypeNameType_CA_BaseContainers TestInstructionContainerTypeNameType = "CA Base containers"

	// *************************************
	// *** TestInstructionContainer ***
	TestInstructionContainerUUID_CA_EmptyParallellContainer OriginalElementUUIDType          = "aa1b9734-5dce-43c9-8d77-3368940cf126"
	TestInstructionContainerName_CA_EmptyParallellContainer TestInstructionContainerNameType = "Empty parallelled processed Turbo TestInstructionsContainer"
)

// *** TestInstructionContainers - Fenix ***
const (
	// *************************************
	// *** TestInstructionContainerType ***
	TestInstructionContainerTypeUUID_Fenix_BaseContainers     TestInstructionContainerTypeUUIDType = "b107bdd9-4152-4020-b3f0-fc750b45885e"
	TestInstructionContainerTypeNameType_Fenix_BaseContainers TestInstructionContainerTypeNameType = "Base containers"

	// *************************************
	// *** TestInstructionContainer ***
	TestInstructionContainerUUID_Fenix_EmptySerialContainer OriginalElementUUIDType          = "f81b9734-5dce-43c9-8d77-3368940cf126"
	TestInstructionContainerName_Fenix_EmptySerialContainer TestInstructionContainerNameType = "Emtpy serial processed TestInstructionsContainer"

	// *** DropZone ***
	TestInstructionContainerDropZoneUUID_Fenix_EmptySerialContainer DropZoneUUIDType = "c5e37024-e40c-49f7-8667-eab485c65105"
	TestInstructionContainerDropZoneName_Fenix_EmptySerialContainer DropZoneUUIDType = "My first DropZone for a TestInstructionContainer"

	// *** TestInstructionContainer ***
	TestInstructionContainerUUID_Fenix_EmptyParallellContainer OriginalElementUUIDType          = "aa1b9734-5dce-43c9-8d77-3368940cf126"
	TestInstructionContainerName_Fenix_EmptyParallelContainer  TestInstructionContainerNameType = "Emtpy parallelled processed TestInstructionsContainer"
)

type TestInstructions []struct {
	DomainUUID                   DomainUUIDType              `json:"DomainUuid"`
	DomainName                   DomainNameType              `json:"DomainName"`
	TestInstructionUUID          OriginalElementUUIDType     `json:"TestInstructionUuid"`
	TestInstructionName          TestInstructionNameType     `json:"TestInstructionName"`
	TestInstructionTypeUUID      TestInstructionTypeUUIDType `json:"TestInstructionTypeUuid"`
	TestInstructionTypeName      TestInstructionTypeNameType `json:"TestInstructionTypeName"`
	TestInstructionDescription   string                      `json:"TestInstructionDescription"`
	TestInstructionMouseOverText string                      `json:"TestInstructionMouseOverText"`
	Deprecated                   bool                        `json:"Deprecated"`
	Enabled                      bool                        `json:"Enabled"`
	MajorVersionNumber           int                         `json:"MajorVersionNumber"`
	MinorVersionNumber           int                         `json:"MinorVersionNumber"`
	UpdatedTimeStamp             UpdatedTimeStampType        `json:"UpdatedTimeStamp"`
}

type BasicTestInstructionInformation []struct {
	DomainUUID                   DomainUUIDType              `json:"DomainUuid"`
	DomainName                   DomainNameType              `json:"DomainName"`
	TestInstructionUUID          OriginalElementUUIDType     `json:"TestInstructionUuid"`
	TestInstructionName          TestInstructionNameType     `json:"TestInstructionName"`
	TestInstructionTypeUUID      TestInstructionTypeUUIDType `json:"TestInstructionTypeUuid"`
	TestInstructionTypeName      TestInstructionTypeNameType `json:"TestInstructionTypeName"`
	Deprecated                   bool                        `json:"Deprecated"`
	MajorVersionNumber           int                         `json:"MajorVersionNumber"`
	MinorVersionNumber           int                         `json:"MinorVersionNumber"`
	UpdatedTimeStamp             UpdatedTimeStampType        `json:"UpdatedTimeStamp"`
	TestInstructionColor         colorType                   `json:"TestInstructionColor"`
	TCRuleDeletion               TCRuleDeletionType          `json:"TCRuleDeletion"`
	TCRuleSwap                   TCRuleSwapType              `json:"TCRuleSwap"`
	TestInstructionDescription   string                      `json:"TestInstructionDescription"`
	TestInstructionMouseOverText string                      `json:"TestInstructionMouseOverText"`
	Enabled                      bool                        `json:"Enabled"`
}

type ImmatureTestInstructionInformation []struct {
	DomainUUID                   DomainUUIDType                   `json:"DomainUuid"`
	DomainName                   DomainNameType                   `json:"DomainName"`
	TestInstructionUUID          OriginalElementUUIDType          `json:"TestInstructionUuid"`
	TestInstructionName          TestInstructionNameType          `json:"TestInstructionName"`
	DropZoneUUID                 DropZoneUUIDType                 `json:"DropZoneUuid"`
	DropZoneName                 DropZoneNameType                 `json:"DropZoneName"`
	DropZoneDescription          string                           `json:"DropZoneDescription"`
	DropZoneMouseOver            string                           `json:"DropZoneMouseOver"`
	DropZoneColor                colorType                        `json:"DropZoneColor"`
	TestInstructionAttributeType TestInstructionAttributeTypeType `json:"TestInstructionAttributeType"` //TEXTBOX...
	TestInstructionAttributeUUID TestInstructionAttributeUUIDType `json:"TestInstructionAttributeUuid"`
	TestInstructionAttributeName TestInstructionAttributeNameType `json:"TestInstructionAttributeName"`
	AttributeValueAsString       AttributeValueAsStringType       `json:"AttributeValueAsString"`
	AttributeValueUUID           AttributeValueUUIDType           `json:"AttributeValueUuid"`
	FirstImmatureElementUUID     OriginalElementUUIDType          `json:"FirstImmatureElementUuid"`
	AttributeActionCommand       AttributeActionCommandType       `json:"AttributeActionCommand"`
}

type TestInstructionAttributes []struct {
	DomainUUID                                    DomainUUIDType                        `json:"DomainUuid"`
	DomainName                                    DomainNameType                        `json:"DomainName"`
	TestInstructionUUID                           OriginalElementUUIDType               `json:"TestInstructionUuid"`
	TestInstructionName                           TestInstructionNameType               `json:"TestInstructionName"`
	TestInstructionAttributeUUID                  TestInstructionAttributeUUIDType      `json:"TestInstructionAttributeUuid"`
	TestInstructionAttributeName                  TestInstructionAttributeNameType      `json:"TestInstructionAttributeName"`
	TestInstructionAttributeDescription           string                                `json:"TestInstructionAttributeDescription"`
	TestInstructionAttributeMouseOver             string                                `json:"TestInstructionAttributeMouseOver"`
	TestInstructionAttributeTypeUUID              TestInstructionAttributeTypeUUIDType  `json:"TestInstructionAttributeTypeUuid"`
	TestInstructionAttributeTypeName              TestInstructionAttributeTypeNameType  `json:"TestInstructionAttributeTypeName"`
	TestInstructionAttributeValueAsString         AttributeValueAsStringType            `json:"TestInstructionAttributeValueAsString"`
	TestInstructionAttributeValueUUID             AttributeValueUUIDType                `json:"TestInstructionAttributeValueUuid"`
	TestInstructionAttributeVisible               bool                                  `json:"TestInstructionAttributeVisible"`
	TestInstructionAttributeEnabled               bool                                  `json:"TestInstructionAttributeEnabled"`
	TestInstructionAttributeMandatory             bool                                  `json:"TestInstructionAttributeMandatory"`
	TestInstructionAttributeVisibleInTestCaseArea bool                                  `json:"TestInstructionAttributeVisibleInTestCaseArea"`
	TestInstructionAttributeIsDeprecated          bool                                  `json:"TestInstructionAttributeIsDeprecated"`
	TestInstructionAttributeInputMask             TestInstructionAttributeInputMaskType `json:"TestInstructionAttributeInputMask"`
	TestInstructionAttributeType                  TestInstructionAttributeTypeType      `json:"TestInstructionAttributeType"` // TEXTBOX...
}

type TestInstructionContainers []struct {
	DomainUUID                            DomainUUIDType                       `json:"DomainUuid"`
	DomainName                            DomainNameType                       `json:"DomainName"`
	TestInstructionContainerUUID          OriginalElementUUIDType              `json:"TestInstructionContainerUuid"`
	TestInstructionContainerName          TestInstructionContainerNameType     `json:"TestInstructionContainerName"`
	TestInstructionContainerTypeUUID      TestInstructionContainerTypeUUIDType `json:"TestInstructionContainerTypeUuid"`
	TestInstructionContainerTypeName      TestInstructionContainerTypeNameType `json:"TestInstructionContainerTypeName"`
	TestInstructionContainerDescription   string                               `json:"TestInstructionContainerDescription"`
	TestInstructionContainerMouseOverText string                               `json:"TestInstructionContainerMouseOverText"`
	Deprecated                            bool                                 `json:"Deprecated"`
	Enabled                               bool                                 `json:"Enabled"`
	MajorVersionNumber                    int                                  `json:"MajorVersionNumber"`
	MinorVersionNumber                    int                                  `json:"MinorVersionNumber"`
	UpdatedTimeStamp                      UpdatedTimeStampType                 `json:"UpdatedTimeStamp"`
	ChildrenIsParallelProcessed           bool                                 `json:"ChildrenIsParallelProcessed"`
}

type BasicTestInstructionContainerInformation []struct {
	DomainUUID                            DomainUUIDType                            `json:"DomainUuid"`
	DomainName                            DomainNameType                            `json:"DomainName"`
	TestInstructionContainerUUID          OriginalElementUUIDType                   `json:"TestInstructionContainerUuid"`
	TestInstructionContainerName          TestInstructionContainerNameType          `json:"TestInstructionContainerName"`
	TestInstructionContainerTypeUUID      TestInstructionContainerTypeUUIDType      `json:"TestInstructionContainerTypeUuid"`
	TestInstructionContainerTypeName      TestInstructionContainerTypeNameType      `json:"TestInstructionContainerTypeName"`
	Deprecated                            bool                                      `json:"Deprecated"`
	MajorVersionNumber                    int                                       `json:"MajorVersionNumber"`
	MinorVersionNumber                    int                                       `json:"MinorVersionNumber"`
	UpdatedTimeStamp                      UpdatedTimeStampType                      `json:"UpdatedTimeStamp"`
	TestInstructionContainerColor         colorType                                 `json:"TestInstructionContainerColor"`
	TCRuleDeletion                        TCRuleDeletionType                        `json:"TCRuleDeletion"`
	TCRuleSwap                            TCRuleSwapType                            `json:"TCRuleSwap"`
	TestInstructionContainerDescription   string                                    `json:"TestInstructionContainerDescription"`
	TestInstructionContainerMouseOverText string                                    `json:"TestInstructionContainerMouseOverText"`
	Enabled                               bool                                      `json:"Enabled"`
	TestInstructionContainerExecutionType TestInstructionContainerExecutionTypeType `json:"TestInstructionContainerExecutionType"`
}

type ImmatureTestInstructionContainerMessage []struct {
	DomainUUID                   DomainUUIDType                   `json:"DomainUuid"`
	DomainName                   DomainNameType                   `json:"DomainName"`
	TestInstructionContainerUUID OriginalElementUUIDType          `json:"TestInstructionContainerUuid"`
	TestInstructionContainerName TestInstructionContainerNameType `json:"TestInstructionContainerName"`
	DropZoneUUID                 DropZoneUUIDType                 `json:"DropZoneUuid"`
	DropZoneName                 DropZoneNameType                 `json:"DropZoneName"`
	DropZoneDescription          string                           `json:"DropZoneDescription"`
	DropZoneMouseOver            string                           `json:"DropZoneMouseOver"`
	DropZoneColor                colorType                        `json:"DropZoneColor"`
	TestInstructionAttributeType TestInstructionAttributeTypeType `json:"TestInstructionAttributeType"` //TEXTBOX
	TestInstructionAttributeUUID TestInstructionAttributeUUIDType `json:"TestInstructionAttributeUuid"`
	TestInstructionAttributeName TestInstructionAttributeNameType `json:"TestInstructionAttributeName"`
	AttributeValueAsString       AttributeValueAsStringType       `json:"AttributeValueAsString"`
	AttributeValueUUID           AttributeValueUUIDType           `json:"AttributeValueUuid"`
	FirstImmatureElementUUID     OriginalElementUUIDType          `json:"FirstImmatureElementUuid"`
}

type ImmatureElementModelMessage []struct {
	DomainUUID               DomainUUIDType               `json:"DomainUuid"`
	DomainName               DomainNameType               `json:"DomainName"`
	ImmatureElementUUID      OriginalElementUUIDType      `json:"ImmatureElementUuid"`
	ImmatureElementName      OriginalElementNameType      `json:"ImmatureElementName"`
	PreviousElementUUID      OriginalElementUUIDType      `json:"PreviousElementUuid"`
	NextElementUUID          OriginalElementUUIDType      `json:"NextElementUuid"`
	FirstChildElementUUID    OriginalElementUUIDType      `json:"FirstChildElementUuid"`
	ParentElementUUID        OriginalElementUUIDType      `json:"ParentElementUuid"`
	TestCaseModelElementType TestCaseModelElementTypeType `json:"TestCaseModelElementType"`
	OriginalElementUUID      OriginalElementUUIDType      `json:"OriginalElementUuid"`
	TopImmatureElementUUID   OriginalElementUUIDType      `json:"TopImmatureElementUuid"`
	IsTopElement             bool                         `json:"IsTopElement"`
}

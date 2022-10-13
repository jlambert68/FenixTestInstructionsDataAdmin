package TypeAndStructs

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
package TestInstructionContainers

import (
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (
	// *** TestInstructionContainer 'EmptySerialContainer'
	TestInstructionContainerUUID_Fenix_EmptySerialContainer                        TypeAndStructs.OriginalElementUUIDType          = "f81b9734-5dce-43c9-8d77-3368940cf126"
	TestInstructionContainerName_Fenix_EmptySerialContainer                        TypeAndStructs.TestInstructionContainerNameType = "Empty serial processed TestInstructionsContainer"
	TestInstructionContainerTypeUUID_Fenix_EmptySerialContainer                                                                    = TestInstructionContainerTypeUUID_Fenix_BaseContainers
	TestInstructionContainerTypeName_Fenix_EmptySerialContainer                                                                    = TestInstructionContainerTypeNameType_Fenix_BaseContainers
	TestInstructionContainerDescription_Fenix_EmptySerialContainer                 string                                          = ""
	TestInstructionContainerMouseOverText_Fenix_EmptySerialContainer               string                                          = ""
	TestInstructionContainerDeprecated_Fenix_EmptySerialContainer                  bool                                            = false
	TestInstructionContainerEnabled_Fenix_EmptySerialContainer                     bool                                            = true
	TestInstructionContainerMajorVersionNumber_Fenix_EmptySerialContainer          int                                             = 1
	TestInstructionContainerMinorVersionNumber_Fenix_EmptySerialContainer          int                                             = 0
	TestInstructionContainerChildrenIsParallelProcessed_Fenix_EmptySerialContainer bool                                            = false
	TestInstructionContainerColor_Fenix_EmptySerialContainer                       TypeAndStructs.ColorType                        = "#AAAAAA"
	TCRuleDeletion_Fenix_EmptySerialContainer                                      TypeAndStructs.TCRuleDeletionType               = "TCRuleDeletion011"
	TCRuleSwap_CA_Fenix_EmptySerialContainer                                       TypeAndStructs.TCRuleSwapType                   = "TCRuleSwap012"

	// *** DropZone ***
	TestInstructionContainerDropZoneUUID_Fenix_EmptySerialContainer TypeAndStructs.DropZoneUUIDType = "c5e37024-e40c-49f7-8667-eab485c65105"
	TestInstructionContainerDropZoneName_Fenix_EmptySerialContainer TypeAndStructs.DropZoneUUIDType = "My first DropZone for a TestInstructionContainer"
)

// TestInstructionContainer_Fenix_SerialStruct
// Struct for holding all data for the TestInstructionContainer
type TestInstructionContainer_Fenix_SerialStruct struct {
	TestInstructionContainer                 TypeAndStructs.TestInstructionContainerStruct
	BasicTestInstructionContainerInformation TypeAndStructs.BasicTestInstructionContainerInformationStruct
	ImmatureTestInstructionContainer         []TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
	ImmatureElementModel                     []TypeAndStructs.ImmatureElementModelMessageStruct
}

// TestInstructionContainer_Fenix_Serial
// Variable that holds the data for the TestInstructionContainer
var TestInstructionContainer_Fenix_Serial TestInstructionContainer_Fenix_SerialStruct

// Initiate_TestInstructionContainer_Fenix_Serial
// Function that creates all data for the TestInstructionContainer
func Initiate_TestInstructionContainer_Fenix_Serial() {

	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstructionContainer - 'EmptySerialContainer'
	TestInstructionContainer_Fenix_Serial.TestInstructionContainer = TypeAndStructs.TestInstructionContainerStruct{
		DomainUUID:                            Domains.DomainUUID_Fenix,
		DomainName:                            Domains.DomainName_Fenix,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		TestInstructionContainerName:          TestInstructionContainerName_Fenix_EmptySerialContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_Fenix_EmptySerialContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_Fenix_EmptySerialContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_Fenix_EmptySerialContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_Fenix_EmptySerialContainer,
		Deprecated:                            TestInstructionContainerDeprecated_Fenix_EmptySerialContainer,
		Enabled:                               TestInstructionContainerEnabled_Fenix_EmptySerialContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_Fenix_EmptySerialContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_Fenix_EmptySerialContainer,
		UpdatedTimeStamp:                      updatedTimeStamp,
		ChildrenIsParallelProcessed:           TestInstructionContainerChildrenIsParallelProcessed_Fenix_EmptySerialContainer,
	}

	TestInstructionContainer_Fenix_Serial.BasicTestInstructionContainerInformation = TypeAndStructs.BasicTestInstructionContainerInformationStruct{
		DomainUUID:                            Domains.DomainUUID_Fenix,
		DomainName:                            Domains.DomainName_Fenix,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		TestInstructionContainerName:          TestInstructionContainerName_Fenix_EmptySerialContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_Fenix_EmptySerialContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_Fenix_EmptySerialContainer,
		Deprecated:                            TestInstructionContainerDeprecated_Fenix_EmptySerialContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_Fenix_EmptySerialContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_Fenix_EmptySerialContainer,
		UpdatedTimeStamp:                      updatedTimeStamp,
		TestInstructionContainerColor:         TestInstructionContainerColor_Fenix_EmptySerialContainer,
		TCRuleDeletion:                        TCRuleDeletion_Fenix_EmptySerialContainer,
		TCRuleSwap:                            TCRuleSwap_CA_Fenix_EmptySerialContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_Fenix_EmptySerialContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_Fenix_EmptySerialContainer,
		Enabled:                               TestInstructionContainerEnabled_Fenix_EmptySerialContainer,
		TestInstructionContainerExecutionType: Domains.TestInstructionContainerExecutionType_SERIAL_PROCESSED,
	}
}

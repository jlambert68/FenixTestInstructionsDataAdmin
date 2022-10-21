package TestInstructionContainers

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Bonds"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

const (
	// *** TestInstructionContainer 'EmptySerialContainer'
	TestInstructionContainerUUID_Fenix_EmptySerialContainer                        TypeAndStructs.OriginalElementUUIDType          = "f81b9734-5dce-43c9-8d77-3368940cf126"
	TestInstructionContainerName_Fenix_EmptySerialContainer                        TypeAndStructs.TestInstructionContainerNameType = "Empty serial processed TestInstructionsContainer"
	TestInstructionContainerTypeUUID_Fenix_EmptySerialContainer                                                                    = TestInstructionContainerTypeUUID_Fenix_BaseContainers
	TestInstructionContainerTypeName_Fenix_EmptySerialContainer                                                                    = TestInstructionContainerTypeNameType_Fenix_BaseContainers
	TestInstructionContainerDescription_Fenix_EmptySerialContainer                 string                                          = "Children of this container is processed in serial"
	TestInstructionContainerMouseOverText_Fenix_EmptySerialContainer               string                                          = "Children of this container is processed in serial"
	TestInstructionContainerDeprecated_Fenix_EmptySerialContainer                  bool                                            = false
	TestInstructionContainerEnabled_Fenix_EmptySerialContainer                     bool                                            = true
	TestInstructionContainerMajorVersionNumber_Fenix_EmptySerialContainer          int                                             = 1
	TestInstructionContainerMinorVersionNumber_Fenix_EmptySerialContainer          int                                             = 0
	TestInstructionContainerChildrenIsParallelProcessed_Fenix_EmptySerialContainer bool                                            = false
	TestInstructionContainerColor_Fenix_EmptySerialContainer                       TypeAndStructs.ColorType                        = "#AAAAAA"
	TCRuleDeletion_Fenix_EmptySerialContainer                                      TypeAndStructs.TCRuleDeletionType               = "TCRuleDeletion011"
	TCRuleSwap_CA_Fenix_EmptySerialContainer                                       TypeAndStructs.TCRuleSwapType                   = "TCRuleSwap012"

	// *** DropZone ***
	// No DropZone for 'EmptySerialContainer'
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

	// BasicTestInstructionContainerInformation - 'EmptySerialContainer'
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

	// ImmatureTestInstructionContainerMessage - 'EmptySerialContainer'
	// No DropZone for 'EmptySerialContainer'

	// ImmatureElementModelMessage - 'EmptySerialContainer' - 'TIC' in 'TIC(B10)'
	var ImmatureElementModel_TIC TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIC = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_Fenix,
		DomainName:               Domains.DomainName_Fenix,
		ImmatureElementUUID:      TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionContainerName_Fenix_EmptySerialContainer),
		PreviousElementUUID:      TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		NextElementUUID:          TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TIC,
		OriginalElementUUID:      TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		TopImmatureElementUUID:   TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		IsTopElement:             true,
	}
	TestInstructionContainer_Fenix_Serial.ImmatureElementModel = append(TestInstructionContainer_Fenix_Serial.ImmatureElementModel, ImmatureElementModel_TIC)

	// ImmatureElementModelMessage - 'EmptySerialContainer' - 'B10' in 'TIC(B10)'
	var ImmatureElementModel_B10 TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B10 = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_Fenix,
		DomainName:               Domains.DomainName_Fenix,
		ImmatureElementUUID:      Bonds.Bond_B10_BondUuid,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(Bonds.Bond_B0_BondName),
		PreviousElementUUID:      Bonds.Bond_B10_BondUuid,
		NextElementUUID:          Bonds.Bond_B10_BondUuid,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B10,
		OriginalElementUUID:      Bonds.Bond_B10_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_Fenix_EmptySerialContainer,
		IsTopElement:             false,
	}
	TestInstructionContainer_Fenix_Serial.ImmatureElementModel = append(TestInstructionContainer_Fenix_Serial.ImmatureElementModel, ImmatureElementModel_B10)

}

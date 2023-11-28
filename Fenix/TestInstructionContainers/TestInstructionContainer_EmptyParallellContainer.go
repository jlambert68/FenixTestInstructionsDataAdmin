package TestInstructionContainers

import (
	"FenixTestInstructionsDataAdmin/Bonds"
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstructionContainer ***
	TestInstructionContainerUUID_Fenix_EmptyParallellContainer                        TypeAndStructs.OriginalElementUUIDType          = "aa1b9734-5dce-43c9-8d77-3368940cf126"
	TestInstructionContainerName_Fenix_EmptyParallellContainer                        TypeAndStructs.TestInstructionContainerNameType = "Emtpy parallelled processed TestInstructionsContainer"
	TestInstructionContainerTypeUUID_Fenix_EmptyParallellContainer                                                                    = TestInstructionContainerTypeUUID_Fenix_BaseContainers
	TestInstructionContainerTypeName_Fenix_EmptyParallellContainer                                                                    = TestInstructionContainerTypeNameType_Fenix_BaseContainers
	TestInstructionContainerDescription_Fenix_EmptyParallellContainer                 string                                          = "Children of this container is processed in parallell"
	TestInstructionContainerMouseOverText_Fenix_EmptyParallellContainer               string                                          = "Children of this container is processed in parallell"
	TestInstructionContainerDeprecated_Fenix_EmptyParallellContainer                  bool                                            = false
	TestInstructionContainerEnabled_Fenix_EmptyParallellContainer                     bool                                            = true
	TestInstructionContainerMajorVersionNumber_Fenix_EmptyParallellContainer          int                                             = 1
	TestInstructionContainerMinorVersionNumber_Fenix_EmptyParallellContainer          int                                             = 0
	TestInstructionContainerChildrenIsParallelProcessed_Fenix_EmptyParallellContainer bool                                            = true
	TestInstructionContainerColor_Fenix_EmptyParallellContainer                       TypeAndStructs.ColorType                        = "#AAAAAA"
	TCRuleDeletion_Fenix_EmptyParallellContainer                                      TypeAndStructs.TCRuleDeletionType               = "TCRuleDeletion011"
	TCRuleSwap_CA_Fenix_EmptyParallellContainer                                       TypeAndStructs.TCRuleSwapType                   = "TCRuleSwap012"

	// *** DropZone ***
	// No DropZone for 'EmptyParallellContainer'
)

// TestInstructionContainer_Fenix_ParallellStruct
// Struct for holding all data for the TestInstructionContainer
type TestInstructionContainer_Fenix_ParallellStruct struct {
	TestInstructionContainer                 TypeAndStructs.TestInstructionContainerStruct
	BasicTestInstructionContainerInformation TypeAndStructs.BasicTestInstructionContainerInformationStruct
	ImmatureTestInstructionContainer         []TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
	ImmatureElementModel                     []TypeAndStructs.ImmatureElementModelMessageStruct
}

// TestInstructionContainer_Fenix_Parallell
// Variable that holds the data for the TestInstructionContainer
var TestInstructionContainer_Fenix_Parallell TestInstructionContainer_Fenix_ParallellStruct

// Initiate_TestInstructionContainer_Fenix_Parallell
// Function that creates all data for the TestInstructionContainer
func Initiate_TestInstructionContainer_Fenix_Parallell() {

	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstructionContainer - 'EmptyParallellContainer'
	TestInstructionContainer_Fenix_Parallell.TestInstructionContainer = TypeAndStructs.TestInstructionContainerStruct{
		DomainUUID:                            Domains.DomainUUID_Fenix,
		DomainName:                            Domains.DomainName_Fenix,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		TestInstructionContainerName:          TestInstructionContainerName_Fenix_EmptyParallellContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_Fenix_EmptyParallellContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_Fenix_EmptyParallellContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_Fenix_EmptyParallellContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_Fenix_EmptyParallellContainer,
		Deprecated:                            TestInstructionContainerDeprecated_Fenix_EmptyParallellContainer,
		Enabled:                               TestInstructionContainerEnabled_Fenix_EmptyParallellContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_Fenix_EmptyParallellContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_Fenix_EmptyParallellContainer,
		UpdatedTimeStamp:                      updatedTimeStamp,
		ChildrenIsParallelProcessed:           TestInstructionContainerChildrenIsParallelProcessed_Fenix_EmptyParallellContainer,
	}

	// BasicTestInstructionContainerInformation - 'EmptyParallellContainer'
	TestInstructionContainer_Fenix_Parallell.BasicTestInstructionContainerInformation = TypeAndStructs.BasicTestInstructionContainerInformationStruct{
		DomainUUID:                            Domains.DomainUUID_Fenix,
		DomainName:                            Domains.DomainName_Fenix,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		TestInstructionContainerName:          TestInstructionContainerName_Fenix_EmptyParallellContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_Fenix_EmptyParallellContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_Fenix_EmptyParallellContainer,
		Deprecated:                            TestInstructionContainerDeprecated_Fenix_EmptyParallellContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_Fenix_EmptyParallellContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_Fenix_EmptyParallellContainer,
		UpdatedTimeStamp:                      updatedTimeStamp,
		TestInstructionContainerColor:         TestInstructionContainerColor_Fenix_EmptyParallellContainer,
		TCRuleDeletion:                        TCRuleDeletion_Fenix_EmptyParallellContainer,
		TCRuleSwap:                            TCRuleSwap_CA_Fenix_EmptyParallellContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_Fenix_EmptyParallellContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_Fenix_EmptyParallellContainer,
		Enabled:                               TestInstructionContainerEnabled_Fenix_EmptyParallellContainer,
		TestInstructionContainerExecutionType: Domains.TestInstructionContainerExecutionType_PARALLELLED_PROCESSED,
	}

	// ImmatureTestInstructionContainerMessage - 'EmptyParallellContainer'
	// No DropZone for 'EmptyParallellContainer'

	// ImmatureElementModelMessage - 'EmptyParallellContainer' - 'TIC' in 'TIC(B10)'
	var ImmatureElementModel_TIC TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIC = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_Fenix,
		DomainName:               Domains.DomainName_Fenix,
		ImmatureElementUUID:      TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionContainerName_Fenix_EmptyParallellContainer),
		PreviousElementUUID:      TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		NextElementUUID:          TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TIC,
		OriginalElementUUID:      TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		TopImmatureElementUUID:   TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		IsTopElement:             true,
	}
	TestInstructionContainer_Fenix_Parallell.ImmatureElementModel = append(TestInstructionContainer_Fenix_Parallell.ImmatureElementModel, ImmatureElementModel_TIC)

	// ImmatureElementModelMessage - 'EmptyParallellContainer' - 'B10' in 'TIC(B10)'
	var ImmatureElementModel_B10 TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B10 = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_Fenix,
		DomainName:               Domains.DomainName_Fenix,
		ImmatureElementUUID:      Bonds.Bond_B10_BondUuid,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(Bonds.Bond_B0_BondName),
		PreviousElementUUID:      Bonds.Bond_B10_BondUuid,
		NextElementUUID:          Bonds.Bond_B10_BondUuid,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B10,
		OriginalElementUUID:      Bonds.Bond_B10_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_Fenix_EmptyParallellContainer,
		IsTopElement:             false,
	}
	TestInstructionContainer_Fenix_Parallell.ImmatureElementModel = append(TestInstructionContainer_Fenix_Parallell.ImmatureElementModel, ImmatureElementModel_B10)

}

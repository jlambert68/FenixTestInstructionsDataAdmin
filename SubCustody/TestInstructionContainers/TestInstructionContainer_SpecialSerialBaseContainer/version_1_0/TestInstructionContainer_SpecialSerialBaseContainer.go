package version_1_0

import (
	"FenixTestInstructionsDataAdmin/Bonds"
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/SubCustody/DomainData"
	"FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers"
	fixedValuesOverVersions "FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer"
	"FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	"FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	"FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"FenixTestInstructionsDataAdmin/TestInstructionAndTestInstuctionContainerTypes"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
)

const (
	// *** TestInstructionContainer 'SpecialSerialBaseContainer'
	TestInstructionContainerUUID_SC_SpecialSerialBaseContainer                        TypeAndStructs.OriginalElementUUIDType          = fixedValuesOverVersions.TestInstructionContainerUUID_SC_SpecialSerialBaseContainer
	TestInstructionContainerName_SC_SpecialSerialBaseContainer                        TypeAndStructs.TestInstructionContainerNameType = "Empty serial processed TestInstructionsContainer"
	TestInstructionContainerTypeUUID_SC_SpecialSerialBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeUUID_SC_BaseContainers
	TestInstructionContainerTypeName_SC_SpecialSerialBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeNameType_SC_BaseContainers
	TestInstructionContainerDescription_SC_SpecialSerialBaseContainer                 string                                          = "Children of this container is processed in serial"
	TestInstructionContainerMouseOverText_SC_SpecialSerialBaseContainer               string                                          = "Children of this container is processed in serial"
	TestInstructionContainerDeprecated_SC_SpecialSerialBaseContainer                  bool                                            = false
	TestInstructionContainerEnabled_SC_SpecialSerialBaseContainer                     bool                                            = true
	TestInstructionContainerMajorVersionNumber_SC_SpecialSerialBaseContainer          int                                             = 1
	TestInstructionContainerMinorVersionNumber_SC_SpecialSerialBaseContainer          int                                             = 0
	TestInstructionContainerChildrenIsParallelProcessed_SC_SpecialSerialBaseContainer bool                                            = false
	TestInstructionContainerColor_SC_SpecialSerialBaseContainer                       TypeAndStructs.ColorType                        = "#AAAAAA"
	TCRuleDeletion_SC_SpecialSerialBaseContainer                                      TypeAndStructs.TCRuleDeletionType               = "TCRuleDeletion011"
	TCRuleSwap_SC_SpecialSerialBaseContainer                                          TypeAndStructs.TCRuleSwapType                   = "TCRuleSwap012"
	TestInstructionContainerCreatingTimeStamp                                         TypeAndStructs.UpdatedTimeStampType             = "2023-11-28 14:00:00"

	// *** DropZone ***
	// No DropZone for 'SpecialSerialBaseContainer'
)

// TestInstructionContainer_SC_SpecialSerialBase
// Variable that holds the data for the TestInstructionContainer
var TestInstructionContainer_SC_SpecialSerialBase *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct

// Initiate_TestInstructionContainer_SC_Serial
// Function that creates all data for the TestInstructionContainer
func Initiate_TestInstructionContainer_SC_Serial(testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) {

	// Initiate the full structure
	TestInstructionContainer_SC_SpecialSerialBase = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct{}

	// TestInstructionContainer - 'SpecialSerialBaseContainer'
	TestInstructionContainer_SC_SpecialSerialBase.TestInstructionContainer = &TypeAndStructs.TestInstructionContainerStruct{
		DomainUUID:                            DomainData.DomainUUID_Fenix,
		DomainName:                            DomainData.DomainName_Fenix,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_SC_SpecialSerialBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_SC_SpecialSerialBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_SC_SpecialSerialBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_SC_SpecialSerialBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_SC_SpecialSerialBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_SC_SpecialSerialBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_SC_SpecialSerialBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_SC_SpecialSerialBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_SC_SpecialSerialBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		ChildrenIsParallelProcessed:           TestInstructionContainerChildrenIsParallelProcessed_SC_SpecialSerialBaseContainer,
	}

	// BasicTestInstructionContainerInformation - 'SpecialSerialBaseContainer'
	TestInstructionContainer_SC_SpecialSerialBase.BasicTestInstructionContainerInformation = &TypeAndStructs.BasicTestInstructionContainerInformationStruct{
		DomainUUID:                            DomainData.DomainUUID_Fenix,
		DomainName:                            DomainData.DomainName_Fenix,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_SC_SpecialSerialBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_SC_SpecialSerialBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_SC_SpecialSerialBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_SC_SpecialSerialBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_SC_SpecialSerialBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_SC_SpecialSerialBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		TestInstructionContainerColor:         TestInstructionContainerColor_SC_SpecialSerialBaseContainer,
		TCRuleDeletion:                        TCRuleDeletion_SC_SpecialSerialBaseContainer,
		TCRuleSwap:                            TCRuleSwap_SC_SpecialSerialBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_SC_SpecialSerialBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_SC_SpecialSerialBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_SC_SpecialSerialBaseContainer,
		TestInstructionContainerExecutionType: Domains.TestInstructionContainerExecutionType_SERIAL_PROCESSED,
	}

	// ImmatureTestInstructionContainerMessage - 'SpecialSerialBaseContainer'
	// No DropZone for 'SpecialSerialBaseContainer'

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TIC' in 'TIC(B11fx-TI-B12-TI-B11lx)'
	var ImmatureElementModel_TIC *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIC = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_Fenix,
		DomainName:               DomainData.DomainName_Fenix,
		ImmatureElementUUID:      TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionContainerName_SC_SpecialSerialBaseContainer),
		PreviousElementUUID:      TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		NextElementUUID:          TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		FirstChildElementUUID:    Bonds.Bond_B10_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TIC,
		OriginalElementUUID:      TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TopImmatureElementUUID:   TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:             true,
	}
	TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_TIC)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B11fx_1' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B11fx_1 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B11fx_1 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          DomainData.DomainUUID_SC,
		DomainName:          DomainData.DomainName_SC,
		ImmatureElementUUID: Bonds.Bond_B11fx_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B11fx_BondName),
		PreviousElementUUID: Bonds.Bond_B11fx_BondUuid,
		NextElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		FirstChildElementUUID:    Bonds.Bond_B11fx_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B11fx,
		OriginalElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_B11fx_1)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TI_1' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_TI_1 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TI_1 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID: DomainData.DomainUUID_SC,
		DomainName: DomainData.DomainName_SC,
		ImmatureElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName),
		PreviousElementUUID: Bonds.Bond_B11fx_BondUuid,
		NextElementUUID:     Bonds.Bond_B10_BondUuid,
		FirstChildElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_TI_1)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B12' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B12 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B12 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          DomainData.DomainUUID_Fenix,
		DomainName:          DomainData.DomainName_Fenix,
		ImmatureElementUUID: Bonds.Bond_B12_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B0_BondName),
		PreviousElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		NextElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		FirstChildElementUUID:    Bonds.Bond_B12_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B10,
		OriginalElementUUID:      Bonds.Bond_B12_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:             false,
	}
	TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_B12)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TI_2' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_TI_2 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TI_2 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID: DomainData.DomainUUID_SC,
		DomainName: DomainData.DomainName_SC,
		ImmatureElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName),
		PreviousElementUUID: Bonds.Bond_B12_BondUuid,
		NextElementUUID:     Bonds.Bond_B11lx_BondUuid,
		FirstChildElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_TI_2)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B10' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B11lx *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B11lx = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          DomainData.DomainUUID_Fenix,
		DomainName:          DomainData.DomainName_Fenix,
		ImmatureElementUUID: Bonds.Bond_B11lx_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B0_BondName),
		PreviousElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		NextElementUUID:          Bonds.Bond_B11lx_BondUuid,
		FirstChildElementUUID:    Bonds.Bond_B11lx_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B10,
		OriginalElementUUID:      Bonds.Bond_B11lx_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:             false,
	}
	TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_SC_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_B11lx)

}

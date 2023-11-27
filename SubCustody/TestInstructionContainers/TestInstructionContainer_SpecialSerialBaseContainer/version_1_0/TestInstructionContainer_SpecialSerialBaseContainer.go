package version_1_0

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Bonds"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers"
	fixedValuesOverVersions "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
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

	// *** DropZone ***
	// No DropZone for 'SpecialSerialBaseContainer'
)

// TestInstructionContainer_SC_SerialStruct
// Struct for holding all data for the TestInstructionContainer
type TestInstructionContainer_SC_SerialStruct struct {
	TestInstructionContainer                 TypeAndStructs.TestInstructionContainerStruct
	BasicTestInstructionContainerInformation TypeAndStructs.BasicTestInstructionContainerInformationStruct
	ImmatureTestInstructionContainer         []TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
	ImmatureElementModel                     []TypeAndStructs.ImmatureElementModelMessageStruct
}

// TestInstructionContainer_SC_Serial
// Variable that holds the data for the TestInstructionContainer
var TestInstructionContainer_SC_Serial TestInstructionContainer_SC_SerialStruct

// Initiate_TestInstructionContainer_SC_Serial
// Function that creates all data for the TestInstructionContainer
func Initiate_TestInstructionContainer_SC_Serial() {

	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstructionContainer - 'SpecialSerialBaseContainer'
	TestInstructionContainer_SC_Serial.TestInstructionContainer = TypeAndStructs.TestInstructionContainerStruct{
		DomainUUID:                            Domains.DomainUUID_Fenix,
		DomainName:                            Domains.DomainName_Fenix,
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
		UpdatedTimeStamp:                      updatedTimeStamp,
		ChildrenIsParallelProcessed:           TestInstructionContainerChildrenIsParallelProcessed_SC_SpecialSerialBaseContainer,
	}

	// BasicTestInstructionContainerInformation - 'SpecialSerialBaseContainer'
	TestInstructionContainer_SC_Serial.BasicTestInstructionContainerInformation = TypeAndStructs.BasicTestInstructionContainerInformationStruct{
		DomainUUID:                            Domains.DomainUUID_Fenix,
		DomainName:                            Domains.DomainName_Fenix,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_SC_SpecialSerialBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_SC_SpecialSerialBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_SC_SpecialSerialBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_SC_SpecialSerialBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_SC_SpecialSerialBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_SC_SpecialSerialBaseContainer,
		UpdatedTimeStamp:                      updatedTimeStamp,
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
	var ImmatureElementModel_TIC TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIC = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_Fenix,
		DomainName:               Domains.DomainName_Fenix,
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
	TestInstructionContainer_SC_Serial.ImmatureElementModel = append(TestInstructionContainer_SC_Serial.ImmatureElementModel, ImmatureElementModel_TIC)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B11fx_1' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B11fx_1 TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B11fx_1 = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          Domains.DomainUUID_SC,
		DomainName:          Domains.DomainName_SC,
		ImmatureElementUUID: Bonds.Bond_B11fx_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B11fx_BondName),
		PreviousElementUUID: Bonds.Bond_B11fx_BondUuid,
		NextElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		FirstChildElementUUID:    Bonds.Bond_B11fx_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B11fx,
		OriginalElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_SC_Serial.ImmatureElementModel = append(TestInstructionContainer_SC_Serial.ImmatureElementModel, ImmatureElementModel_B11fx_1)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TI_1' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_TI_1 TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TI_1 = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID: Domains.DomainUUID_SC,
		DomainName: Domains.DomainName_SC,
		ImmatureElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName),
		PreviousElementUUID: Bonds.Bond_B11fx_BondUuid,
		NextElementUUID:     Bonds.Bond_B10_BondUuid,
		FirstChildElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_SC_Serial.ImmatureElementModel = append(TestInstructionContainer_SC_Serial.ImmatureElementModel, ImmatureElementModel_TI_1)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B12' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B12 TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B12 = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          Domains.DomainUUID_Fenix,
		DomainName:          Domains.DomainName_Fenix,
		ImmatureElementUUID: Bonds.Bond_B12_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B0_BondName),
		PreviousElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		NextElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		FirstChildElementUUID:    Bonds.Bond_B12_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B10,
		OriginalElementUUID:      Bonds.Bond_B12_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:             false,
	}
	TestInstructionContainer_SC_Serial.ImmatureElementModel = append(TestInstructionContainer_SC_Serial.ImmatureElementModel, ImmatureElementModel_B12)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TI_2' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_TI_2 TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TI_2 = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID: Domains.DomainUUID_SC,
		DomainName: Domains.DomainName_SC,
		ImmatureElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName),
		PreviousElementUUID: Bonds.Bond_B12_BondUuid,
		NextElementUUID:     Bonds.Bond_B11lx_BondUuid,
		FirstChildElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ParentElementUUID:        TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_SC_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_SC_Serial.ImmatureElementModel = append(TestInstructionContainer_SC_Serial.ImmatureElementModel, ImmatureElementModel_TI_2)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B10' in 'TIC(B11fx_1-TI_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B11lx TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B11lx = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          Domains.DomainUUID_Fenix,
		DomainName:          Domains.DomainName_Fenix,
		ImmatureElementUUID: Bonds.Bond_B11lx_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B0_BondName),
		PreviousElementUUID: SubCustody.TestInstructions_SC.TestInstructions.
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
	TestInstructionContainer_SC_Serial.ImmatureElementModel = append(TestInstructionContainer_SC_Serial.ImmatureElementModel, ImmatureElementModel_B11lx)

}

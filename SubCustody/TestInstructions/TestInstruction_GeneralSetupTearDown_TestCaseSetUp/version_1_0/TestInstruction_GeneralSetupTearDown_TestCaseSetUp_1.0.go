package version_1_0

import (
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/LocalExecutionMethods"
	"FenixTestInstructionsDataAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
	"FenixTestInstructionsDataAdmin/SubCustody/DomainData"
	"FenixTestInstructionsDataAdmin/SubCustody/TestInstructions"
	fixedValuesOverVersions "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	"FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"FenixTestInstructionsDataAdmin/TestInstructionAndTestInstuctionContainerTypes"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseSetUp'
	TestInstructionUUID_SC_TestCaseSetUp               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SC_TestCaseSetUp
	TestInstructionName_SC_TestCaseSetUp               TypeAndStructs.TestInstructionNameType = "TestCaseSetUp"
	TestInstructionTypeUUID_SC_TestCaseSetUp                                                  = TestInstructions.TestInstructionTypeUUID_SC_GeneralSetUpTearDown
	TestInstructionTypeName_SC_TestCaseSetUp                                                  = TestInstructions.TestInstructionTypeName_SC_GeneralSetUpTearDown
	TestInstructionDescription_SC_TestCaseSetUp        string                                 = "Initiate _SCs execution engine to be able to execute TestInstructionsMap"
	TestInstructionMouseOverText_SC_TestCaseSetUp      string                                 = "Initiate _SCs execution engine to be able to execute TestInstructionsMap"
	TestInstructionDeprecated_SC_TestCaseSetUp         bool                                   = true
	TestInstructionEnabled_SC_TestCaseSetUp            bool                                   = true
	TestInstructionMajorVersionNumber_SC_TestCaseSetUp int                                    = 1
	TestInstructionMinorVersionNumber_SC_TestCaseSetUp int                                    = 0
	TestInstructionColor_SC_TestCaseSetUp              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SC_TestCaseSetUp                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SC_TestCaseSetUp                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                   TypeAndStructs.UpdatedTimeStampType    = "2023-11-27 13:00:00"

	// *** DropZone *** 'TestCaseSetUp_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SC_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "60a2492f-9b0d-4a0e-a374-5845ffabf37d"
	TestInstructionDropZoneName_SC_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseSetUp_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SC_TestCaseSetUp_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SC_TestCaseSetUp_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SC_TestCaseSetUp_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "f4682904-8f60-447c-b851-e713f2b4a03d" //TestInstructionAttributeUUID_SC_ExpectedToBePassed
	TestInstructionAttributeName_SC_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SC_ExpectedToBePassed
	TestInstructionAttributeType_SC_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SC_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SC_TestCaseSetUp_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SC_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SC_TestCaseSetUp_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SC_TestCaseSetUp_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SC_TestCaseSetUp_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_SC_TestCaseSetUp
// Variable that holds the data for the TestInstruction
var TestInstruction_SC_TestCaseSetUp *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SC_TestCaseSetUp
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SC_TestCaseSetUp() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to store all TestInstruction data
	TestInstruction_SC_TestCaseSetUp = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: &LocalExecutionMethods.MethodsForLocalExecutionsStruct{
			FangEngineClassesMethodsAttributes: &FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
				TestInstructionOriginalUUID: TestInstructionUUID_SC_TestCaseSetUp,
				TestInstructionName:         TestInstructionName_SC_TestCaseSetUp,
				FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralSetupTearDown,
				FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralSetupTearDown,
				FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_SC_GeneralSetupTearDown_Setup,
				FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_SC_GeneralSetupTearDown_Setup,
				Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct),
			},
		},
	}

	// TestInstruction - TestCaseSetUp
	TestInstruction_SC_TestCaseSetUp.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SC,
		DomainName:                   DomainData.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_SC_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SC_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_SC_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_SC_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SC_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_SC_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_SC_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SC_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SC_TestCaseSetUp,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseSetUp
	TestInstruction_SC_TestCaseSetUp.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SC,
		DomainName:                   DomainData.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_SC_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SC_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_SC_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_SC_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SC_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SC_TestCaseSetUp,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SC_TestCaseSetUp,
		TCRuleDeletion:               TCRuleDeletion_SC_TestCaseSetUp,
		TCRuleSwap:                   TCRuleSwap_SC_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_SC_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SC_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_SC_TestCaseSetUp,
	}

	// DropZone 'TestCaseSetUp_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseSetUp_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SC_TestCaseSetUp_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SC_TestCaseSetUp_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SC,
		DomainName:                   DomainData.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_SC_TestCaseSetUp,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SC_TestCaseSetUp_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SC_TestCaseSetUp_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SC_TestCaseSetUp_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SC_TestCaseSetUp_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SC_TestCaseSetUp_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SC_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SC_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SC_TestCaseSetUp_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SC_TestCaseSetUp,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SC_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		TestInstruction_SC_TestCaseSetUp.ImmatureTestInstructionInformation,
		TestInstruction_SC_TestCaseSetUp_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SC_TestCaseSetUp_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SC_TestCaseSetUp_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SC,
		DomainName:                                    DomainData.DomainName_SC,
		TestInstructionUUID:                           TestInstructionUUID_SC_TestCaseSetUp,
		TestInstructionName:                           TestInstructionName_SC_TestCaseSetUp,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_SC_ExpectedToPass,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_SC_ExpectedToPass,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_SC_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseSetUp.TestInstructionAttribute = append(
		TestInstruction_SC_TestCaseSetUp.TestInstructionAttribute,
		TestInstructionAttribute_SC_TestCaseSetUp_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed *FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = &FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_SC_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseSetUp.LocalExecutionMethods.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseSetUp
	var TestInstructionImmatureElementModel_SC_TestCaseSetUp *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SC_TestCaseSetUp = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SC,
		DomainName:               DomainData.DomainName_SC,
		ImmatureElementUUID:      TestInstructionUUID_SC_TestCaseSetUp,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SC_TestCaseSetUp),
		PreviousElementUUID:      TestInstructionUUID_SC_TestCaseSetUp,
		NextElementUUID:          TestInstructionUUID_SC_TestCaseSetUp,
		FirstChildElementUUID:    TestInstructionUUID_SC_TestCaseSetUp,
		ParentElementUUID:        TestInstructionUUID_SC_TestCaseSetUp,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SC_TestCaseSetUp,
		TopImmatureElementUUID:   TestInstructionUUID_SC_TestCaseSetUp,
		IsTopElement:             true,
	}
	TestInstruction_SC_TestCaseSetUp.ImmatureElementModel = append(
		TestInstruction_SC_TestCaseSetUp.ImmatureElementModel,
		TestInstructionImmatureElementModel_SC_TestCaseSetUp)

	return TestInstruction_SC_TestCaseSetUp
}

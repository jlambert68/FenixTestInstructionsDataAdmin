package version_1_0

import (
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/LocalExecutionMethods"
	"FenixTestInstructionsDataAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
	"FenixTestInstructionsDataAdmin/SubCustody/DomainData"
	"FenixTestInstructionsDataAdmin/SubCustody/TestInstructions"
	fixedValuesOverVersions "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	"FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"FenixTestInstructionsDataAdmin/TestInstructionAndTestInstuctionContainerTypes"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseTearDown'
	TestInstructionUUID_SC_TestCaseTearDown               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SC_TestCaseTearDown
	TestInstructionName_SC_TestCaseTearDown               TypeAndStructs.TestInstructionNameType = "TestCaseTearDown"
	TestInstructionTypeUUID_SC_TestCaseTearDown                                                  = TestInstructions.TestInstructionTypeUUID_SC_GeneralSetUpTearDown
	TestInstructionTypeName_SC_TestCaseTearDown                                                  = TestInstructions.TestInstructionTypeName_SC_GeneralSetUpTearDown
	TestInstructionDescription_SC_TestCaseTearDown        string                                 = "TearDown the _SCs execution engine after executing TestInstructionsMap"
	TestInstructionMouseOverText_SC_TestCaseTearDown      string                                 = "TearDown the _SCs execution engine after executing TestInstructionsMap"
	TestInstructionDeprecated_SC_TestCaseTearDown         bool                                   = false
	TestInstructionEnabled_SC_TestCaseTearDown            bool                                   = true
	TestInstructionMajorVersionNumber_SC_TestCaseTearDown int                                    = 1
	TestInstructionMinorVersionNumber_SC_TestCaseTearDown int                                    = 0
	TestInstructionColor_SC_TestCaseTearDown              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SC_TestCaseTearDown                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SC_TestCaseTearDown                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                      TypeAndStructs.UpdatedTimeStampType    = "2023-11-28 14:00:00"

	// *** DropZone *** 'TestCaseTearDown_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SC_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "67734e1c-9fdc-463b-904a-950300703bd2"
	TestInstructionDropZoneName_SC_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseTearDown_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SC_TestCaseTearDown_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SC_TestCaseTearDown_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SC_TestCaseTearDown_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "5cfff942-1766-4d78-9973-c57c11609c43" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SC_ExpectedToBePassed
	TestInstructionAttributeName_SC_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SC_ExpectedToBePassed
	TestInstructionAttributeType_SC_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SC_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SC_TestCaseTearDown_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SC_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SC_TestCaseTearDown_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SC_TestCaseTearDown_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SC_TestCaseTearDown_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	FangEngine_Class_Name_SC_TestCaseTearDown = "GeneralSetupTearDown"
)

// TestInstruction_SC_TestCaseTearDown
// Variable that holds the data for the TestInstruction
var TestInstruction_SC_TestCaseTearDown *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SC_TestCaseTearDown
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SC_TestCaseTearDown() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	var fangMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct
	fangMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct)

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_SC_TestCaseTearDown = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDuration: 360,
				},
				FangEngineClassesMethodsAttributes: &FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
					TestInstructionOriginalUUID: TestInstructionUUID_SC_TestCaseTearDown,
					TestInstructionName:         TestInstructionName_SC_TestCaseTearDown,
					FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralSetupTearDown,
					FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralSetupTearDown,
					FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_SC_GeneralSetupTearDown_TearDown,
					FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_SC_GeneralSetupTearDown_TearDown,
					Attributes:                  fangMethodAttributeMap, //make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct),
				},
			},
		},
	}

	// TestInstruction - TestCaseTearDown
	TestInstruction_SC_TestCaseTearDown.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SC,
		DomainName:                   DomainData.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_SC_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SC_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_SC_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_SC_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SC_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_SC_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_SC_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SC_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SC_TestCaseTearDown,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseTearDown
	TestInstruction_SC_TestCaseTearDown.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SC,
		DomainName:                   DomainData.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_SC_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SC_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_SC_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_SC_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SC_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SC_TestCaseTearDown,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SC_TestCaseTearDown,
		TCRuleDeletion:               TCRuleDeletion_SC_TestCaseTearDown,
		TCRuleSwap:                   TCRuleSwap_SC_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_SC_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SC_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_SC_TestCaseTearDown,
	}

	// DropZone 'TestCaseTearDown_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseTearDown_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SC_TestCaseTearDown_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SC_TestCaseTearDown_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SC,
		DomainName:                   DomainData.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_SC_TestCaseTearDown,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SC_TestCaseTearDown_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SC_TestCaseTearDown_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SC_TestCaseTearDown_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SC_TestCaseTearDown_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SC_TestCaseTearDown_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SC_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SC_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SC_TestCaseTearDown_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SC_TestCaseTearDown,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SC_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseTearDown.ImmatureTestInstructionInformation = append(
		TestInstruction_SC_TestCaseTearDown.ImmatureTestInstructionInformation,
		TestInstruction_SC_TestCaseTearDown_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SC_TestCaseTearDown_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SC_TestCaseTearDown_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SC,
		DomainName:                                    DomainData.DomainName_SC,
		TestInstructionUUID:                           TestInstructionUUID_SC_TestCaseTearDown,
		TestInstructionName:                           TestInstructionName_SC_TestCaseTearDown,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SC_TestCaseTearDown_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SC_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseTearDown.TestInstructionAttribute = append(
		TestInstruction_SC_TestCaseTearDown.TestInstructionAttribute,
		TestInstructionAttribute_SC_TestCaseTearDown_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed *FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = &FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_SC_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralAttribute_ExpectedToBePassed,
	}
	fangMethodAttributeMap[TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed
	//TestInstruction_SC_TestCaseTearDown.LocalExecutionMethods.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseTearDown
	var TestInstructionImmatureElementModel_SC_TestCaseTearDown *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SC_TestCaseTearDown = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SC,
		DomainName:               DomainData.DomainName_SC,
		ImmatureElementUUID:      TestInstructionUUID_SC_TestCaseTearDown,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SC_TestCaseTearDown),
		PreviousElementUUID:      TestInstructionUUID_SC_TestCaseTearDown,
		NextElementUUID:          TestInstructionUUID_SC_TestCaseTearDown,
		FirstChildElementUUID:    TestInstructionUUID_SC_TestCaseTearDown,
		ParentElementUUID:        TestInstructionUUID_SC_TestCaseTearDown,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SC_TestCaseTearDown,
		TopImmatureElementUUID:   TestInstructionUUID_SC_TestCaseTearDown,
		IsTopElement:             true,
	}
	TestInstruction_SC_TestCaseTearDown.ImmatureElementModel = append(
		TestInstruction_SC_TestCaseTearDown.ImmatureElementModel,
		TestInstructionImmatureElementModel_SC_TestCaseTearDown)

	return TestInstruction_SC_TestCaseTearDown
}

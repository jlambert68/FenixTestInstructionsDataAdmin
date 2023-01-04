package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseTearDown'
	TestInstructionUUID_SC_TestCaseTearDown               TypeAndStructs.OriginalElementUUIDType = "89959686-1599-4a5f-bcf6-5b427cbdce93"
	TestInstructionName_SC_TestCaseTearDown               TypeAndStructs.TestInstructionNameType = "TestCaseTearDown"
	TestInstructionTypeUUID_SC_TestCaseTearDown                                                  = TestInstructionTypeUUID_SC_GeneralSetUpTearDown
	TestInstructionTypeName_SC_TestCaseTearDown                                                  = TestInstructionTypeName_SC_GeneralSetUpTearDown
	TestInstructionDescription_SC_TestCaseTearDown        string                                 = "TearDown the _SCs execution engine after executing TestInstructions"
	TestInstructionMouseOverText_SC_TestCaseTearDown      string                                 = "TearDown the _SCs execution engine after executing TestInstructions"
	TestInstructionDeprecated_SC_TestCaseTearDown         bool                                   = false
	TestInstructionEnabled_SC_TestCaseTearDown            bool                                   = true
	TestInstructionMajorVersionNumber_SC_TestCaseTearDown int                                    = 0
	TestInstructionMinorVersionNumber_SC_TestCaseTearDown int                                    = 1
	TestInstructionColor_SC_TestCaseTearDown              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SC_TestCaseTearDown                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SC_TestCaseTearDown                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"

	// *** DropZone *** 'TestCaseTearDown_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SC_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "67734e1c-9fdc-463b-904a-950300703bd2"
	TestInstructionDropZoneName_SC_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseTearDown_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SC_TestCaseTearDown_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SC_TestCaseTearDown_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SC_TestCaseTearDown_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "5cfff942-1766-4d78-9973-c57c11609c43" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SC_ExpectedToBePassed
	TestInstructionAttributeName_SC_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_SC_ExpectedToBePassed
	TestInstructionAttributeType_SC_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_SC_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SC_TestCaseTearDown_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SC_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SC_TestCaseTearDown_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SC_TestCaseTearDown_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SC_TestCaseTearDown_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	FangEngine_Class_Name_SC_TestCaseTearDown = "GeneralSetupTearDown"
)

// TestInstruction_SC_TestCaseTearDown
// Variable that holds the data for the TestInstruction
var TestInstruction_SC_TestCaseTearDown TestInstruction_SC_TestCaseSetUpStruct

// Initate_TestInstruction_SC_TestCaseTearDown
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SC_TestCaseTearDown() TestInstruction_SC_TestCaseSetUpStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_SC_TestCaseTearDown = TestInstruction_SC_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_SC_TestCaseTearDown,
			TestInstructionName:         TestInstructionName_SC_TestCaseTearDown,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralSetupTearDown,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralSetupTearDown,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_SC_GeneralSetupTearDown_TearDown,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_SC_GeneralSetupTearDown_TearDown,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - TestCaseTearDown
	TestInstruction_SC_TestCaseTearDown.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_SC,
		DomainName:                   Domains.DomainName_SC,
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
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseTearDown
	TestInstruction_SC_TestCaseTearDown.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_SC,
		DomainName:                   Domains.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_SC_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SC_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_SC_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_SC_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SC_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SC_TestCaseTearDown,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_SC_TestCaseTearDown,
		TCRuleDeletion:               TCRuleDeletion_SC_TestCaseTearDown,
		TCRuleSwap:                   TCRuleSwap_SC_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_SC_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SC_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_SC_TestCaseTearDown,
	}

	// DropZone 'TestCaseTearDown_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseTearDown_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SC_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SC_TestCaseTearDown_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_SC,
		DomainName:                   Domains.DomainName_SC,
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
	var TestInstructionAttribute_SC_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SC_TestCaseTearDown_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_SC,
		DomainName:                                    Domains.DomainName_SC,
		TestInstructionUUID:                           TestInstructionUUID_SC_TestCaseTearDown,
		TestInstructionName:                           TestInstructionName_SC_TestCaseTearDown,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:              TestInstructionAttributeTypeUUID_SC_ExpectedToPass,
		TestInstructionAttributeTypeName:              TestInstructionAttributeTypeName_SC_ExpectedToPass,
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
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SC_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_SC_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseTearDown.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_SC_TestCaseTearDown_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseTearDown
	var TestInstructionImmatureElementModel_SC_TestCaseTearDown TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SC_TestCaseTearDown = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_SC,
		DomainName:               Domains.DomainName_SC,
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

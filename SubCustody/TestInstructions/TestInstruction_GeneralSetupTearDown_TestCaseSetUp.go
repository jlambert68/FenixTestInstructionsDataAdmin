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
	// *** TestInstruction *** 'TestCaseSetUp'
	TestInstructionUUID_SC_TestCaseSetUp               TypeAndStructs.OriginalElementUUIDType = "26d38886-c112-48ef-a20f-4da8fb9a5ccb"
	TestInstructionName_SC_TestCaseSetUp               TypeAndStructs.TestInstructionNameType = "TestCaseSetUp"
	TestInstructionTypeUUID_SC_TestCaseSetUp                                                  = TestInstructionTypeUUID_SC_GeneralSetUpTearDown
	TestInstructionTypeName_SC_TestCaseSetUp                                                  = TestInstructionTypeName_SC_GeneralSetUpTearDown
	TestInstructionDescription_SC_TestCaseSetUp        string                                 = "Initiate _SCs execution engine to be able to execute TestInstructions"
	TestInstructionMouseOverText_SC_TestCaseSetUp      string                                 = "Initiate _SCs execution engine to be able to execute TestInstructions"
	TestInstructionDeprecated_SC_TestCaseSetUp         bool                                   = false
	TestInstructionEnabled_SC_TestCaseSetUp            bool                                   = true
	TestInstructionMajorVersionNumber_SC_TestCaseSetUp int                                    = 0
	TestInstructionMinorVersionNumber_SC_TestCaseSetUp int                                    = 1
	TestInstructionColor_SC_TestCaseSetUp              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SC_TestCaseSetUp                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SC_TestCaseSetUp                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"

	// *** DropZone *** 'TestCaseSetUp_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SC_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "60a2492f-9b0d-4a0e-a374-5845ffabf37d"
	TestInstructionDropZoneName_SC_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseSetUp_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SC_TestCaseSetUp_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SC_TestCaseSetUp_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SC_TestCaseSetUp_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "f4682904-8f60-447c-b851-e713f2b4a03d" //TestInstructionAttributeUUID_SC_ExpectedToBePassed
	TestInstructionAttributeName_SC_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_SC_ExpectedToBePassed
	TestInstructionAttributeType_SC_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_SC_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SC_TestCaseSetUp_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SC_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SC_TestCaseSetUp_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SC_TestCaseSetUp_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SC_TestCaseSetUp_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_SC_TestCaseSetUp
// Variable that holds the data for the TestInstruction
var TestInstruction_SC_TestCaseSetUp TestInstruction_SC_TestCaseSetUpStruct

// Initate_TestInstruction_SC_TestCaseSetUp
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SC_TestCaseSetUp() TestInstruction_SC_TestCaseSetUpStruct {

	// Initiate variable to store all TestInstruction data
	TestInstruction_SC_TestCaseSetUp = TestInstruction_SC_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_SC_TestCaseSetUp,
			TestInstructionName:         TestInstructionName_SC_TestCaseSetUp,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralSetupTearDown,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralSetupTearDown,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_SC_GeneralSetupTearDown_Setup,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_SC_GeneralSetupTearDown_Setup,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - TestCaseSetUp
	TestInstruction_SC_TestCaseSetUp.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_SC,
		DomainName:                   Domains.DomainName_SC,
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
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseSetUp
	TestInstruction_SC_TestCaseSetUp.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_SC,
		DomainName:                   Domains.DomainName_SC,
		TestInstructionUUID:          TestInstructionUUID_SC_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_SC_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SC_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_SC_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_SC_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SC_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SC_TestCaseSetUp,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_SC_TestCaseSetUp,
		TCRuleDeletion:               TCRuleDeletion_SC_TestCaseSetUp,
		TCRuleSwap:                   TCRuleSwap_SC_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_SC_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SC_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_SC_TestCaseSetUp,
	}

	// DropZone 'TestCaseSetUp_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseSetUp_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SC_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SC_TestCaseSetUp_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_SC,
		DomainName:                   Domains.DomainName_SC,
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
	var TestInstructionAttribute_SC_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SC_TestCaseSetUp_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_SC,
		DomainName:                                    Domains.DomainName_SC,
		TestInstructionUUID:                           TestInstructionUUID_SC_TestCaseSetUp,
		TestInstructionName:                           TestInstructionName_SC_TestCaseSetUp,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SC_TestCaseSetUp_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SC_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseSetUp.TestInstructionAttribute = append(
		TestInstruction_SC_TestCaseSetUp.TestInstructionAttribute,
		TestInstructionAttribute_SC_TestCaseSetUp_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SC_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_SC_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_SC_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_SC_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_SC_TestCaseSetUp.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_SC_TestCaseSetUp_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseSetUp
	var TestInstructionImmatureElementModel_SC_TestCaseSetUp TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SC_TestCaseSetUp = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_SC,
		DomainName:               Domains.DomainName_SC,
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

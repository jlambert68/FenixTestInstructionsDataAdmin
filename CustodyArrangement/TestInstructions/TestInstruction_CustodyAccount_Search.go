package TestInstructions

import (
	"FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"FenixTestInstructionsDataAdmin/Domains"
	"FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *************************************
	// *** TestInstruction *** 'Search'
	TestInstructionUUID_CA_CustodyAccount_Search               TypeAndStructs.OriginalElementUUIDType = "f3f7723e-d117-434a-9583-30db87d198ce"
	TestInstructionName_CA_CustodyAccount_Search               TypeAndStructs.TestInstructionNameType = "Search"
	TestInstructionTypeUUID_CA_CustodyAccount_Search                                                  = TestInstructionTypeUUID_CA_CustodyAccount
	TestInstructionTypeName_CA_CustodyAccount_Search                                                  = TestInstructionTypeName_CA_CustodyAccount
	TestInstructionDescription_CA_CustodyAccount_Search        string                                 = "Search for a specific Custody Account"
	TestInstructionMouseOverText_CA_CustodyAccount_Search      string                                 = "Search for a specific Custody Account"
	TestInstructionDeprecated_CA_CustodyAccount_Search         bool                                   = false
	TestInstructionEnabled_CA_CustodyAccount_Search            bool                                   = true
	TestInstructionMajorVersionNumber_CA_CustodyAccount_Search int                                    = 1
	TestInstructionMinorVersionNumber_CA_CustodyAccount_Search int                                    = 0
	TestInstructionColor_CA_CustodyAccount_Search              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_CA_CustodyAccount_Search                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_CustodyAccount_Search                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// *** DropZone *** 'SearchExpectsToSucceed'
	// No Dropzone needed

	// Attribute - 'CustodyAccountId'
	TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId               TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructionAttributeUUID_CA_CustodyAccountId
	TestInstructionAttributeName_CA_CustodyAccount_Search_CustodyAccountId               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_CA_CustodyAccountId
	TestInstructionAttributeType_CA_CustodyAccount_Search_CustodyAccountId               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_CA_CustodyAccountId
	TestInstructionAttributeActionCommand_CA_CustodyAccount_Search_CustodyAccountId      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_CA_CustodyAccount_Search_CustodyAccountId TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_CA_CustodyAccount_Search_CustodyAccountId          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_CA_CustodyAccount_Search_CustodyAccountId        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_CA_CustodyAccount_Search_CustodyAccountId      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_CA_CustodyAccount_SearchStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_CA_CustodyAccount_SearchStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}

// TestInstruction_CA_CustodyAccount_Search
// Variable that holds the data for the TestInstruction
var TestInstruction_CA_CustodyAccount_Search TestInstruction_CA_CustodyAccount_SearchStruct

// Initate_TestInstruction_CA_CustodyAccount_Search
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_CustodyAccount_Search() {

	// Initiate variable to store all TestInstruction data
	TestInstruction_CA_TestCaseSetUp = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			FangEngineClassNameUUID:  FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_CustodyAccount,
			FangEngineClassNameNAME:  FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_CustodyAccount,
			FangEngineMethodNameUUID: FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_CustodyAccount_Search,
			FangEngineMethodNameNAME: FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_CustodyAccount_Search,
			Attributes:               make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - Search
	TestInstruction_CA_CustodyAccount_Search.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:          TestInstructionName_CA_CustodyAccount_Search,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CustodyAccount_Search,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CustodyAccount_Search,
		TestInstructionDescription:   TestInstructionDescription_CA_CustodyAccount_Search,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_CustodyAccount_Search,
		Deprecated:                   TestInstructionDeprecated_CA_CustodyAccount_Search,
		Enabled:                      TestInstructionEnabled_CA_CustodyAccount_Search,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_CustodyAccount_Search,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_CustodyAccount_Search,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - CustodyAccount_Search
	TestInstruction_CA_CustodyAccount_Search.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:          TestInstructionName_CA_CustodyAccount_Search,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CustodyAccount_Search,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CustodyAccount_Search,
		Deprecated:                   TestInstructionDeprecated_CA_CustodyAccount_Search,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_CustodyAccount_Search,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_CustodyAccount_Search,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_CustodyAccount_Search,
		TCRuleDeletion:               TCRuleDeletion_CA_CustodyAccount_Search,
		TCRuleSwap:                   TCRuleSwap_CA_CustodyAccount_Search,
		TestInstructionDescription:   TestInstructionDescription_CA_CustodyAccount_Search,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_CustodyAccount_Search,
		Enabled:                      TestInstructionEnabled_CA_CustodyAccount_Search,
	}

	// DropZone 'SearchExpectsToSucceed'
	// No DropZone needed

	// TestInstruction Attribute - 'CustodyAccountId'
	var TestInstructionAttribute_CA_CustodyAccount_Search_CustodyAccountId TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_CustodyAccount_Search_CustodyAccountId = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:                           TestInstructionName_CA_CustodyAccount_Search,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeTypeUUID:              TestInstructionAttributeTypeUUID_CA_Standard,
		TestInstructionAttributeTypeName:              TestInstructionAttributeTypeName_CA_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_CustodyAccount_Search_CustodyAccountId,
	}
	TestInstruction_CA_CustodyAccount_Search.TestInstructionAttribute = append(
		TestInstruction_CA_CustodyAccount_Search.TestInstructionAttribute,
		TestInstructionAttribute_CA_CustodyAccount_Search_CustodyAccountId)

	// Add FangEngine relation for Attribute - 'CustodyAccountId'
	var tempFangEngineAttributeCustodyAccountId FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeCustodyAccountId = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_CustodyAccount_Search_CustodyAccountId,
		FangEngineAttributeNameUUID:  FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_CustodyAccount_Search_CustodyAccountId,
		FangEngineAttributeNameName:  FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_CustodyAccount_Search_CustodyAccountId,
	}
	TestInstruction_CA_CustodyAccount_Search.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId] = tempFangEngineAttributeCustodyAccountId

	// ImmatureElementModel - Search
	var TestInstructionImmatureElementModel_CA_CustodyAccount_Search TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_CustodyAccount_Search = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_CustodyAccount_Search,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_CustodyAccount_Search),
		PreviousElementUUID:      TestInstructionUUID_CA_CustodyAccount_Search,
		NextElementUUID:          TestInstructionUUID_CA_CustodyAccount_Search,
		FirstChildElementUUID:    TestInstructionUUID_CA_CustodyAccount_Search,
		ParentElementUUID:        TestInstructionUUID_CA_CustodyAccount_Search,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_CustodyAccount_Search,
		TopImmatureElementUUID:   TestInstructionUUID_CA_CustodyAccount_Search,
		IsTopElement:             true,
	}
	TestInstruction_CA_CustodyAccount_Search.ImmatureElementModel = append(
		TestInstruction_CA_CustodyAccount_Search.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_CustodyAccount_Search)
}

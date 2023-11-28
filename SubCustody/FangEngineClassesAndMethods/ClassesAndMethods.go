package FangEngineClassesAndMethods

import "FenixTestInstructionsDataAdmin/TypeAndStructs"

// Types for FangEngine
type FangEngine_ClassName_UUID_SC_Type string
type FangEngine_ClassName_Name_SC_Type string
type FangEngine_MethodName_UUID_SC_Type string
type FangEngine_MethodName_Name_SC_Type string
type FangEngine_AttributeName_UUID_SC_Type string
type FangEngine_AttributeName_Name_SC_Type string

// Type this is used for specifying Classes, Methods and Attributes for FangEngine which is used by _SCs TestAutomation
type FangEngineClassesMethodsAttributesStruct struct {
	TestInstructionOriginalUUID TypeAndStructs.OriginalElementUUIDType
	TestInstructionName         TypeAndStructs.TestInstructionNameType
	FangEngineClassNameUUID     FangEngine_ClassName_UUID_SC_Type
	FangEngineClassNameNAME     FangEngine_ClassName_Name_SC_Type
	FangEngineMethodNameUUID    FangEngine_MethodName_UUID_SC_Type
	FangEngineMethodNameNAME    FangEngine_MethodName_Name_SC_Type
	Attributes                  map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineAttributesStruct
}

type FangEngineAttributesStruct struct {
	TestInstructionAttributeUUID     TypeAndStructs.TestInstructionAttributeUUIDType
	TestInstructionAttributeName     TypeAndStructs.TestInstructionAttributeNameType
	TestInstructionAttributeTypeUUID TypeAndStructs.TestInstructionAttributeTypeUUIDType
	FangEngineAttributeNameUUID      FangEngine_AttributeName_UUID_SC_Type
	FangEngineAttributeNameName      FangEngine_AttributeName_Name_SC_Type
}

// Classes, Methods and their Parameters in FangEngine for _SC
const (

	// General Attribute - ''
	FangEngine_ClassName_UUID_SC_GeneralAttribute_ExpectedToBePassed FangEngine_AttributeName_UUID_SC_Type = "dd6575f9-8a20-4626-9897-286902f7a144"
	FangEngine_ClassName_Name_SC_GeneralAttribute_ExpectedToBePassed FangEngine_AttributeName_Name_SC_Type = "expectedToBePassed"

	// ClassName - ***** 'GeneralSetupTearDown' *****
	FangEngine_ClassName_UUID_SC_GeneralSetupTearDown FangEngine_ClassName_UUID_SC_Type = "6bae70af-e1d1-4f9d-91d7-7d581cb8f278"
	FangEngine_ClassName_Name_SC_GeneralSetupTearDown FangEngine_ClassName_Name_SC_Type = "GeneralSetupTearDown"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'Setup'
	FangEngine_MethodName_UUID_SC_GeneralSetupTearDown_Setup FangEngine_MethodName_UUID_SC_Type = "b0df960b-e488-467d-bff7-a6e324ea4197"
	FangEngine_MethodName_Name_SC_GeneralSetupTearDown_Setup FangEngine_MethodName_Name_SC_Type = "Setup"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'TearDown'
	FangEngine_MethodName_UUID_SC_GeneralSetupTearDown_TearDown FangEngine_MethodName_UUID_SC_Type = "35da6f14-146f-452a-85d0-106781a8e6a3"
	FangEngine_MethodName_Name_SC_GeneralSetupTearDown_TearDown FangEngine_MethodName_Name_SC_Type = "TearDown"
)

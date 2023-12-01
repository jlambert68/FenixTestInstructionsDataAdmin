package SubCustody

import (
	testInstructionContainer_SpecialSerialBaseContainer "FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer"
	testInstructionContainer_SpecialSerialBaseContainer_1_0 "FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer/version_1_0"
	generalSetupTearDown_TestCaseSetUp "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	generalSetupTearDown_TestCaseSetUp_1_0 "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_0"
	generalSetupTearDown_TestCaseSetUp_1_1 "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_1"
	generalSetupTearDown_TestCaseTearDown "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	generalSetupTearDown_TestCaseTearDown_1_0 "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/version_1_0"
	"FenixTestInstructionsDataAdmin/TestInstructionAndTestInstuctionContainerTypes"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
	"FenixTestInstructionsDataAdmin/shared_code"
	"time"
)

var TestInstructionsAndTestInstructionContainers_SC *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct

func GenerateTestInstructions_SC() {

	// Load Allowed users from json-file
	shared_code.ImportAllowedUsersFromFile()

	// Generate TestInstructions
	// GeneralSetupTearDown::TestCaseSetUp
	generalSetupTearDown_TestCaseSetUp_1_1.Initate_TestInstruction_SC_TestCaseSetUp()
	generalSetupTearDown_TestCaseSetUp_1_0.Initate_TestInstruction_SC_TestCaseSetUp()

	// GeneralSetupTearDown::TestCaseTearDown
	generalSetupTearDown_TestCaseTearDown_1_0.Initate_TestInstruction_SC_TestCaseTearDown()

	// Build structure for all TestInstructions & TestInstructionContainers to be sent over gRPC to Fenix Backend
	TestInstructionsAndTestInstructionContainers_SC = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{

		// TestInstructions
		TestInstructions: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{

				//TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				generalSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						//Version 'generalSetupTearDown_TestCaseSetUp_1.1'
						&TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  "HASH",
						},

						// Version 'generalSetupTearDown_TestCaseSetUp_1.0'
						&TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  "HASH",
						},
					},
					TestInstructionVersionsHash: "HASH",
				},

				// TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				generalSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'generalSetupTearDown_TestCaseSetUp_1.0'
						&TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  "HASH",
						},
					},
					TestInstructionVersionsHash: "HASH",
				},
			},
			TestInstructionsHash: "HASH",
		},

		// TestInstructionContainers
		TestInstructionContainers: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{},
		AllowedUsers:              shared_code.AllowedUsersLoadFronJsonFile.AllowedUsers,
		TestInstructionsAndTestInstructionsContainersMessageHash:        "HASH",
		MessageCreationTimeStamp:                                        time.Now(),
		ForceNewBaseLineForTestInstructionsAndTestInstructionContainers: false,
	}

	// Generate TestInstructionContainers
	// testInstructionContainer_SpecialSerialBaseContainer
	testInstructionContainer_SpecialSerialBaseContainer_1_0.Initiate_TestInstructionContainer_SC_Serial(TestInstructionsAndTestInstructionContainers_SC)

	// TestInstructionContainers
	var testInstructionContainers *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct
	testInstructionContainers = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{
		TestInstructionContainersMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{

			// 'SC_SpecialSerialBaseContainer'
			testInstructionContainer_SpecialSerialBaseContainer.TestInstructionContainerUUID_SC_SpecialSerialBaseContainer: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{
				TestInstructionContainerVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{

					//Version 'TestInstructionContainer_SpecialSerialBaseContainer_1.0'
					&TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{
						TestInstructionContainerInstance:             testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SC_SpecialSerialBase,
						TestInstructionContainerInstanceMajorVersion: testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SC_SpecialSerialBase.TestInstructionContainer.MajorVersionNumber,
						TestInstructionContainerInstanceMinorVersion: testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SC_SpecialSerialBase.TestInstructionContainer.MinorVersionNumber,
						Deprecated:                           testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SC_SpecialSerialBase.TestInstructionContainer.Deprecated,
						Enabled:                              testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SC_SpecialSerialBase.TestInstructionContainer.Enabled,
						TestInstructionContainerInstanceHash: "HASH",
					},
				},
				TestInstructionContainerVersionsHash: "HASH",
			},
		},

		TestInstructionContainersHash: "HASH",
	}
	TestInstructionsAndTestInstructionContainers_SC.TestInstructionContainers = testInstructionContainers

	// TODO Calculate alla Hashes for TestInstructions-block and TestInstructionContainersMap-block
	shared_code.CalculateTestInstructionAndTestInstructionContainerMessageHashes(TestInstructionsAndTestInstructionContainers_SC)

}

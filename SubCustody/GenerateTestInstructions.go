package SubCustody

import (
	testInstructionContainer_SpecialSerialBaseContainer "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer"
	testInstructionContainer_SpecialSerialBaseContainer_1_0 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer/"
	generalSetupTearDown_TestCaseSetUp "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	generalSetupTearDown_TestCaseSetUp_1_0 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_0"
	generalSetupTearDown_TestCaseSetUp_1_1 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_1"
	generalSetupTearDown_TestCaseTearDown "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	generalSetupTearDown_TestCaseTearDown_1_0 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"time"
)

var TestInstructions_SC TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct

func generateTestInstructions_SC() {

	// GeneralSetupTearDown::TestCaseSetUp
	generalSetupTearDown_TestCaseSetUp_1_1.Initate_TestInstruction_SC_TestCaseSetUp()
	generalSetupTearDown_TestCaseSetUp_1_0.Initate_TestInstruction_SC_TestCaseSetUp()

	// GeneralSetupTearDown::TestCaseTearDown
	generalSetupTearDown_TestCaseTearDown_1_0.Initate_TestInstruction_SC_TestCaseTearDown()

	// Generate all TestInstructions & TestInstructionContainers
	TestInstructions_SC = TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{

		// TestInstructions
		TestInstructions: TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap: map[TypeAndStructs.OriginalElementUUIDType]TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{

				//TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				generalSetupTearDown_TestCaseSetUp.TestInstructionUUID_SC_TestCaseSetUp: TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						//Version 'generalSetupTearDown_TestCaseSetUp_1.1'
						TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceHash:         "HASH",
						},

						// Version 'generalSetupTearDown_TestCaseSetUp_1.0'
						TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
							TestInstructionInstanceHash:         "HASH",
						},
					},
					TestInstructionVersionsHash: "HASH",
				},

				// TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				generalSetupTearDown_TestCaseTearDown.TestInstructionUUID_SC_TestCaseTearDown: TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'generalSetupTearDown_TestCaseSetUp_1.0'
						TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.MajorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SC_TestCaseTearDown.TestInstruction.Enabled,
							TestInstructionInstanceHash:         "HASH",
						},
					},
					TestInstructionVersionsHash: "HASH",
				},
			},
			TestInstructionsHash: "HASH",
		},

		// TestInstructionContainers
		TestInstructionContainers: TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{
			TestInstructionContainersMap: map[TypeAndStructs.OriginalElementUUIDType]TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{

				// 'SC_SpecialSerialBaseContainer'
				testInstructionContainer_SpecialSerialBaseContainer.TestInstructionContainerUUID_SC_SpecialSerialBaseContainer: TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{
				TestInstructionContainerVersions:     []TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{

					//Version 'TestInstructionContainer_SpecialSerialBaseContainer_1.0'
					TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{
						TestInstructionContainerInstance:             TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct{},
						TestInstructionContainerInstanceMajorVersion: 0,
						TestInstructionContainerInstanceMinorVersion: 0,
						Deprecated:                           false,
						Enabled:                              false,
						TestInstructionContainerInstanceHash: "",
					},
				},
				TestInstructionContainerVersionsHash: "",
			},
			},

				MessageCreationTimeStamp:                                 time.Now(),
				TestInstructionsAndTestInstructionsContainersMessageHash: "HASH",
			},
		},
	}

	// TODO Calculate alla Hashes for TestInstructions-block

	// Generate all TestInstructionContainersMap
	var TestInstructionContainers map[TypeAndStructs.OriginalElementUUIDType]TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct
	TestInstructionContainers = map[TypeAndStructs.OriginalElementUUIDType]TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{}

	var xx TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct
	xx = TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{
		TestInstructionContainerVersions:     nil,
		TestInstructionContainerVersionsHash: "",
	}
	// TODO Calculate alla Hashes for TestInstructionContainersMap-block
}

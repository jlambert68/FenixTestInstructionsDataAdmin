package SubCustody

import (
	"FenixTestInstructionsDataAdmin/Domains"
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
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"os"
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
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},

						// Version 'generalSetupTearDown_TestCaseSetUp_1.0'
						&TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
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
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},
			},
			TestInstructionsHash: shared_code.InitialValueBeforeHashed,
		},

		// TestInstructionContainers
		TestInstructionContainers: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{},
		AllowedUsers:              shared_code.AllowedUsersLoadFronJsonFile,
		TestInstructionsAndTestInstructionsContainersAndUsersMessageHash: shared_code.InitialValueBeforeHashed,
		MessageCreationTimeStamp: time.Now(),
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
						TestInstructionContainerInstanceHash: shared_code.InitialValueBeforeHashed,
					},
				},
				TestInstructionContainerVersionsHash: shared_code.InitialValueBeforeHashed,
			},
		},

		TestInstructionContainersHash: shared_code.InitialValueBeforeHashed,
	}
	TestInstructionsAndTestInstructionContainers_SC.TestInstructionContainers = testInstructionContainers

	// Calculate hashes that is included in the Supported TestInstructions, TestInstructionContainers and Allowed Users message
	var err error
	err = shared_code.CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(TestInstructionsAndTestInstructionContainers_SC)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Convert supported TestInstructions, TestInstructionContainers and Allowed Users message into a gRPC-version of the message
	var supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage *fenixExecutionWorkerGrpcApi.SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage, err = shared_code.
		GenerateTestInstructionAndTestInstructionContainerAndUserGrpcMessage(string(Domains.DomainUUID_SC), TestInstructionsAndTestInstructionContainers_SC)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Convert back supported TestInstructions, TestInstructionContainers and Allowed Users message from a gRPC-version of the message and check correctness of Hashes
	var testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct
	testInstructionsAndTestInstructionContainersMessage, err = shared_code.
		GenerateStandardFromGrpcMessageForTestInstructionsAndUsers(supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Verify recreated Hashes from gRPC-message
	var errorSlice []error
	errorSlice = shared_code.VerifyTestInstructionAndTestInstructionContainerAndUsersMessageHashes(testInstructionsAndTestInstructionContainersMessage)
	if errorSlice != nil {
		for _, errFromSlice := range errorSlice {
			fmt.Println(errFromSlice)
		}
		os.Exit(1)
	}

	// Send supported TestInstructions, TestInstructionContainers and Allowed Users message to Worker
	err = shared_code.SendGrpcTestInstructionsAndTestInstructionContainersAndAllowedUsersMessageToWorker(supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

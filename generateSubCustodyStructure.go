package main

import (
	"fmt"
	generalSetupTearDown_TestCaseSetUp_0_1 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/0_1"
	generalSetupTearDown_TestCaseSetUp_0_2 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/0_2"
	generalSetupTearDown_TestCaseTearDown_0_1 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/0_1"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestInstructionAndTestInstuctionContainerTypes"
	"time"
)

func generateTestInstructionsAndTestInstructionsContainers_SC() {

	// GeneralSetupTearDown::TestCaseSetUp
	generalSetupTearDown_TestCaseSetUp_0_1.Initate_TestInstruction_SC_TestCaseSetUp()
	generalSetupTearDown_TestCaseSetUp_0_2.Initate_TestInstruction_SC_TestCaseSetUp()

	/*

		var TestInstructionInstanceVersion_0_1 TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct
		TestInstructionInstanceVersion_0_1 = TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
			TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp,
			TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
			TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
			Deprecated    :               generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
			Enabled        :              generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
			TestInstructionInstanceHash:         "HASH",
		}
		var TestInstructionInstanceVersion_0_2 TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct
		TestInstructionInstanceVersion_0_2 = TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
			TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp,
			TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
			TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
			TestInstructionInstanceHash:         "HASH",
		}

		var TestInstructionInstanceVersions []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct
		TestInstructionInstanceVersions = []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

			// generalSetupTearDown_TestCaseSetUp_0_1
			TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
				TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp,
				TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
				TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
				Deprecated    :               generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
				Enabled        :              generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
				TestInstructionInstanceHash:         "HASH",
			},

			//generalSetupTearDown_TestCaseSetUp_0_2
			TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
				TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp,
				TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
				TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
				TestInstructionInstanceHash:         "HASH",
			},
		}

		var TestInstruction_no1 TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct
		TestInstruction_no1 = TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
			TestInstructionVersions:      []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

				// generalSetupTearDown_TestCaseSetUp_0_1
				TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
					TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp,
					TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
					TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
					Deprecated    :               generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
					Enabled        :              generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
					TestInstructionInstanceHash:         "HASH",
				},

				//generalSetupTearDown_TestCaseSetUp_0_2
				TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
					TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp,
					TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
					TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
					TestInstructionInstanceHash:         "HASH",
				},
			},
			TestInstructionVersionsHash: "HASH",
		}

		var TestInstruction_no2 TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct
		TestInstruction_no2 = TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
			TestInstructionVersions:      []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

				// generalSetupTearDown_TestCaseSetUp_0_1
				TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
					TestInstructionInstance:             generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1,
					TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.MajorVersionNumber,
					TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.MajorVersionNumber,
					Deprecated    :               generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.Deprecated,
					Enabled        :              generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.Enabled,
					TestInstructionInstanceHash:         "HASH",
				},
			},
			TestInstructionVersionsHash: "HASH",
		}



		var TestInstructions TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct
		TestInstructions = TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructions:     []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
				TestInstruction_no1,
				TestInstruction_no2,
			},
			TestInstructionsHash: "HASH",
		}


	*/

	var testInstructionsAndTestInstructionsContainers TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct
	testInstructionsAndTestInstructionsContainers = TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{
		TestInstructions: TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructions: []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{

				//TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'generalSetupTearDown_TestCaseSetUp_0_1'
						TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseSetUp_0_1.TestInstruction_SC_TestCaseSetUp.TestInstruction.Enabled,
							TestInstructionInstanceHash:         "HASH",
						},

						//Version 'generalSetupTearDown_TestCaseSetUp_0_2'
						TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_0_2.TestInstruction_SC_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceHash:         "HASH",
						},
					},
					TestInstructionVersionsHash: "HASH",
				},

				// TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'generalSetupTearDown_TestCaseSetUp_0_1'
						TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
							TestInstructionInstance:             generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.MajorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseTearDown_0_1.TestInstruction_SC_TestCaseTearDown_0_1.TestInstruction.Enabled,
							TestInstructionInstanceHash:         "HASH",
						},
					},
					TestInstructionVersionsHash: "HASH",
				},
			},
			TestInstructionsHash: "HASH",
		},
		TestInstructionContainers:                                TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{},
		MessageCreationTimeStamp:                                 time.Time{},
		TestInstructionsAndTestInstructionsContainersMessageHash: "HASH",
	}

	fmt.Println(testInstructionsAndTestInstructionsContainers)

}

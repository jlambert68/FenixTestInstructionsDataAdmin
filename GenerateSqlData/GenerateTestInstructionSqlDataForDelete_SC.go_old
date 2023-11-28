package GenerateSqlData

import (
	"FenixTestInstructionsDataAdmin/BaseSQL"
	"FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_0"
	__12 "FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/version_1_0"
	"FenixTestInstructionsDataAdmin/shared_code"
	"fmt"
)

// Function that creates all data for the TestInstruction
func GenerateSqlDelete_For_TestInstructions_SC() {

	testInstructionsSQL := BaseSQL.TestInstructionsSQLDelete
	basicTestInstructionInformationSQL := BaseSQL.BasicTestInstructionInformationSQLDelete
	immatureTestInstructionInformationSQL := BaseSQL.ImmatureTestInstructionInformationSQLDelete
	testInstructionAttributesSQL := BaseSQL.TestInstructionAttributesSQLDelete
	immatureElementModelMessageSQL := BaseSQL.ImmatureElementModelMessageSQLDelete

	var testInstructionUuidSlice []string

	// GeneralSetUpTearDown
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(version_1_0.TestInstructionUUID_SC_TestCaseSetUp))
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(__12.TestInstructionUUID_SC_TestCaseTearDown))

	testInstructionsSQLToBeAdded := shared_code.GenerateSQLINArray(testInstructionUuidSlice)
	testInstructionsSQLToBeAdded = testInstructionsSQLToBeAdded + ";"

	testInstructionsSQL = testInstructionsSQL + testInstructionsSQLToBeAdded
	basicTestInstructionInformationSQL = basicTestInstructionInformationSQL + testInstructionsSQLToBeAdded
	immatureTestInstructionInformationSQL = immatureTestInstructionInformationSQL + testInstructionsSQLToBeAdded
	testInstructionAttributesSQL = testInstructionAttributesSQL + testInstructionsSQLToBeAdded
	immatureElementModelMessageSQL = immatureElementModelMessageSQL + testInstructionsSQLToBeAdded

	// Output all SQL-data
	fmt.Println(immatureElementModelMessageSQL)
	fmt.Println(testInstructionAttributesSQL)
	fmt.Println(immatureTestInstructionInformationSQL)
	fmt.Println(basicTestInstructionInformationSQL)
	fmt.Println(testInstructionsSQL)

}

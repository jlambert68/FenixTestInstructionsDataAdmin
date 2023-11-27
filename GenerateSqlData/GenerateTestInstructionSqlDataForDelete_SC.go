package GenerateSqlData

import (
	"fmt"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/BaseSQL"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/0_1"
	__12 "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/0_1"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
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
		string(__1.TestInstructionUUID_SC_TestCaseSetUp))
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

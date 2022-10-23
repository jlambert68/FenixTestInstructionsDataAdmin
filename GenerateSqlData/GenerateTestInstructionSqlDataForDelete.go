package GenerateSqlData

import (
	"fmt"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/BaseSQL"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

// Initate_TestInstruction_CA_CustodyAccount_Search
// Function that creates all data for the TestInstruction
func GenerateSqlDelete_For_TestInstructions(testInstructionSetUp *TestInstructions.TestInstruction_CA_TestCaseSetUpStruct) {

	testInstructionsSQL := BaseSQL.TestInstructionsSQLDelete
	basicTestInstructionInformationSQL := BaseSQL.BasicTestInstructionInformationSQLDelete
	immatureTestInstructionInformationSQL := BaseSQL.ImmatureTestInstructionInformationSQLDelete
	testInstructionAttributesSQL := BaseSQL.TestInstructionAttributesSQLDelete
	immatureElementModelMessageSQL := BaseSQL.ImmatureElementModelMessageSQLDelete

	var testInstructuinUuidSlice []string

	// Add all TestInstructions for Custody Arrangement
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_CustodyAccount_Search))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_AddOrDeleteSelectedSwift))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_TestCaseSetUp))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_TestCaseTearDown))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift))
	testInstructuinUuidSlice = append(testInstructuinUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_Edit))

	testInstructionsSQLToBeAdded := shared_code.GenerateSQLINArray(testInstructuinUuidSlice)
	testInstructionsSQLToBeAdded = testInstructionsSQLToBeAdded + ";"

	testInstructionsSQL = testInstructionsSQL + testInstructionsSQLToBeAdded
	basicTestInstructionInformationSQL = basicTestInstructionInformationSQL + testInstructionsSQLToBeAdded
	immatureTestInstructionInformationSQL = immatureTestInstructionInformationSQL + testInstructionsSQLToBeAdded
	testInstructionAttributesSQL = testInstructionAttributesSQL + testInstructionsSQLToBeAdded
	immatureElementModelMessageSQL = immatureElementModelMessageSQL + testInstructionsSQLToBeAdded

	// Output all SQL-data
	fmt.Println(testInstructionsSQL)
	fmt.Println(basicTestInstructionInformationSQL)
	fmt.Println(immatureTestInstructionInformationSQL)
	fmt.Println(testInstructionAttributesSQL)
	fmt.Println(immatureElementModelMessageSQL)

}

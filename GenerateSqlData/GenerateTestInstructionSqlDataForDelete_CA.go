package GenerateSqlData

import (
	"fmt"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/BaseSQL"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

// Initate_TestInstruction_CA_CustodyAccount_Search
// Function that creates all data for the TestInstruction
func GenerateSqlDelete_For_TestInstructions_CA() {

	testInstructionsSQL := BaseSQL.TestInstructionsSQLDelete
	basicTestInstructionInformationSQL := BaseSQL.BasicTestInstructionInformationSQLDelete
	immatureTestInstructionInformationSQL := BaseSQL.ImmatureTestInstructionInformationSQLDelete
	testInstructionAttributesSQL := BaseSQL.TestInstructionAttributesSQLDelete
	immatureElementModelMessageSQL := BaseSQL.ImmatureElementModelMessageSQLDelete

	var testInstructionUuidSlice []string

	// Add all TestInstructions for Custody Arrangement
	// CustodyAccount
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_CustodyAccount_Search))

	// FundExecutionAgreementManagement
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_AddOrDeleteSelectedSwift))

	// GeneralSetUpTearDown
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_TestCaseSetUp))
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_TestCaseTearDown))

	// SettlementAgreement
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_AddSelectedInstructedParties))
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_AddSelectedSwift))
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties))
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift))
	testInstructionUuidSlice = append(testInstructionUuidSlice,
		string(TestInstructions.TestInstructionUUID_CA_SettlementAgreement_Edit))

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

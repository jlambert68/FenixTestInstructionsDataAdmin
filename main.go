package main

import (
	"FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"FenixTestInstructionsDataAdmin/GenerateSqlData"
)

func main() {

	// CustodyAccount::Search
	//TestInstructions.Initate_TestInstruction_CA_CustodyAccount_Search()
	//GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&TestInstructions.TestInstruction_CA_CustodyAccount_Search)

	//TestInstructions.Initate_TestInstruction_CA_TestCaseSetUp()
	//GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&TestInstructions.TestInstruction_CA_TestCaseSetUp)

	//TestInstructions.Initate_TestInstruction_CA_TestCaseTearDown()
	//GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&TestInstructions.TestInstruction_CA_TestCaseTearDown)

	//TestInstructions.Initate_TestInstruction_CA_SettlementAgreement_Edit()
	//GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&TestInstructions.TestInstruction_CA_SettlementAgreement_Edit)

	TestInstructions.Initate_TestInstruction_CA_SettlementAgreement_AddSelectedSwift()
	GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&TestInstructions.TestInstruction_CA_SettlementAgreement_AddSelectedSwift)

}

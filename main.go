package main

import (
	"fmt"
	CustodyArrangementTestInstructions "github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/GenerateSqlData"
)

func main() {

	doInsert := true

	if doInsert == true {

		// CustodyAccount::Search
		CustodyArrangementTestInstructions.Initate_TestInstruction_CA_CustodyAccount_Search()
		GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&CustodyArrangementTestInstructions.TestInstruction_CA_CustodyAccount_Search)

		// GeneralSetupTearDown::TestCaseSetUp
		CustodyArrangementTestInstructions.Initate_TestInstruction_CA_TestCaseSetUp()
		GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&CustodyArrangementTestInstructions.TestInstruction_CA_TestCaseSetUp)

		// GeneralSetupTearDown::TestCaseTearDown
		CustodyArrangementTestInstructions.Initate_TestInstruction_CA_TestCaseTearDown()
		GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&CustodyArrangementTestInstructions.TestInstruction_CA_TestCaseTearDown)

		// SettlementAgreement::Edit
		CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_Edit()
		GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_Edit)

		// SettlementAgreement::AddSelectedSwift
		CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_AddSelectedSwift()
		GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_AddSelectedSwift)

		// SettlementAgreement::DeleteSelectedSwift
		CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift()
		GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift)

		// SettlementAgreement::AddSelectedInstructedParties
		CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties()
		GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties)
	} else {
		// SettlementAgreement::DeleteSelectedInstructedParties
		GenerateSqlData.GenerateSqlDelete_For_TestInstructions()
	}

	// Delete Data In Execution Server
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestCaseExecutionQueue\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestCasesUnderExecution\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestInstructionAttributesUnderExecution\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestInstructionExecutionQueue\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestInstructionsUnderExecution\";")

}

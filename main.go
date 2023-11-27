package main

import (
	"fmt"
	CustodyArrangementTestInstructions "github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/GenerateSqlData"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/0_1"
	SubCustodyTestInstructions "github.com/jlambert68/FenixTestInstructionsDataAdmin/SubCustody/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/0_1"
	"log"
)

func main() {

	// Which system to process data for
	system := "Sub Custody"
	//System := "Custody Arrangement"

	switch system {
	case "Custody Arrangement":

		// Should data be Inserted into DB or Deleted from DB
		doInsert := true

		if doInsert == true {

			// CustodyAccount::Search
			CustodyArrangementTestInstructions.Initate_TestInstruction_CA_CustodyAccount_Search()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_CA(&CustodyArrangementTestInstructions.TestInstruction_CA_CustodyAccount_Search)

			// GeneralSetupTearDown::TestCaseSetUp
			CustodyArrangementTestInstructions.Initate_TestInstruction_CA_TestCaseSetUp()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_CA(&CustodyArrangementTestInstructions.TestInstruction_CA_TestCaseSetUp)

			// GeneralSetupTearDown::TestCaseTearDown
			CustodyArrangementTestInstructions.Initate_TestInstruction_CA_TestCaseTearDown()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_CA(&CustodyArrangementTestInstructions.TestInstruction_CA_TestCaseTearDown)

			// SettlementAgreement::Edit
			CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_Edit()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_CA(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_Edit)

			// SettlementAgreement::AddSelectedSwift
			CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_AddSelectedSwift()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_CA(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_AddSelectedSwift)

			// SettlementAgreement::DeleteSelectedSwift
			CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_CA(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift)

			// SettlementAgreement::AddSelectedInstructedParties
			CustodyArrangementTestInstructions.Initate_TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_CA(&CustodyArrangementTestInstructions.TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties)
		} else {
			// Delete all TestInstruction-data
			GenerateSqlData.GenerateSqlDelete_For_TestInstructions_CA()
		}

	case "Sub Custody":

		// Should data be Inserted into DB or Deleted from DB
		doInsert := true

		if doInsert == true {

			// GeneralSetupTearDown::TestCaseSetUp
			__1.Initate_TestInstruction_SC_TestCaseSetUp()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_SC(&__1.TestInstruction_SC_TestCaseSetUp)

			// GeneralSetupTearDown::TestCaseTearDown
			SubCustodyTestInstructions.Initate_TestInstruction_SC_TestCaseTearDown0_1()
			GenerateSqlData.GenerateSqlInsert_For_TestInstructions_SC(&SubCustodyTestInstructions.TestInstruction_SC_TestCaseTearDown_0_1)

		} else {
			// Delete all TestInstruction-data
			GenerateSqlData.GenerateSqlDelete_For_TestInstructions_SC()
		}
	default:
		log.Fatalln("Unknown system: " + system)
	}

	// Delete Data In Execution Server
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestCaseExecutionQueue\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestCasesUnderExecution\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestInstructionAttributesUnderExecution\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestInstructionExecutionQueue\";")
	fmt.Println("DELETE FROM \"FenixExecution\".\"TestInstructionsUnderExecution\";")

}

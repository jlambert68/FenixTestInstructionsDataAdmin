package main

import (
	"FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"FenixTestInstructionsDataAdmin/GenerateSqlData"
)

func main() {

	TestInstructions.Initate_TestInstruction_CA_CustodyAccount_Search()
	GenerateSqlData.GenerateSqlInsert_For_TestInstructions(&TestInstructions.TestInstruction_CA_CustodyAccount_Search)

}

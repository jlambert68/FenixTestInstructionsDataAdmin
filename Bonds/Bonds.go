package Bonds

import "FenixTestInstructionsDataAdmin/TypeAndStructs"

const (
// B0_BOND

Bond_B0_BondUuid                              TypeAndStructs.BondUUIDType = "14dc7af0-5e46-4073-b920-8dffee7ca307"
Bond_B0_BondName                              TypeAndStructs.BondNameType = "B0_BOND"
Bond_B0_BondDescription                       TypeAndStructs.BondDescriptionType =  "Bond with zero connection. This is the base for any TestCase: \"B0\"
Bond_B0_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType = "Bond with zero connection. This is the base for any TestCase: \"B0\""
Bond_B0_Deprecated                            TypeAndStructs.BondDeprecatedType = false
Bond_B0_Enabled                               TypeAndStructs.BondEnabledType = true
Bond_B0_Visible                               TypeAndStructs.BondVisibleType = true
Bond_B0_BondColor                             TypeAndStructs.BondColorType = "#ffff00"
Bond_B0_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType =  false
Bond_B0_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType =  false
Bond_B0_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType =   true
Bond_B0_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType = "TCRuleDeletion001"
Bond_B0_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType = "TCRuleSwap001"
Bond_B0_TestCaseModelElementType              TypeAndStructs.BondTestCaseModelElementTypeType ="B0_BOND"
Bond_B0_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.BondTestCaseModelElementTypeGrpcMappingIDType = 0 //

)


[
{
"BondUuid":
"BondName": ,
"BondDescription":
"BondMouseOverText": ,
"Deprecated": ,
"Enabled": ,
"Visible": ,
"BondColor":
"CanBeDeleted": ,
"CanBeSwappedOut": ,
"ShowBondAttributes": ,
"TCRuleDeletion": ,
"TCRuleSwap": ,
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": ,
"TestCaseModelElementTypeGrpcMappingId": 0
},
{
"BondUuid": "2faab7f8-9f08-44ad-bc02-15b205146f48",
"BondName": "B1f_BOND_NONE_SWAPPABLE",
"BondDescription": "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers",
"BondMouseOverText": "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion002",
"TCRuleSwap": "TCRuleSwap002",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B1f_BOND_NONE_SWAPPABLE",
"TestCaseModelElementTypeGrpcMappingId": 1
},
{
"BondUuid": "8a978853-f067-4978-aea0-07cceae8a1e1",
"BondName": "B10ox_BOND",
"BondDescription": "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x) into TIC(B11fx-n-B11l).",
"BondMouseOverText": "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x) into TIC(B11fx-n-B11l).",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion008",
"TCRuleSwap": "TCRuleSwap008",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B10ox_BOND",
"TestCaseModelElementTypeGrpcMappingId": 8
},
{
"BondUuid": "f3277f00-c2e7-496d-948b-28ce1534785f",
"BondName": "B10xo_BOND",
"BondDescription": "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10x*) into TIC(B11f-n-B11lx).",
"BondMouseOverText": "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10x*) into TIC(B11f-n-B11lx).",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion009",
"TCRuleSwap": "TCRuleSwap009",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B10xo_BOND",
"TestCaseModelElementTypeGrpcMappingId": 9
},
{
"BondUuid": "a2361365-1cc6-449b-bcf4-525608adcf8f",
"BondName": "B11fx_BOND_NONE_SWAPPABLE",
"BondDescription": " Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers before(first) structure.",
"BondMouseOverText": " Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers before(first) structure.",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion010",
"TCRuleSwap": "TCRuleSwap010",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B11fx_BOND_NONE_SWAPPABLE",
"TestCaseModelElementTypeGrpcMappingId": 10
},
{
"BondUuid": "efe03475-91a2-4525-bac1-89fcfb246491",
"BondName": "B11f_BOND",
"BondDescription": "Bond with one connection which appear as first element inside a TestInstructionContainer",
"BondMouseOverText": "Bond with one connection which appear as first element inside a TestInstructionContainer",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion004",
"TCRuleSwap": "TCRuleSwap004",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B11f_BOND",
"TestCaseModelElementTypeGrpcMappingId": 4
},
{
"BondUuid": "78f4c8ef-b8f3-40a2-a335-b28d2ba02e76",
"BondName": "B12_BOND",
"BondDescription": "Bond with two connections which connects TestInstructions or TestInstructionContainers",
"BondMouseOverText": "Bond with two connections which connects TestInstructions or TestInstructionContainers",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion006",
"TCRuleSwap": "TCRuleSwap006",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B12_BOND",
"TestCaseModelElementTypeGrpcMappingId": 6
},
{
"BondUuid": "b1aaae68-076d-4527-a167-228958c9ab01",
"BondName": "B11lx_BOND_NONE_SWAPPABLE",
"BondDescription": "Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers after(last) structure.",
"BondMouseOverText": "Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers after(last) structure.",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion011",
"TCRuleSwap": "TCRuleSwap011",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B11lx_BOND_NONE_SWAPPABLE",
"TestCaseModelElementTypeGrpcMappingId": 11
},
{
"BondUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
"BondName": "B10_BOND",
"BondDescription": "Bond with zero connections which appear as an element in a new TestInstructionContainer without any other elements",
"BondMouseOverText": "Bond with zero connections which appear as an element in a new TestInstructionContainer without any other elements",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion003",
"TCRuleSwap": "TCRuleSwap003",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B10_BOND",
"TestCaseModelElementTypeGrpcMappingId": 3
},
{
"BondUuid": "2c45d283-4625-40e3-bf8e-cb070f1a6210",
"BondName": "B11l_BOND",
"BondDescription": "Bond with one connection which appear as last element inside a TestInstructionContainer",
"BondMouseOverText": "Bond with one connection which appear as last element inside a TestInstructionContainer",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion005",
"TCRuleSwap": "TCRuleSwap005",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B11l_BOND",
"TestCaseModelElementTypeGrpcMappingId": 5
},
{
"BondUuid": "f1cc879d-3b95-4dea-84e5-e2a166c5fc95",
"BondName": "B10oxo_BOND",
"BondDescription": "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x*) into TIC(B11fx-n-B11lx).",
"BondMouseOverText": "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x*) into TIC(B11fx-n-B11lx).",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion007",
"TCRuleSwap": "TCRuleSwap007",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B10oxo_BOND",
"TestCaseModelElementTypeGrpcMappingId": 7
},
{
"BondUuid": "123ab7f8-9f08-44ad-bc02-15b205146f48",
"BondName": "B1l_BOND_NONE_SWAPPABLE",
"BondDescription": "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers",
"BondMouseOverText": "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion002",
"TCRuleSwap": "TCRuleSwap002",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B1l_BOND_NONE_SWAPPABLE",
"TestCaseModelElementTypeGrpcMappingId": 2
},
{
"BondUuid": "9e613527-daa6-4559-9c73-adb8213eccd9",
"BondName": "B12x_BOND_NONE_SWAPPABLE",
"BondDescription": "Bond with two connections which connects TestInstructions or TestInstructiB12x_BOND_NONE_SWAPPABLEonContainers. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers within a pre-created container.",
"BondMouseOverText": "Bond with two connections which connects TestInstructions or TestInstructiB12x_BOND_NONE_SWAPPABLEonContainers. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers within a pre-created container.",
"Deprecated": false,
"Enabled": true,
"Visible": true,
"BondColor": "#ffff00",
"CanBeDeleted": false,
"CanBeSwappedOut": false,
"ShowBondAttributes": true,
"TCRuleDeletion": "TCRuleDeletion012",
"TCRuleSwap": "TCRuleSwap012",
"UpdatedTimeStamp": "2022-06-19 13:57:56.000000",
"TestCaseModelElementType": "B12x_BOND_NONE_SWAPPABLE",
"TestCaseModelElementTypeGrpcMappingId": 12
}
]
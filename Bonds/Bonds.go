package Bonds

import (
	"FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"FenixTestInstructionsDataAdmin/TypeAndStructs"
)

const (
	// B0_BOND
	Bond_B0_BondUuid                              TypeAndStructs.BondUUIDType                          = "14dc7af0-5e46-4073-b920-8dffee7ca307"
	Bond_B0_BondName                              TypeAndStructs.BondNameType                          = "B0_BOND"
	Bond_B0_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with zero connection. This is the base for any TestCase: 'B0'"
	Bond_B0_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with zero connection. This is the base for any TestCase: 'B0'"
	Bond_B0_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B0_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B0_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B0_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B0_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B0_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B0_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B0_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion001"
	Bond_B0_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap001"
	Bond_B0_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B0
	Bond_B0_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B0

	// B1f_BOND
	Bond_B1f_BondUuid                              TypeAndStructs.BondUUIDType                          = "2faab7f8-9f08-44ad-bc02-15b205146f48"
	Bond_B1f_BondName                              TypeAndStructs.BondNameType                          = "B1f_BOND_NONE_SWAPPABLE"
	Bond_B1f_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers"
	Bond_B1f_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers"
	Bond_B1f_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B1f_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B1f_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B1f_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B1f_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B1f_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B1f_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B1f_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion002"
	Bond_B1f_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap002"
	Bond_B1f_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B1f
	Bond_B1f_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B1f

	// B10ox_BOND
	Bond_B10ox_BondUuid                              TypeAndStructs.BondUUIDType                          = "8a978853-f067-4978-aea0-07cceae8a1e1"
	Bond_B10ox_BondName                              TypeAndStructs.BondNameType                          = "B10ox_BOND"
	Bond_B10ox_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x) into TIC(B11fx-n-B11l)."
	Bond_B10ox_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x) into TIC(B11fx-n-B11l)."
	Bond_B10ox_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B10ox_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B10ox_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B10ox_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B10ox_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B10ox_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B10ox_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B10ox_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion008"
	Bond_B10ox_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap008"
	Bond_B10ox_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B10ox
	Bond_B10ox_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B10ox

	// B10xo_BOND
	Bond_B10xo_BondUuid                              TypeAndStructs.BondUUIDType                          = "f3277f00-c2e7-496d-948b-28ce1534785f"
	Bond_B10xo_BondName                              TypeAndStructs.BondNameType                          = "B10xo_BOND"
	Bond_B10xo_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10x*) into TIC(B11f-n-B11lx)."
	Bond_B10xo_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10x*) into TIC(B11f-n-B11lx)."
	Bond_B10xo_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B10xo_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B10xo_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B10xo_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B10xo_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B10xo_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B10xo_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B10xo_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion009"
	Bond_B10xo_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap009"
	Bond_B10xo_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B10xo
	Bond_B10xo_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B10xo

	// B11fx_BOND
	Bond_B11fx_BondUuid                              TypeAndStructs.BondUUIDType                          = "a2361365-1cc6-449b-bcf4-525608adcf8f"
	Bond_B11fx_BondName                              TypeAndStructs.BondNameType                          = "B11fx_BOND_NONE_SWAPPABLE"
	Bond_B11fx_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers before(first) structure."
	Bond_B11fx_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers before(first) structure."
	Bond_B11fx_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B11fx_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B11fx_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B11fx_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B11fx_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B11fx_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B11fx_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B11fx_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion010"
	Bond_B11fx_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap010"
	Bond_B11fx_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B11fx
	Bond_B11fx_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B11fx

	// B11f_BOND
	Bond_B11f_BondUuid                              TypeAndStructs.BondUUIDType                          = "efe03475-91a2-4525-bac1-89fcfb246491"
	Bond_B11f_BondName                              TypeAndStructs.BondNameType                          = "B11f_BOND"
	Bond_B11f_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with one connection which appear as first element inside a TestInstructionContainer"
	Bond_B11f_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with one connection which appear as first element inside a TestInstructionContainer"
	Bond_B11f_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B11f_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B11f_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B11f_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B11f_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B11f_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B11f_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B11f_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion004"
	Bond_B11f_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap004"
	Bond_B11f_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B11f
	Bond_B11f_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B11f

	// B12_BOND
	Bond_B12_BondUuid                              TypeAndStructs.BondUUIDType                          = "78f4c8ef-b8f3-40a2-a335-b28d2ba02e76"
	Bond_B12_BondName                              TypeAndStructs.BondNameType                          = "B12_BOND"
	Bond_B12_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with two connections which connects TestInstructions or TestInstructionContainers"
	Bond_B12_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with two connections which connects TestInstructions or TestInstructionContainers"
	Bond_B12_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B12_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B12_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B12_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B12_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B12_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B12_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B12_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion006"
	Bond_B12_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap006"
	Bond_B12_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B12
	Bond_B12_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B12

	// B11lx_BOND
	Bond_B11lx_BondUuid                              TypeAndStructs.BondUUIDType                          = "b1aaae68-076d-4527-a167-228958c9ab01"
	Bond_B11lx_BondName                              TypeAndStructs.BondNameType                          = "B11lx_BOND_NONE_SWAPPABLE"
	Bond_B11lx_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers after(last) structure."
	Bond_B11lx_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers after(last) structure."
	Bond_B11lx_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B11lx_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B11lx_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B11lx_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B11lx_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B11lx_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B11lx_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B11lx_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion011"
	Bond_B11lx_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap011"
	Bond_B11lx_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B11lx
	Bond_B11lx_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B11lx

	// B10_BOND
	Bond_B10_BondUuid                              TypeAndStructs.BondUUIDType                          = "0883d538-1cff-4be1-ba1f-4dc1f68f6242"
	Bond_B10_BondName                              TypeAndStructs.BondNameType                          = "B10_BOND"
	Bond_B10_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with zero connections which appear as an element in a new TestInstructionContainer without any other elements"
	Bond_B10_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with zero connections which appear as an element in a new TestInstructionContainer without any other elements"
	Bond_B10_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B10_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B10_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B10_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B10_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B10_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B10_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B10_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion003"
	Bond_B10_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap003"
	Bond_B10_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B10
	Bond_B10_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B10

	// B11l_BOND
	Bond_B11l_BondUuid                              TypeAndStructs.BondUUIDType                          = "2c45d283-4625-40e3-bf8e-cb070f1a6210"
	Bond_B11l_BondName                              TypeAndStructs.BondNameType                          = "B11l_BOND"
	Bond_B11l_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with one connection which appear as last element inside a TestInstructionContainer"
	Bond_B11l_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with one connection which appear as last element inside a TestInstructionContainer"
	Bond_B11l_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B11l_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B11l_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B11l_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B11l_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B11l_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B11l_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B11l_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion005"
	Bond_B11l_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap005"
	Bond_B11l_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B11l
	Bond_B11l_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B11l

	// B10oxo_BOND
	Bond_B10oxo_BondUuid                              TypeAndStructs.BondUUIDType                          = "f1cc879d-3b95-4dea-84e5-e2a166c5fc95"
	Bond_B10oxo_BondName                              TypeAndStructs.BondNameType                          = "B10oxo_BOND"
	Bond_B10oxo_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x*) into TIC(B11fx-n-B11lx)."
	Bond_B10oxo_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x*) into TIC(B11fx-n-B11lx)."
	Bond_B10oxo_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B10oxo_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B10oxo_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B10oxo_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B10oxo_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B10oxo_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B10oxo_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B10oxo_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion007"
	Bond_B10oxo_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap007"
	Bond_B10oxo_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B10oxo
	Bond_B10oxo_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B10oxo

	// B1l_BOND
	Bond_B1l_BondUuid                              TypeAndStructs.BondUUIDType                          = "123ab7f8-9f08-44ad-bc02-15b205146f48"
	Bond_B1l_BondName                              TypeAndStructs.BondNameType                          = "B1l_BOND_NONE_SWAPPABLE"
	Bond_B1l_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers"
	Bond_B1l_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with one connection. Use for starting and ending a TestCase and is only found in pre-created TestInstructionContainers"
	Bond_B1l_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B1l_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B1l_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B1l_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B1l_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B1l_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B1l_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B1l_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion002"
	Bond_B1l_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap002"
	Bond_B1l_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B1l
	Bond_B1l_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B1l

	// B12x_BOND
	Bond_B12x_BondUuid                              TypeAndStructs.BondUUIDType                          = "9e613527-daa6-4559-9c73-adb8213eccd9"
	Bond_B12x_BondName                              TypeAndStructs.BondNameType                          = "B12x_BOND_NONE_SWAPPABLE"
	Bond_B12x_BondDescription                       TypeAndStructs.BondDescriptionType                   = "Bond with two connections which connects TestInstructions or TestInstructiB12x_BOND_NONE_SWAPPABLEonContainers. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers within a pre-created container."
	Bond_B12x_BondMouseOverText                     TypeAndStructs.BondMouseOverTextType                 = "Bond with two connections which connects TestInstructions or TestInstructiB12x_BOND_NONE_SWAPPABLEonContainers. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructions or TestInstructionsContainers within a pre-created container."
	Bond_B12x_Deprecated                            TypeAndStructs.BondDeprecatedType                    = false
	Bond_B12x_Enabled                               TypeAndStructs.BondEnabledType                       = true
	Bond_B12x_Visible                               TypeAndStructs.BondVisibleType                       = true
	Bond_B12x_BondColor                             TypeAndStructs.BondColorType                         = "#ffff00"
	Bond_B12x_CanBeDeleted                          TypeAndStructs.BondCanBeDeletedType                  = false
	Bond_B12x_CanBeSwappedOut                       TypeAndStructs.BondCanBeSwappedOutType               = false
	Bond_B12x_ShowBondAttributes                    TypeAndStructs.BondShowBondAttributesType            = true
	Bond_B12x_TCRuleDeletion                        TypeAndStructs.BondTCRuleDeletionType                = "TCRuleDeletion012"
	Bond_B12x_TCRuleSwap                            TypeAndStructs.BondTCRuleSwapType                    = "TCRuleSwap012"
	Bond_B12x_TestCaseModelElementType              TypeAndStructs.TestCaseModelElementTypeType          = TestCaseModelElementTypes.TestCaseModelElementType_B12x
	Bond_B12x_TestCaseModelElementTypeGrpcMappingID TypeAndStructs.TestCaseModelElementGrpcMappingIdType = TestCaseModelElementTypes.TestCaseModelElementGrpcMappingId_B12x
)

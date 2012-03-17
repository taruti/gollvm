package llvm

/*
#define __STDC_LIMIT_MACROS
#define __STDC_CONSTANT_MACROS
#include <llvm-c/Transforms/IPO.h>
*/
import "C"

// helpers
func boolToUnsigned(b bool) C.unsigned {
	if b {
		return 1
	}
	return 0
}

func (pm PassManager) AddArgumentPromotionPass()     { C.LLVMAddArgumentPromotionPass(pm.C) }
func (pm PassManager) AddConstantMergePass()         { C.LLVMAddConstantMergePass(pm.C) }
func (pm PassManager) AddDeadArgEliminationPass()    { C.LLVMAddDeadArgEliminationPass(pm.C) }
func (pm PassManager) AddFunctionAttrsPass()         { C.LLVMAddFunctionAttrsPass(pm.C) }
func (pm PassManager) AddFunctionInliningPass()      { C.LLVMAddFunctionInliningPass(pm.C) }
func (pm PassManager) AddGlobalDCEPass()             { C.LLVMAddGlobalDCEPass(pm.C) }
func (pm PassManager) AddGlobalOptimizerPass()       { C.LLVMAddGlobalOptimizerPass(pm.C) }
func (pm PassManager) AddIPConstantPropagationPass() { C.LLVMAddIPConstantPropagationPass(pm.C) }
func (pm PassManager) AddPruneEHPass()               { C.LLVMAddPruneEHPass(pm.C) }
func (pm PassManager) AddIPSCCPPass()                { C.LLVMAddIPSCCPPass(pm.C) }
func (pm PassManager) AddInternalizePass(allButMain bool) {
	C.LLVMAddInternalizePass(pm.C, boolToUnsigned(allButMain))
}
func (pm PassManager) AddStripDeadPrototypesPass() { C.LLVMAddStripDeadPrototypesPass(pm.C) }

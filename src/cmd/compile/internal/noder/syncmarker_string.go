// Code generated by "stringer -type=syncMarker -trimprefix=sync"; DO NOT EDIT.

package noder

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[syncNode-1]
	_ = x[syncBool-2]
	_ = x[syncInt64-3]
	_ = x[syncUint64-4]
	_ = x[syncString-5]
	_ = x[syncPos-6]
	_ = x[syncPkg-7]
	_ = x[syncSym-8]
	_ = x[syncSelector-9]
	_ = x[syncKind-10]
	_ = x[syncType-11]
	_ = x[syncTypePkg-12]
	_ = x[syncSignature-13]
	_ = x[syncParam-14]
	_ = x[syncOp-15]
	_ = x[syncObject-16]
	_ = x[syncExpr-17]
	_ = x[syncStmt-18]
	_ = x[syncDecl-19]
	_ = x[syncConstDecl-20]
	_ = x[syncFuncDecl-21]
	_ = x[syncTypeDecl-22]
	_ = x[syncVarDecl-23]
	_ = x[syncPragma-24]
	_ = x[syncValue-25]
	_ = x[syncEOF-26]
	_ = x[syncMethod-27]
	_ = x[syncFuncBody-28]
	_ = x[syncUse-29]
	_ = x[syncUseObj-30]
	_ = x[syncObjectIdx-31]
	_ = x[syncTypeIdx-32]
	_ = x[syncBOF-33]
	_ = x[syncEntry-34]
	_ = x[syncOpenScope-35]
	_ = x[syncCloseScope-36]
	_ = x[syncGlobal-37]
	_ = x[syncLocal-38]
	_ = x[syncDefine-39]
	_ = x[syncDefLocal-40]
	_ = x[syncUseLocal-41]
	_ = x[syncDefGlobal-42]
	_ = x[syncUseGlobal-43]
	_ = x[syncTypeParams-44]
	_ = x[syncUseLabel-45]
	_ = x[syncDefLabel-46]
	_ = x[syncFuncLit-47]
	_ = x[syncCommonFunc-48]
	_ = x[syncBodyRef-49]
	_ = x[syncLinksymExt-50]
	_ = x[syncHack-51]
	_ = x[syncSetlineno-52]
	_ = x[syncName-53]
	_ = x[syncImportDecl-54]
	_ = x[syncDeclNames-55]
	_ = x[syncDeclName-56]
	_ = x[syncExprList-57]
	_ = x[syncExprs-58]
	_ = x[syncWrapname-59]
	_ = x[syncTypeExpr-60]
	_ = x[syncTypeExprOrNil-61]
	_ = x[syncChanDir-62]
	_ = x[syncParams-63]
	_ = x[syncCloseAnotherScope-64]
	_ = x[syncSum-65]
	_ = x[syncUnOp-66]
	_ = x[syncBinOp-67]
	_ = x[syncStructType-68]
	_ = x[syncInterfaceType-69]
	_ = x[syncPackname-70]
	_ = x[syncEmbedded-71]
	_ = x[syncStmts-72]
	_ = x[syncStmtsFall-73]
	_ = x[syncStmtFall-74]
	_ = x[syncBlockStmt-75]
	_ = x[syncIfStmt-76]
	_ = x[syncForStmt-77]
	_ = x[syncSwitchStmt-78]
	_ = x[syncRangeStmt-79]
	_ = x[syncCaseClause-80]
	_ = x[syncCommClause-81]
	_ = x[syncSelectStmt-82]
	_ = x[syncDecls-83]
	_ = x[syncLabeledStmt-84]
	_ = x[syncCompLit-85]
	_ = x[sync1-86]
	_ = x[sync2-87]
	_ = x[sync3-88]
	_ = x[sync4-89]
	_ = x[syncN-90]
	_ = x[syncDefImplicit-91]
	_ = x[syncUseName-92]
	_ = x[syncUseObjLocal-93]
	_ = x[syncAddLocal-94]
	_ = x[syncBothSignature-95]
	_ = x[syncSetUnderlying-96]
	_ = x[syncLinkname-97]
	_ = x[syncStmt1-98]
	_ = x[syncStmtsEnd-99]
	_ = x[syncDeclare-100]
	_ = x[syncTopDecls-101]
	_ = x[syncTopConstDecl-102]
	_ = x[syncTopFuncDecl-103]
	_ = x[syncTopTypeDecl-104]
	_ = x[syncTopVarDecl-105]
	_ = x[syncObject1-106]
	_ = x[syncAddBody-107]
	_ = x[syncLabel-108]
	_ = x[syncFuncExt-109]
	_ = x[syncMethExt-110]
	_ = x[syncOptLabel-111]
	_ = x[syncScalar-112]
	_ = x[syncStmtDecls-113]
	_ = x[syncDeclLocal-114]
	_ = x[syncObjLocal-115]
	_ = x[syncObjLocal1-116]
	_ = x[syncDeclareLocal-117]
	_ = x[syncPublic-118]
	_ = x[syncPrivate-119]
	_ = x[syncRelocs-120]
	_ = x[syncReloc-121]
	_ = x[syncUseReloc-122]
	_ = x[syncVarExt-123]
	_ = x[syncPkgDef-124]
	_ = x[syncTypeExt-125]
	_ = x[syncVal-126]
	_ = x[syncCodeObj-127]
	_ = x[syncPosBase-128]
	_ = x[syncLocalIdent-129]
}

const _syncMarker_name = "NodeBoolInt64Uint64StringPosPkgSymSelectorKindTypeTypePkgSignatureParamOpObjectExprStmtDeclConstDeclFuncDeclTypeDeclVarDeclPragmaValueEOFMethodFuncBodyUseUseObjObjectIdxTypeIdxBOFEntryOpenScopeCloseScopeGlobalLocalDefineDefLocalUseLocalDefGlobalUseGlobalTypeParamsUseLabelDefLabelFuncLitCommonFuncBodyRefLinksymExtHackSetlinenoNameImportDeclDeclNamesDeclNameExprListExprsWrapnameTypeExprTypeExprOrNilChanDirParamsCloseAnotherScopeSumUnOpBinOpStructTypeInterfaceTypePacknameEmbeddedStmtsStmtsFallStmtFallBlockStmtIfStmtForStmtSwitchStmtRangeStmtCaseClauseCommClauseSelectStmtDeclsLabeledStmtCompLit1234NDefImplicitUseNameUseObjLocalAddLocalBothSignatureSetUnderlyingLinknameStmt1StmtsEndDeclareTopDeclsTopConstDeclTopFuncDeclTopTypeDeclTopVarDeclObject1AddBodyLabelFuncExtMethExtOptLabelScalarStmtDeclsDeclLocalObjLocalObjLocal1DeclareLocalPublicPrivateRelocsRelocUseRelocVarExtPkgDefTypeExtValCodeObjPosBaseLocalIdent"

var _syncMarker_index = [...]uint16{0, 4, 8, 13, 19, 25, 28, 31, 34, 42, 46, 50, 57, 66, 71, 73, 79, 83, 87, 91, 100, 108, 116, 123, 129, 134, 137, 143, 151, 154, 160, 169, 176, 179, 184, 193, 203, 209, 214, 220, 228, 236, 245, 254, 264, 272, 280, 287, 297, 304, 314, 318, 327, 331, 341, 350, 358, 366, 371, 379, 387, 400, 407, 413, 430, 433, 437, 442, 452, 465, 473, 481, 486, 495, 503, 512, 518, 525, 535, 544, 554, 564, 574, 579, 590, 597, 598, 599, 600, 601, 602, 613, 620, 631, 639, 652, 665, 673, 678, 686, 693, 701, 713, 724, 735, 745, 752, 759, 764, 771, 778, 786, 792, 801, 810, 818, 827, 839, 845, 852, 858, 863, 871, 877, 883, 890, 893, 900, 907, 917}

func (i syncMarker) String() string {
	i -= 1
	if i < 0 || i >= syncMarker(len(_syncMarker_index)-1) {
		return "syncMarker(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _syncMarker_name[_syncMarker_index[i]:_syncMarker_index[i+1]]
}
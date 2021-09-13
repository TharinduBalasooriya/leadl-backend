//
// Created by tharindu on 8/27/2021.
//

#include <string>
#include "LDAL_Wrapper.h"
#include "DefFileReader.h"
#include "ScriptReader.h"
#include "MetaData.h"
#include "ExecutionTemplateList.h"
#include "ExecutionContext.h"
#include "Node.h"
#include "MemMan.h"
#include "ELInterpretter.h"
#include "CommonIncludes.h"
#include "Debugger.h"


std::string LDAL_Wrapper::GetLDALResult(std::string defFilePath) {

    DefFileReader dfr;
    MetaData *pMD = dfr.Read(defFilePath);

    std::cout << defFilePath;
    ScriptReader sr;
    ScriptReaderOutput op;
    bool bSucc = sr.ProcessScript(pMD->s_RuleFileName, pMD, op);
    if (!bSucc) {
        std::wcout << "\nFailed to read script\n";
    }
    ExecutionContext ec;
    ec.p_mapFunctions = &op.map_Functions;
    ec.p_MD = pMD;
    Node *pLog = MemoryManager::Inst.CreateNode(1);
    Node *pY = MemoryManager::Inst.CreateNode(2);
    Node *pOut = MemoryManager::Inst.CreateNode(3);

    ec.map_Var["LOG"] = pLog;
    ec.map_Var["OUTPUT"] = pOut;
    ec.map_Var["Y"] = pY;

    MSTRING Location = pMD->s_TREELocation;
    pLog->ReadValueFromFile(Location.c_str());


    op.p_ETL->Execute(&ec);


    MSTRING result = pOut->GetAggregatedValue();
    Debugger b;
    b.DebugResult(&ec.map_Var,pMD);


    pLog->DestroyWithSubTree();
    pY->DestroyWithSubTree();
    pOut->DestroyWithSubTree();


    return result;


}
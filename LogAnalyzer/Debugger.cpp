//
// Created by Tharindu-Balasuriya on 9/7/2021.
//


#include <iomanip>
#include "Debugger.h"
#include "ExecutionContext.h"
#include "Null.h"
#include "string"
#include "json.hpp"
#include "Int.h"
#include "Node.h"
#include "StringOperations.h"
#include "Bool.h"
#include "EntityList.h"

using json = nlohmann::json;

void Debugger::DebugResult(MAP_STR_ENTITYPTR *ecVarMap,MetaData *pMD) {


    json j;
    json varibleObj;


    for (auto const &x : *ecVarMap) {


        if(x.first !="LOG"){

            switch (x.second->ul_Type) {


                case ENTITY_TYPE_NULL:{

                    //ep:entity pointer
                    PNull ep = (PNull)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "NULL ENTITY";

                    MSTRING val;
                    if(ep->IsNull()){
                        val="TRUE";
                    }
                    varibleObj["details"] =  "Is Null : " +val ;

                    j["variables"] += varibleObj;

                    break;
                }
                case ENTITY_TYPE_INT:{
                    PInt ep = (PInt)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "INTEGER";

                    if(ep->IsNull()){
                        varibleObj["details"] = "VALUE:NULL" ;
                    }else{
                        varibleObj["details"] = "VALUE: " + std::to_string(ep->GetValue());
                    }

                    j["variables"] += varibleObj;

                    break;
                }
                case ENTITY_TYPE_STRING:{

                    PString ep = (PString)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "STRING";
                    if(ep->IsNull()){
                        varibleObj["details"] = "VALUE:NULL" ;
                    }else{
                        varibleObj["details"] = "VALUE: " + ep->GetValue();
                    }
                    break;
                }
                case ENTITY_TYPE_NODE:{
                    PNODE ep = (PNODE)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "NODE";
                    json  nodeDetailsObj;
                    if(ep->IsNull()){
                        varibleObj["details"] = "VALUE:NULL NODE" ;
                    }else{

                         if(ep->GetValue()){
                             const char *s = ep->GetValue();
                             std::string zValue(s);
                             nodeDetailsObj["z_Value"]=zValue;

                         }else{

                             nodeDetailsObj["z_Value"]="NULL";
                         }
                        if(ep->GetLVal()){
                            const char *s = ep->GetLVal();
                            std::string lValue(s);
                            nodeDetailsObj["l_Value"]=lValue;
                        }else{
                            nodeDetailsObj["l_Value"]="NULL";
                        }


                         if(ep->GetRVal()){
                             const char *s = ep->GetRVal();
                             std::string rValue(s);
                             nodeDetailsObj["r_Value"]=rValue;
                         }else{
                             nodeDetailsObj["r_Value"]="NULL";
                         }

                         if(ep->GetCustomString()){
                             const char *s = ep->GetCustomString();
                             std::string cValue(s);
                             nodeDetailsObj["customString"]=cValue;
                         }else{
                             nodeDetailsObj["customString"]="NULL";
                         }

                         nodeDetailsObj["child_count"] = ep->GetChildCount();
                        varibleObj["details"] = nodeDetailsObj;

                    }
                    j["variables"] += varibleObj;
                    break;
                }
                case ENTITY_TYPE_LIST:{

                    PENTITYLIST ep = (PENTITYLIST)x.second;
                    varibleObj["name"] = x.first;
                    varibleObj["dataType"] = "LIST";
                    if(ep->IsNull()){
                        varibleObj["details"] = "VALUE:NULL" ;
                    }else{
                        varibleObj["details"] = "LIST SIZE : " + std::to_string(ep->size());
                    }
                    break;

                }

            }
        }


    }
  
    std::ofstream o(pMD->s_DebugJSON_File);
    o << std::setw(4) << j <<"\n";



}
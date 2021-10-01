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
#include "Strings.h"

using json = nlohmann::json;

void Debugger::DebugResult(MAP_STR_ENTITYPTR *ecVarMap,MetaData *pMD) {


    json j;
    json varibleObj;
    json childObj;


    for (auto const &x : *ecVarMap) {


//        if(x.first !="LOG" && x.first !="X"){

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
                    varibleObj["details"] = ep->GetValue();
                    j["variables"] += varibleObj;
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

                         nodeDetailsObj["child"];
                         PNODE child = ep->GetFirstChild();
                         while(child != NULL){

                             if(child->GetValue()){
                                 const char *s = ep->GetValue();
                                 std::string zValue(s);
                                 childObj["z_Value"]=zValue;

                             }else{

                                 childObj["z_Value"]="NULL";
                             }
                             if(ep->GetLVal()){
                                 const char *s = ep->GetLVal();
                                 std::string lValue(s);
                                 childObj["l_Value"]=lValue;
                             }else{
                                 childObj["l_Value"]="NULL";
                             }


                             if(ep->GetRVal()){
                                 const char *s = ep->GetRVal();
                                 std::string rValue(s);
                                 childObj["r_Value"]=rValue;
                             }else{
                                 childObj["r_Value"]="NULL";
                             }

                             if(ep->GetCustomString()){
                                 const char *s = ep->GetCustomString();
                                 std::string cValue(s);
                                 childObj["customString"]=cValue;
                             }else{
                                 childObj["customString"]="NULL";
                             }
                             nodeDetailsObj["child"]+=childObj;

                             child = child->GetRightSibling();
                         }
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
                    j["variables"] += varibleObj;
                    break;
                }

            }
        }


//    }


  
    std::ofstream o(pMD->s_DebugJSON_File);
    o << std::setw(4) << j <<"\n";

}

void appendNode(json &j ,json &varibleObj,const std::pair<const std::basic_string<char>, PENTITY> &x){

}
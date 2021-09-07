//
// Created by Isini_Dananjana on 2021-07-12.
//


#include "FCLWrapper.h"
#include "ELInterpretter.h"
#include "CommonIncludes.h"
#include "CommonIncludes.h"
#include "LDAL_Wrapper.h"
#include <string>



void FCLWrapper::RunELInterpretter(const char *defFilepath) {
    ELInterpretter intp;
    intp.EvaluateCase(defFilepath);

}

std::string FCLWrapper::GetLDALResult(const char *defFilePath) {
    LDAL_Wrapper ldalWrapper ;
    return ldalWrapper.GetLDALResult(defFilePath);


}
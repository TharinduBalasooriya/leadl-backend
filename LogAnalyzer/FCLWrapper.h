//
// Created by Isini-Dananjana on 2021-07-12.
//


#ifndef CODE_FCLWRAPPER_H
#define CODE_FCLWRAPPER_H
#include <string>

class FCLWrapper{
public:
    void RunELInterpretter(const char* defFilepath);
    std::string GetLDALResult(const char* defFilePath);

};

#endif //CODE_FCLWRAPPER_H

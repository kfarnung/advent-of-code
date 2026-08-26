#ifndef DAY11_LIB_H
#define DAY11_LIB_H

#include <string>
#include <vector>

namespace day11
{
    void increment_password(std::vector<char> &input);
    bool validate_password(const std::vector<char> &input);
    std::string next_valid_password(const std::string &input);
}

#endif
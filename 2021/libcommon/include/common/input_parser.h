#ifndef COMMON_INPUT_PARSER_H
#define COMMON_INPUT_PARSER_H

#include <string>
#include <vector>

namespace common
{
    std::vector<std::string> splitlines(std::istream &input);
    std::vector<std::string> splitlines(const char *input);
    std::vector<std::string> splitstr(const std::string &input, char ch);
}

#endif
#include <common/string_convert.h>

int common::string_to_int(const std::string &input)
{
    return std::atoi(input.c_str());
}
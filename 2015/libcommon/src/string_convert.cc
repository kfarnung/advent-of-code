#include <common/string_convert.h>

#include <cstdlib>

int common::string_to_int(const std::string &input)
{
    return static_cast<int>(std::strtol(input.c_str(), nullptr, 10));
}
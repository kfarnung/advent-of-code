#ifndef DAY08_LIB_H
#define DAY08_LIB_H

#include <string>
#include <vector>

namespace day08
{
    size_t calculate_string_overhead(const std::vector<std::string> &input);
    size_t calculate_encoding_overhead(const std::vector<std::string> &input);
}

#endif
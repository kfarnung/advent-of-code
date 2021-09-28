#ifndef COMMON_VECTOR_CONVERT_H
#define COMMON_VECTOR_CONVERT_H

#include <string>
#include <vector>

namespace common
{
    std::vector<int32_t> vector_parse_int(const std::vector<std::string> &input);
}

#endif
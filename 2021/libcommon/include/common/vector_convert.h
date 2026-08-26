#ifndef COMMON_VECTOR_CONVERT_H
#define COMMON_VECTOR_CONVERT_H

#include <string>
#include <vector>

namespace common
{
    std::vector<int64_t> vector_parse_int(const std::vector<std::string> &input);
}

#endif
#include <common/vector_convert.h>

#include <algorithm>
#include <iterator>

namespace
{
    int32_t string_to_int(const std::string &input)
    {
        return std::stol(input);
    }
}

std::vector<int32_t> common::vector_parse_int(const std::vector<std::string> &input)
{
    std::vector<int32_t> output;
    std::transform(std::begin(input), std::end(input), std::back_inserter(output), string_to_int);
    return output;
}
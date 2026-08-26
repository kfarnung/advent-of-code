#include <common/vector_convert.h>

#include <algorithm>
#include <iterator>

namespace
{
    int64_t string_to_int64(const std::string &input)
    {
        return std::stoll(input);
    }
}

std::vector<int64_t> common::vector_parse_int(const std::vector<std::string> &input)
{
    std::vector<int64_t> output;
    std::transform(std::begin(input), std::end(input), std::back_inserter(output), string_to_int64);
    return output;
}
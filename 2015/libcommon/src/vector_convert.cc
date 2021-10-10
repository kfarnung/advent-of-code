#include <common/vector_convert.h>

#include <algorithm>
#include <iterator>

namespace
{
    int str_to_int(std::string str)
    {
        return std::atoi(str.c_str());
    }
}

std::vector<int> common::vector_parse_int(std::vector<std::string> &input)
{
    std::vector<int> output;
    std::transform(std::begin(input), std::end(input), std::back_inserter(output), str_to_int);
    return output;
}
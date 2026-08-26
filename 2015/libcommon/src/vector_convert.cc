#include <common/vector_convert.h>

#include <common/string_convert.h>

#include <algorithm>
#include <iterator>

std::vector<int> common::vector_parse_int(std::vector<std::string> &input)
{
    std::vector<int> output;
    std::transform(std::begin(input), std::end(input), std::back_inserter(output), common::string_to_int);
    return output;
}
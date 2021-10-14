#ifndef DAY05_LIB_H
#define DAY05_LIB_H

#include <string>
#include <vector>

namespace day05
{
    bool is_nice_string_part1(const std::string &input);
    size_t count_nice_strings_part1(const std::vector<std::string> &input);
    bool is_nice_string_part2(const std::string &input);
    size_t count_nice_strings_part2(const std::vector<std::string> &input);
}

#endif
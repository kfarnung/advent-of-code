#ifndef DAY14_LIB_H
#define DAY14_LIB_H

#include <string>
#include <vector>

namespace day14
{
    int32_t find_longest_distance(const std::vector<std::string> &input, int32_t target_time);
    int32_t find_part2_winner(const std::vector<std::string> &input, int32_t target_time);
}

#endif
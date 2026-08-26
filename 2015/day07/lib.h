#ifndef DAY07_LIB_H
#define DAY07_LIB_H

#include <cstdint>
#include <string>
#include <vector>

namespace day07
{
    uint16_t get_wire_output(const std::vector<std::string> &input, const std::string &wire);
    uint16_t get_wire_output_part2(const std::vector<std::string> &input, const std::string &wire);
}

#endif
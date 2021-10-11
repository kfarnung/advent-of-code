#ifndef DAY02_LIB_H
#define DAY02_LIB_H

#include <string>
#include <vector>

namespace day02
{
    int64_t calculate_wrapping_paper(const std::string &input);
    int64_t calculate_wrapping_paper(const std::vector<std::string> &input);
    int64_t calculate_ribbon(const std::string &input);
    int64_t calculate_ribbon(const std::vector<std::string> &input);
}

#endif
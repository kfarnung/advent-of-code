#include "lib.h"

#include <common/input_parser.h>
#include <common/vector_convert.h>

#include <algorithm>
#include <limits>
#include <numeric>

namespace
{
    int32_t get_middle(const std::vector<int32_t> &input)
    {
        auto size = input.size();
        auto middle = size / 2;

        if (size % 2 == 0)
        {
            return (input[middle - 1] + input[middle]) / 2;
        }

        return input[middle];
    }

    uint32_t calculate_fuel_part2(const std::vector<int32_t> &input, int32_t target)
    {
        uint32_t fuel_cost = 0;
        for (const auto &sub : input)
        {
            auto distance = std::abs(target - sub);
            fuel_cost += distance * (distance + 1) / 2;
        }

        return fuel_cost;
    }
}

uint32_t day07::run_part1(const std::vector<std::string> &input)
{
    auto submarines = common::vector_parse_int(common::splitstr(input[0], ','));
    std::sort(begin(submarines), end(submarines));
    auto median = get_middle(submarines);

    uint32_t fuel_cost = 0;
    for (const auto &sub : submarines)
    {
        fuel_cost += std::abs(median - sub);
    }

    return fuel_cost;
}

uint32_t day07::run_part2(const std::vector<std::string> &input)
{
    auto submarines = common::vector_parse_int(common::splitstr(input[0], ','));

    auto total = std::accumulate(begin(submarines), end(submarines), 0);
    auto mean = total / static_cast<double>(submarines.size());
    auto lower = static_cast<int32_t>(mean - 0.5);
    auto upper = static_cast<int32_t>(mean + 0.5);

    return std::min(
        calculate_fuel_part2(submarines, lower),
        calculate_fuel_part2(submarines, upper));
}

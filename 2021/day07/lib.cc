#include "lib.h"

#include <common/input_parser.h>
#include <common/vector_convert.h>

#include <algorithm>
#include <limits>

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

    std::pair<int32_t, int32_t> get_range(const std::vector<int32_t> &input)
    {
        auto min = std::numeric_limits<int32_t>::max();
        auto max = std::numeric_limits<int32_t>::min();

        for (const auto &value : input)
        {
            min = std::min(min, value);
            max = std::max(max, value);
        }

        return std::make_pair(min, max);
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
    auto range = get_range(submarines);

    auto min_fuel = std::numeric_limits<uint32_t>::max();

    for (int32_t i = range.first; i <= range.second; ++i)
    {
        uint32_t fuel_cost = 0;
        for (const auto &sub : submarines)
        {
            auto distance = std::abs(i - sub);
            fuel_cost += distance * (distance + 1) / 2;
        }

        min_fuel = std::min(min_fuel, fuel_cost);
    }

    return min_fuel;
}

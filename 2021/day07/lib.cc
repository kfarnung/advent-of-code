#include "lib.h"

#include <common/input_parser.h>
#include <common/vector_convert.h>

#include <algorithm>
#include <limits>
#include <numeric>

namespace
{
    int64_t get_middle(const std::vector<int64_t> &input)
    {
        auto size = input.size();
        auto middle = size / 2;

        if (size % 2 == 0)
        {
            return (input[middle - 1] + input[middle]) / 2;
        }

        return input[middle];
    }

    int64_t calculate_fuel_part2(const std::vector<int64_t> &input, int64_t target)
    {
        int64_t fuel_cost = 0;
        for (const auto &sub : input)
        {
            auto distance = std::abs(target - sub);
            fuel_cost += distance * (distance + 1) / 2;
        }

        return fuel_cost;
    }
}

int64_t day07::run_part1(const std::vector<std::string> &input)
{
    auto submarines = common::vector_parse_int(common::splitstr(input[0], ','));
    std::sort(begin(submarines), end(submarines));
    auto median = get_middle(submarines);

    int64_t fuel_cost = 0;
    for (const auto &sub : submarines)
    {
        fuel_cost += std::abs(median - sub);
    }

    return fuel_cost;
}

int64_t day07::run_part2(const std::vector<std::string> &input)
{
    auto submarines = common::vector_parse_int(common::splitstr(input[0], ','));

    auto total = std::accumulate(begin(submarines), end(submarines), 0LL);
    auto mean = total / static_cast<double>(submarines.size());
    auto lower = static_cast<int64_t>(mean - 0.5);
    auto upper = static_cast<int64_t>(mean + 0.5);

    return std::min(
        calculate_fuel_part2(submarines, lower),
        calculate_fuel_part2(submarines, upper));
}

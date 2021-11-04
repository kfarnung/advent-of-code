#include "lib.h"

#include <common/string_convert.h>

#include <set>

namespace
{
    size_t count_combinations(
        const std::set<std::pair<size_t, int32_t>> &remaining,
        size_t path_length,
        int32_t current_sum,
        int32_t target_sum,
        size_t container_limit,
        size_t &min_required)
    {
        if (current_sum == target_sum)
        {
            min_required = std::min(min_required, path_length);
            return 1;
        }
        
        if (current_sum > target_sum || path_length >= container_limit)
        {
            return 0;
        }

        size_t count = 0;
        std::set<std::pair<size_t, int32_t>> next_remaining{remaining};

        for (const auto &entry : remaining)
        {
            next_remaining.erase(entry);
            count += count_combinations(next_remaining, path_length + 1, current_sum + entry.second, target_sum, container_limit, min_required);
        }

        return count;
    }
}

size_t day17::count_combinations(const std::vector<std::string> &input, int32_t target_sum, bool use_fewest)
{
    std::set<std::pair<size_t, int32_t>> containers;
    for (size_t i = 0; i < input.size(); i++)
    {
        containers.insert({i, common::string_to_int(input[i])});
    }

    auto min_required = containers.size();
    auto count = ::count_combinations(
        containers,
        0,
        0,
        target_sum,
        min_required,
        min_required);

    if (use_fewest)
    {
        count = ::count_combinations(
            containers,
            0,
            0,
            target_sum,
            min_required,
            min_required);
    }

    return count;
}
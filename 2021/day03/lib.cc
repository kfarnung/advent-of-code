#include "lib.h"

namespace
{
    size_t get_one_count(const std::vector<std::string> &input, size_t position)
    {
        size_t count = 0;
        for (const auto &current : input)
        {
            if (current[position] == '1')
            {
                ++count;
            }
        }

        return count;
    }
}

int64_t day03::run_part1(const std::vector<std::string> &input)
{
    size_t input_size = input.size();
    size_t line_size = input[0].size();

    int64_t gamma = 0;
    int64_t epsilon = 0;

    for (size_t i = 0; i < line_size; ++i)
    {
        gamma <<= 1;
        epsilon <<= 1;

        size_t one_count = get_one_count(input, i);

        if (one_count >= input_size - one_count)
        {
            gamma |= 1;
        }

        if (one_count <= input_size - one_count)
        {
            epsilon |= 1;
        }
    }

    return gamma * epsilon;
}

int64_t day03::run_part2(const std::vector<std::string> &input)
{
    size_t count_index = 0;
    std::vector<std::string> oxygen_list(input);
    std::vector<std::string> co2_list(input);

    while (oxygen_list.size() > 1)
    {
        std::vector<std::string> new_oxygen;

        size_t one_count = get_one_count(oxygen_list, count_index);
        char ch = one_count >= (oxygen_list.size() - one_count) ? '1' : '0';

        for (const auto &current : oxygen_list)
        {
            if (current[count_index] == ch)
            {
                new_oxygen.push_back(current);
            }
        }

        std::swap(oxygen_list, new_oxygen);
        ++count_index;
    }

    count_index = 0;
    while (co2_list.size() > 1)
    {
        std::vector<std::string> new_co2;
        
        size_t one_count = get_one_count(co2_list, count_index);
        char ch = one_count >= (co2_list.size() - one_count) ? '0' : '1';

        for (const auto &current : co2_list)
        {
            if (current[count_index] == ch)
            {
                new_co2.push_back(current);
            }
        }

        std::swap(co2_list, new_co2);
        ++count_index;
    }

    return std::stoul(oxygen_list[0], nullptr, 2) * std::stoul(co2_list[0], nullptr, 2);
}

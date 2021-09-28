#include "lib.h"

#include <common/vector_convert.h>

#include <limits>

size_t day01::count_increases(const std::vector<std::string> &input, size_t window_size)
{
    size_t increases = 0;

    auto values = common::vector_parse_int(input);
    int32_t previous = std::numeric_limits<int32_t>::max();

    for (size_t i = window_size - 1; i < input.size(); ++i)
    {
        int sum = 0;
        for (size_t j = 0; j < window_size; ++j)
        {
            sum += values[i - j];
        }

        if (sum > previous)
        {
            ++increases;
        }

        previous = sum;
    }

    return increases;
}

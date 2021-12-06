#include "lib.h"

#include <common/input_parser.h>

#include <deque>
#include <numeric>

size_t day06::simulate_lanternfish(const std::vector<std::string> &input, size_t day_count)
{
    std::deque<size_t> queue(9, 0);

    for (const auto &fish : common::splitstr(input[0], ','))
    {
        auto index = std::stoul(fish);
        queue[index] += 1;
    }

    for (size_t day = 0; day < day_count; ++day)
    {
        auto spawner_count = queue.front();
        queue.pop_front();

        queue[6] += spawner_count;
        queue.push_back(spawner_count);
    }

    return std::accumulate(begin(queue), end(queue), 0ULL);
}

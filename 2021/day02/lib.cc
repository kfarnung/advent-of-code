#include "lib.h"

int64_t day02::run_part1(const std::vector<std::string> &input)
{
    int64_t horizontal = 0;
    int64_t vertical = 0;

    for (const auto &line : input)
    {
        if (line.find("forward ") == 0)
        {
            int64_t value = std::stoi(line.substr(8));
            horizontal += value;
        }
        else if (line.find("down ") == 0)
        {
            int64_t value = std::stoi(line.substr(5));
            vertical += value;
        }
        else if (line.find("up ") == 0)
        {
            int64_t value = std::stoi(line.substr(3));
            vertical -= value;
        }
    }

    return horizontal * vertical;
}

int64_t day02::run_part2(const std::vector<std::string> &input)
{
    int64_t horizontal = 0;
    int64_t vertical = 0;
    int64_t aim = 0;

    for (const auto &line : input)
    {
        if (line.find("forward ") == 0)
        {
            int64_t value = std::stoi(line.substr(8));
            horizontal += value;
            vertical += value * aim;
        }
        else if (line.find("down ") == 0)
        {
            int64_t value = std::stoi(line.substr(5));
            aim += value;
        }
        else if (line.find("up ") == 0)
        {
            int64_t value = std::stoi(line.substr(3));
            aim -= value;
        }
    }

    return horizontal * vertical;
}

#include "lib.h"

int32_t day02::run_part1(const std::vector<std::string> &input)
{
    int32_t horizontal = 0;
    int32_t vertical = 0;

    for (const auto &line : input)
    {
        if (line.find("forward ") == 0)
        {
            int32_t value = std::stoi(line.substr(8));
            horizontal += value;
        }
        else if (line.find("down ") == 0)
        {
            int32_t value = std::stoi(line.substr(5));
            vertical += value;
        }
        else if (line.find("up ") == 0)
        {
            int32_t value = std::stoi(line.substr(3));
            vertical -= value;
        }
    }

    return horizontal * vertical;
}

int32_t day02::run_part2(const std::vector<std::string> &input)
{
    int32_t horizontal = 0;
    int32_t vertical = 0;
    int32_t aim = 0;

    for (const auto &line : input)
    {
        if (line.find("forward ") == 0)
        {
            int32_t value = std::stoi(line.substr(8));
            horizontal += value;
            vertical += value * aim;
        }
        else if (line.find("down ") == 0)
        {
            int32_t value = std::stoi(line.substr(5));
            aim += value;
        }
        else if (line.find("up ") == 0)
        {
            int32_t value = std::stoi(line.substr(3));
            aim -= value;
        }
    }

    return horizontal * vertical;
}

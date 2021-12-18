#include "lib.h"

#include <regex>

namespace
{
    struct target_area
    {
        std::pair<int64_t, int64_t> x;
        std::pair<int64_t, int64_t> y;
    };

    target_area parse_input(const std::string &input)
    {
        std::regex re(R"(^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$)");
        std::smatch sm;
        if (std::regex_match(input, sm, re))
        {
            return {
                std::make_pair(std::stoll(sm[1].str()), std::stoll(sm[2].str())),
                std::make_pair(std::stoll(sm[3].str()), std::stoll(sm[4].str())),
            };
        }

        throw std::exception("Invalid input");
    }
}

int64_t day17::run_part1(const std::string &input)
{
    auto area = parse_input(input);

    int64_t y_max = 0;

    for (int64_t y = 0; y <= 1000; ++y)
    {
        int64_t current_position = 0;
        int64_t current_velocity = y;
        int64_t current_max = 0;

        while (current_velocity >= 0 || current_position >= area.y.first)
        {
            current_position += current_velocity;
            current_max = std::max(current_max, current_position);
            current_velocity -= 1;

            if (current_position >= area.y.first && current_position <= area.y.second)
            {
                y_max = std::max(y_max, current_max);
                break;
            }
        }
    }

    return y_max;
}

int64_t day17::run_part2(const std::string &input)
{
    auto area = parse_input(input);

    int64_t count = 0;

    // Calculate the minimum x velocity using n * (n + 1) / 2 = x.first, solve
    // using quadratic formula.
    double min_x = (std::sqrt((area.x.first * 8) + 1) - 1) / 2;

    for (int64_t x = static_cast<int64_t>(min_x + 1); x <= area.x.second; ++x)
    {
        for (int64_t y = area.y.first; y <= 1000; ++y)
        {
            int64_t pos_x = 0;
            int64_t pos_y = 0;
            int64_t vel_x = x;
            int64_t vel_y = y;

            while (pos_y >= area.y.first)
            {
                pos_x += vel_x;
                if (vel_x > 0)
                {
                    --vel_x;
                }

                pos_y += vel_y;
                --vel_y;

                if (pos_x >= area.x.first && pos_x <= area.x.second &&
                    pos_y >= area.y.first && pos_y <= area.y.second)
                {
                    ++count;
                    break;
                }
            }
        }
    }

    return count;
}

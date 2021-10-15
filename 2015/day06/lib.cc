#include "lib.h"

#include <common/point2d.h>
#include <common/string_convert.h>

#include <regex>
#include <unordered_map>

namespace
{
    using instruction = std::tuple<std::string, common::Point2D, common::Point2D>;

    std::vector<instruction> parse_instructions(const std::vector<std::string> &input)
    {
        std::vector<instruction> parsed_instructions;
        std::regex dimensions_regex("^(.+) (\\d+),(\\d+) through (\\d+),(\\d+)$");

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, dimensions_regex))
            {
                parsed_instructions.emplace_back(std::make_tuple(
                    sm[1].str(),
                    common::Point2D{common::string_to_int(sm[2].str()), common::string_to_int(sm[3].str())},
                    common::Point2D{common::string_to_int(sm[4].str()), common::string_to_int(sm[5].str())}));
            }
        }

        return parsed_instructions;
    }

    void process_instructions_part1(
        std::unordered_map<common::Point2D, bool> &grid,
        const std::vector<instruction> &input)
    {
        for (const auto &current : input)
        {
            std::string command;
            common::Point2D start, end;
            std::tie(command, start, end) = current;

            for (int i = start.x; i <= end.x; i++)
            {
                for (int j = start.y; j <= end.y; j++)
                {
                    auto current_position = common::Point2D{i, j};

                    if (command == "turn on")
                    {
                        grid[current_position] = true;
                    }
                    else if (command == "turn off")
                    {
                        grid[current_position] = false;
                    }
                    else if (command == "toggle")
                    {
                        bool current_value = false;

                        auto search = grid.find(current_position);
                        if (search != grid.end())
                        {
                            current_value = search->second;
                        }

                        grid[current_position] = !current_value;
                    }
                }
            }
        }
    }

    void process_instructions_part2(
        std::unordered_map<common::Point2D, int64_t> &grid,
        const std::vector<instruction> &input)
    {
        for (const auto &current : input)
        {
            std::string command;
            common::Point2D start, end;
            std::tie(command, start, end) = current;

            for (int i = start.x; i <= end.x; i++)
            {
                for (int j = start.y; j <= end.y; j++)
                {
                    auto current_position = common::Point2D{i, j};
                    int64_t current_value = 0;

                    auto search = grid.find(current_position);
                    if (search != grid.end())
                    {
                        current_value = search->second;
                    }

                    if (command == "turn on")
                    {
                        grid[current_position] = current_value + 1;
                    }
                    else if (command == "turn off")
                    {
                        if (current_value > 0)
                        {
                            grid[current_position] = current_value - 1;
                        }
                    }
                    else if (command == "toggle")
                    {
                        grid[current_position] = current_value + 2;
                    }
                }
            }
        }
    }
}

size_t day06::count_lit_lights(const std::vector<std::string> &input)
{
    std::unordered_map<common::Point2D, bool> grid;
    auto parsed_instructions = parse_instructions(input);
    process_instructions_part1(grid, parsed_instructions);

    size_t count = 0;
    for (const auto &cell : grid)
    {
        if (cell.second)
        {
            count++;
        }
    }

    return count;
}

int64_t day06::total_brightness(const std::vector<std::string> &input)
{
    std::unordered_map<common::Point2D, int64_t> grid;
    auto parsed_instructions = parse_instructions(input);
    process_instructions_part2(grid, parsed_instructions);

    size_t total = 0;
    for (const auto &cell : grid)
    {
        total += cell.second;
    }

    return total;
}
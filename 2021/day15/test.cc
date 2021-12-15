#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "1163751742",
        "1381373672",
        "2136511328",
        "3694931569",
        "7463417111",
        "1319128137",
        "1359912421",
        "3125421639",
        "1293138521",
        "2311944581",
    };
}

TEST(Day15, Part1)
{
    EXPECT_EQ(day15::run_part1(test_data), 40);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day15::run_part1(lines), 503);
}

TEST(Day15, Part2)
{
    EXPECT_EQ(day15::run_part2(test_data), 315);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day15::run_part2(lines), 2853);
}
#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "start-A",
        "start-b",
        "A-c",
        "A-b",
        "b-d",
        "A-end",
        "b-end",
    };
}

TEST(Day12, Part1)
{
    EXPECT_EQ(day12::run_part1(test_data), 10);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day12::run_part1(lines), 3887);
}

TEST(Day12, Part2)
{
    EXPECT_EQ(day12::run_part2(test_data), 36);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day12::run_part2(lines), 104834);
}
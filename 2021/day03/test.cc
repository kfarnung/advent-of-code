#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "00100",
        "11110",
        "10110",
        "10111",
        "10101",
        "01111",
        "00111",
        "11100",
        "10000",
        "11001",
        "00010",
        "01010",
    };
}

TEST(Day03, Part1)
{
    EXPECT_EQ(day03::run_part1(test_data), 198);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day03::run_part1(lines), 2724524);
}

TEST(Day03, Part2)
{
    EXPECT_EQ(day03::run_part2(test_data), 230);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day03::run_part2(lines), 2775870);
}
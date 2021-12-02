#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "forward 5",
        "down 5",
        "forward 8",
        "up 3",
        "down 8",
        "forward 2",
    };
}

TEST(Day02, Part1)
{
    EXPECT_EQ(day02::run_part1(test_data), 150);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day02::run_part1(lines), 1924923);
}

TEST(Day02, Part2)
{
    EXPECT_EQ(day02::run_part2(test_data), 900);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day02::run_part2(lines), 1982495697);
}
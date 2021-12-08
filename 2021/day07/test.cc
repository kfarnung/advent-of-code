#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "16,1,2,0,4,2,7,1,2,14",
    };
}

TEST(Day07, Part1)
{
    EXPECT_EQ(day07::run_part1(test_data), 37ul);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day07::run_part1(lines), 336120ul);
}

TEST(Day07, Part2)
{
    EXPECT_EQ(day07::run_part2(test_data), 168ul);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day07::run_part2(lines), 96864235ul);
}
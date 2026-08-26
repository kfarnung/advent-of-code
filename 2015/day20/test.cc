#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day20, Part1)
{
    EXPECT_EQ(day20::find_first_house("150"), 8);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day20::find_first_house(lines[0]), 776160);
}

TEST(Day20, Part2)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day20::find_first_house_part2(lines[0]), 786240);
}
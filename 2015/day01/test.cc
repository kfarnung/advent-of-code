#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day01, Part1)
{
    EXPECT_EQ(day01::find_floor("(())"), 0);
    EXPECT_EQ(day01::find_floor("()()"), 0);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day01::find_floor(lines[0]), 74);
}

TEST(Day01, Part2)
{
    EXPECT_EQ(day01::find_basement(")"), 1);
    EXPECT_EQ(day01::find_basement("()())"), 5);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day01::find_basement(lines[0]), 1795);
}
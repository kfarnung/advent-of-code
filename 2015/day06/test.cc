#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day06, Part1)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day06::count_lit_lights(lines), 569999);
}

TEST(Day06, Part2)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day06::total_brightness(lines), 17836115);
}
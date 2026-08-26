#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day17, Part1)
{
    EXPECT_EQ(day17::run_part1("target area: x=20..30, y=-10..-5"), 45);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day17::run_part1(lines[0]), 6903);
}

TEST(Day17, Part2)
{
    EXPECT_EQ(day17::run_part2("target area: x=20..30, y=-10..-5"), 112);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day17::run_part2(lines[0]), 2351);
}
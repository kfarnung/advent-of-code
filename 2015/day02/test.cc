#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day02, Part1)
{
    EXPECT_EQ(day02::calculate_wrapping_paper("2x3x4"), 58);
    EXPECT_EQ(day02::calculate_wrapping_paper("1x1x10"), 43);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day02::calculate_wrapping_paper(lines), 1606483);
}

TEST(Day02, Part2)
{
    EXPECT_EQ(day02::calculate_ribbon("2x3x4"), 34);
    EXPECT_EQ(day02::calculate_ribbon("1x1x10"), 14);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day02::calculate_ribbon(lines), 3842356);
}
#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "2199943210",
        "3987894921",
        "9856789892",
        "8767896789",
        "9899965678",
    };
}

TEST(Day09, Part1)
{
    EXPECT_EQ(day09::run_part1(test_data), 15ul);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day09::run_part1(lines), 585ul);
}

TEST(Day09, Part2)
{
    EXPECT_EQ(day09::run_part2(test_data), 1134ul);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day09::run_part2(lines), 827904ul);
}
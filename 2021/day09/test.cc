#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
std::vector<std::string> test_data()
{
    return {
        "2199943210", "3987894921", "9856789892", "8767896789", "9899965678",
    };
}
} // namespace

TEST(Day09, Part1)
{
    EXPECT_EQ(day09::run_part1(test_data()), 15);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day09::run_part1(lines), 585);
}

TEST(Day09, Part2)
{
    EXPECT_EQ(day09::run_part2(test_data()), 1134);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day09::run_part2(lines), 827904);
}
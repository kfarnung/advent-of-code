#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "5483143223",
        "2745854711",
        "5264556173",
        "6141336146",
        "6357385478",
        "4167524645",
        "2176841721",
        "6882881134",
        "4846848554",
        "5283751526",
    };
}

TEST(Day11, Part1)
{
    EXPECT_EQ(day11::run_part1(test_data), 1656);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day11::run_part1(lines), 1729);
}

TEST(Day11, Part2)
{
    EXPECT_EQ(day11::run_part2(test_data), 195);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day11::run_part2(lines), 237);
}
#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    ".#.#.#",
    "...##.",
    "#....#",
    "..#...",
    "#.#..#",
    "####..",
};

TEST(Day18, Part1)
{
    EXPECT_EQ(day18::run_iterations(test_data, 4, false), 4);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day18::run_iterations(lines, 100, false), 1061);
}

TEST(Day18, Part2)
{
    EXPECT_EQ(day18::run_iterations(test_data, 5, true), 17);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day18::run_iterations(lines, 100, true), 1006);
}
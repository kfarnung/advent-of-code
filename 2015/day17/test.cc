#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    "20",
    "15",
    "10",
    "5",
    "5",
};

TEST(Day17, Part1)
{
    EXPECT_EQ(day17::count_combinations(test_data, 25, false), 4);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day17::count_combinations(lines, 150, false), 4372);
}

TEST(Day17, Part2)
{
    EXPECT_EQ(day17::count_combinations(test_data, 25, true), 3);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day17::count_combinations(lines, 150, true), 4);
}
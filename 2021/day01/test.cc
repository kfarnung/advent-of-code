#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "199",
        "200",
        "208",
        "210",
        "200",
        "207",
        "240",
        "269",
        "260",
        "263",
    };
}

TEST(Day01, Part1)
{
    EXPECT_EQ(day01::count_increases(test_data, 1), 7ul);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day01::count_increases(lines, 1), 1502ul);
}

TEST(Day01, Part2)
{
    EXPECT_EQ(day01::count_increases(test_data, 3), 5ul);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day01::count_increases(lines, 3), 1538ul);
}
#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "0,9 -> 5,9",
        "8,0 -> 0,8",
        "9,4 -> 3,4",
        "2,2 -> 2,1",
        "7,0 -> 7,4",
        "6,4 -> 2,0",
        "0,9 -> 2,9",
        "3,4 -> 1,4",
        "0,0 -> 8,8",
        "5,5 -> 8,2",
    };
}

TEST(Day05, Part1)
{
    EXPECT_EQ(day05::count_overlaps(test_data, false), 5);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day05::count_overlaps(lines, false), 5835);
}

TEST(Day05, Part2)
{
    EXPECT_EQ(day05::count_overlaps(test_data, true), 12);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day05::count_overlaps(lines, true), 17013);
}
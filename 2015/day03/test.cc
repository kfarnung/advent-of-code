#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day03, Part1)
{
    EXPECT_EQ(day03::count_visited_houses(">", 1), 2);
    EXPECT_EQ(day03::count_visited_houses("^>v<", 1), 4);
    EXPECT_EQ(day03::count_visited_houses("^v^v^v^v^v", 1), 2);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day03::count_visited_houses(lines[0], 1), 2565);
}

TEST(Day03, Part2)
{
    EXPECT_EQ(day03::count_visited_houses("^v", 2), 3);
    EXPECT_EQ(day03::count_visited_houses("^>v<", 2), 3);
    EXPECT_EQ(day03::count_visited_houses("^v^v^v^v^v", 2), 11);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day03::count_visited_houses(lines[0], 2), 2639);
}
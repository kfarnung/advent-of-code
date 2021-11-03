#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
    "Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
};

TEST(Day15, Part1)
{
    EXPECT_EQ(day15::find_best_cookie(test_data, -1), 62842880);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day15::find_best_cookie(lines, -1), 13882464);
}

TEST(Day15, Part2)
{
    EXPECT_EQ(day15::find_best_cookie(test_data, 500), 57600000);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day15::find_best_cookie(lines, 500), 11171160);
}
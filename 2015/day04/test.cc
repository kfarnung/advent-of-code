#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day04, Part1)
{
    EXPECT_EQ(day04::mine_adventcoin("abcdef", 5), 609043);
    EXPECT_EQ(day04::mine_adventcoin("pqrstuv", 5), 1048970);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day04::mine_adventcoin(lines[0], 5), 254575);
}

TEST(Day04, Part2)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day04::mine_adventcoin(lines[0], 6), 1038736);
}
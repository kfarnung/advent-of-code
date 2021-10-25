#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day10, Part1)
{
    EXPECT_EQ(day10::look_and_say("1", 1), "11");
    EXPECT_EQ(day10::look_and_say("11", 1), "21");
    EXPECT_EQ(day10::look_and_say("21", 1), "1211");
    EXPECT_EQ(day10::look_and_say("1211", 1), "111221");
    EXPECT_EQ(day10::look_and_say("111221", 1), "312211");
    EXPECT_EQ(day10::look_and_say("1", 5), "312211");

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day10::look_and_say(lines[0], 40).size(), 252594);
}

TEST(Day10, Part2)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day10::look_and_say(lines[0], 50).size(), 3579328);
}
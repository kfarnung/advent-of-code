#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day05, Part1)
{
    EXPECT_EQ(day05::is_nice_string_part1("ugknbfddgicrmopn"), true);
    EXPECT_EQ(day05::is_nice_string_part1("aaa"), true);
    EXPECT_EQ(day05::is_nice_string_part1("jchzalrnumimnmhp"), false);
    EXPECT_EQ(day05::is_nice_string_part1("haegwjzuvuyypxyu"), false);
    EXPECT_EQ(day05::is_nice_string_part1("dvszwmarrgswjxmb"), false);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day05::count_nice_strings_part1(lines), 258);
}

TEST(Day05, Part2)
{
    EXPECT_EQ(day05::is_nice_string_part2("qjhvhtzxzqqjkmpb"), true);
    EXPECT_EQ(day05::is_nice_string_part2("xxyxx"), true);
    EXPECT_EQ(day05::is_nice_string_part2("uurcxstgmygtbstg"), false);
    EXPECT_EQ(day05::is_nice_string_part2("ieodomkazucvgmuy"), false);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day05::count_nice_strings_part2(lines), 53);
}
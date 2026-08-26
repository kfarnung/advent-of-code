#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    R"("")",
    R"("abc")",
    R"("aaa\"aaa")",
    R"("\x27")",
};

TEST(Day08, Part1)
{
    EXPECT_EQ(day08::calculate_string_overhead(test_data), 12);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day08::calculate_string_overhead(lines), 1342);
}

TEST(Day08, Part2)
{
    EXPECT_EQ(day08::calculate_encoding_overhead(test_data), 19);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day08::calculate_encoding_overhead(lines), 2074);
}
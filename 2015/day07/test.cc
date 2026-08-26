#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    "123 -> x",
    "456 -> y",
    "x AND y -> d",
    "x OR y -> e",
    "x LSHIFT 2 -> f",
    "y RSHIFT 2 -> g",
    "NOT x -> h",
    "NOT y -> i",
};

TEST(Day07, Part1)
{
    EXPECT_EQ(day07::get_wire_output(test_data, "d"), 72);
    EXPECT_EQ(day07::get_wire_output(test_data, "e"), 507);
    EXPECT_EQ(day07::get_wire_output(test_data, "f"), 492);
    EXPECT_EQ(day07::get_wire_output(test_data, "g"), 114);
    EXPECT_EQ(day07::get_wire_output(test_data, "h"), 65412);
    EXPECT_EQ(day07::get_wire_output(test_data, "i"), 65079);
    EXPECT_EQ(day07::get_wire_output(test_data, "x"), 123);
    EXPECT_EQ(day07::get_wire_output(test_data, "y"), 456);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day07::get_wire_output(lines, "a"), 956);
}

TEST(Day07, Part2)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day07::get_wire_output_part2(lines, "a"), 40149);
}
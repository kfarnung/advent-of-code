#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    "London to Dublin = 464",
    "London to Belfast = 518",
    "Dublin to Belfast = 141",
};

TEST(Day09, Part1)
{
    EXPECT_EQ(day09::calculate_shortest_distance(test_data), 605);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day09::calculate_shortest_distance(lines), 251);
}

TEST(Day09, Part2)
{
    EXPECT_EQ(day09::calculate_longest_distance(test_data), 982);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day09::calculate_longest_distance(lines), 898);
}
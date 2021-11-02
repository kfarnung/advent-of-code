#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
    "Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
};

TEST(Day14, Part1)
{
    EXPECT_EQ(day14::find_longest_distance(test_data, 1000), 1120);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day14::find_longest_distance(lines, 2503), 2640);
}

TEST(Day14, Part2)
{
    EXPECT_EQ(day14::find_part2_winner(test_data, 1000), 689);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day14::find_part2_winner(lines, 2503), 1102);
}
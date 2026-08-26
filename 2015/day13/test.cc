#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

const std::vector<std::string> test_data{
    "Alice would gain 54 happiness units by sitting next to Bob.",
    "Alice would lose 79 happiness units by sitting next to Carol.",
    "Alice would lose 2 happiness units by sitting next to David.",
    "Bob would gain 83 happiness units by sitting next to Alice.",
    "Bob would lose 7 happiness units by sitting next to Carol.",
    "Bob would lose 63 happiness units by sitting next to David.",
    "Carol would lose 62 happiness units by sitting next to Alice.",
    "Carol would gain 60 happiness units by sitting next to Bob.",
    "Carol would gain 55 happiness units by sitting next to David.",
    "David would gain 46 happiness units by sitting next to Alice.",
    "David would lose 7 happiness units by sitting next to Bob.",
    "David would gain 41 happiness units by sitting next to Carol.",
};

TEST(Day13, Part1)
{
    EXPECT_EQ(day13::find_highest_happiness(test_data, false), 330);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day13::find_highest_happiness(lines, false), 709);
}

TEST(Day13, Part2)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day13::find_highest_happiness(lines, true), 668);
}
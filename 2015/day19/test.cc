#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

std::vector<std::string> test_data()
{
    return {
        "H => HO", "H => OH", "O => HH", "", "HOH",
    };
}

std::vector<std::string> test_data2()
{
    return {
        "e => H", "e => O", "H => HO", "H => OH", "O => HH", "", "HOH",
    };
}

std::vector<std::string> test_data3()
{
    return {
        "e => H", "e => O", "H => HO", "H => OH", "O => HH", "", "HOHOHO",
    };
}

TEST(Day19, Part1)
{
    EXPECT_EQ(day19::count_molecules(test_data()), 4);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day19::count_molecules(lines), 576);
}

TEST(Day19, Part2)
{
    EXPECT_EQ(day19::find_minimum_replacements(test_data2()), 3);
    EXPECT_EQ(day19::find_minimum_replacements(test_data3()), 6);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day19::find_minimum_replacements(lines), 207);
}
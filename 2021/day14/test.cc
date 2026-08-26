#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "NNCB",
        "",
        "CH -> B",
        "HH -> N",
        "CB -> H",
        "NH -> C",
        "HB -> C",
        "HC -> B",
        "HN -> C",
        "NN -> C",
        "BH -> H",
        "NC -> B",
        "NB -> B",
        "BN -> B",
        "BB -> N",
        "BC -> B",
        "CC -> N",
        "CN -> C",
    };
}

TEST(Day14, Part1)
{
    EXPECT_EQ(day14::run_part1(test_data), 1588);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day14::run_part1(lines), 3342);
}

TEST(Day14, Part2)
{
    EXPECT_EQ(day14::run_part2(test_data), 2188189693529);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day14::run_part2(lines), 3776553567525);
}
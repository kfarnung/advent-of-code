#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "3,4,3,1,2",
    };
}

TEST(Day06, Part1)
{
    EXPECT_EQ(day06::simulate_lanternfish(test_data, 80), 5934ull);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day06::simulate_lanternfish(lines, 80), 362346ull);
}

TEST(Day06, Part2)
{
    EXPECT_EQ(day06::simulate_lanternfish(test_data, 256), 26984457539ull);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day06::simulate_lanternfish(lines, 256), 1639643057051ull);
}
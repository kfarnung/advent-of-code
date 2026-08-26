#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

TEST(Day12, Part1)
{
    EXPECT_EQ(day12::sum_all_numbers("[1,2,3]", false), 6);
    EXPECT_EQ(day12::sum_all_numbers("{\"a\":2,\"b\":4}", false), 6);
    EXPECT_EQ(day12::sum_all_numbers("[[[3]]]", false), 3);
    EXPECT_EQ(day12::sum_all_numbers("{\"a\":{\"b\":4},\"c\":-1}", false), 3);
    EXPECT_EQ(day12::sum_all_numbers("{\"a\":[-1,1]}", false), 0);
    EXPECT_EQ(day12::sum_all_numbers("[-1,{\"a\":1}]", false), 0);
    EXPECT_EQ(day12::sum_all_numbers("[]", false), 0);
    EXPECT_EQ(day12::sum_all_numbers("{}", false), 0);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day12::sum_all_numbers(lines[0], false), 156366);
}

TEST(Day12, Part2)
{
    EXPECT_EQ(day12::sum_all_numbers("[1,2,3]", true), 6);
    EXPECT_EQ(day12::sum_all_numbers("[1,{\"c\":\"red\",\"b\":2},3]", true), 4);
    EXPECT_EQ(day12::sum_all_numbers("{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", true), 0);
    EXPECT_EQ(day12::sum_all_numbers("[1,\"red\",5]", true), 6);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day12::sum_all_numbers(lines[0], true), 96852);
}
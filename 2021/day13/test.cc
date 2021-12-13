#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "6,10",
        "0,14",
        "9,10",
        "0,3",
        "10,4",
        "4,11",
        "6,0",
        "6,12",
        "4,1",
        "0,13",
        "10,12",
        "3,4",
        "3,0",
        "8,4",
        "1,10",
        "2,14",
        "8,10",
        "9,0",
        "",
        "fold along y=7",
        "fold along x=5",
    };
}

TEST(Day13, Part1)
{
    EXPECT_EQ(day13::run_part1(test_data), 17);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day13::run_part1(lines), 850);
}

TEST(Day13, Part2)
{
    EXPECT_EQ(
        day13::run_part2(test_data),
        "#####\n"
        "#   #\n"
        "#   #\n"
        "#   #\n"
        "#####");

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(
        day13::run_part2(lines),
        " ##  #  #  ##   ##  ###   ##   ##  #  #\n"
        "#  # #  # #  # #  # #  # #  # #  # #  #\n"
        "#  # #### #    #    #  # #    #  # #  #\n"
        "#### #  # # ## #    ###  # ## #### #  #\n"
        "#  # #  # #  # #  # #    #  # #  # #  #\n"
        "#  # #  #  ###  ##  #     ### #  #  ## ");
}
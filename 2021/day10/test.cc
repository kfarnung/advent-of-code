#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{
        "[({(<(())[]>[[{[]{<()<>>",
        "[(()[<>])]({[<{<<[]>>(",
        "{([(<{}[<>[]}>{[]{[(<()>",
        "(((({<>}<{<{<>}{[]{[]{}",
        "[[<[([]))<([[{}[[()]]]",
        "[{[{({}]{}}([{[{{{}}([]",
        "{<[[]]>}<{[{[{[]{()[[[]",
        "[<(<(<(<{}))><([]([]()",
        "<{([([[(<>()){}]>(<<{{",
        "<{([{{}}[<[[[<>{}]]]>[]]",
    };
}

TEST(Day10, Part1)
{
    EXPECT_EQ(day10::run_part1(test_data), 26397);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day10::run_part1(lines), 341823);
}

TEST(Day10, Part2)
{
    EXPECT_EQ(day10::run_part2(test_data), 288957);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day10::run_part2(lines), 2801302861);
}
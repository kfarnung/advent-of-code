#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    std::vector<char> string_to_vector(const std::string &input)
    {
        return std::vector<char>{begin(input), end(input)};
    }

    std::string vector_to_string(const std::vector<char> &input)
    {
        return std::string{begin(input), end(input)};
    }
}

TEST(Day11, Part1)
{
    auto test_vec = string_to_vector("z");
    day11::increment_password(test_vec);
    EXPECT_EQ(vector_to_string(test_vec), "a");

    test_vec = string_to_vector("az");
    day11::increment_password(test_vec);
    EXPECT_EQ(vector_to_string(test_vec), "ba");

    test_vec = string_to_vector("zzzzzzzz");
    day11::increment_password(test_vec);
    EXPECT_EQ(vector_to_string(test_vec), "aaaaaaaa");

    EXPECT_EQ(day11::validate_password(string_to_vector("hijklmmn")), false);
    EXPECT_EQ(day11::validate_password(string_to_vector("abbceffg")), false);
    EXPECT_EQ(day11::validate_password(string_to_vector("abbcegjk")), false);

    EXPECT_EQ(day11::next_valid_password("abcdefgh"), "abcdffaa");
    EXPECT_EQ(day11::next_valid_password("ghijklmn"), "ghjaabcc");

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day11::next_valid_password(lines[0]), "cqjxxyzz");
}

TEST(Day11, Part2)
{
    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    auto first = day11::next_valid_password(lines[0]);
    EXPECT_EQ(day11::next_valid_password(first), "cqkaabcc");
}
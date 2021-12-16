#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <gtest/gtest.h>

namespace
{
    const std::vector<std::string> test_data{};
}

TEST(Day16, Part1)
{
    EXPECT_EQ(day16::run_part1("8A004A801A8002F478"), 16);
    EXPECT_EQ(day16::run_part1("620080001611562C8802118E34"), 12);
    EXPECT_EQ(day16::run_part1("C0015000016115A2E0802F182340"), 23);
    EXPECT_EQ(day16::run_part1("A0016C880162017C3686B18A3D4780"), 31);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day16::run_part1(lines[0]), 1012);
}

TEST(Day16, Part2)
{
    EXPECT_EQ(day16::run_part2("C200B40A82"), 3);
    EXPECT_EQ(day16::run_part2("04005AC33890"), 54);
    EXPECT_EQ(day16::run_part2("880086C3E88112"), 7);
    EXPECT_EQ(day16::run_part2("CE00C43D881120"), 9);
    EXPECT_EQ(day16::run_part2("D8005AC2A8F0"), 1);
    EXPECT_EQ(day16::run_part2("F600BC2D8F"), 0);
    EXPECT_EQ(day16::run_part2("9C005AC2F8F0"), 0);
    EXPECT_EQ(day16::run_part2("9C0141080250320F1802104A08"), 1);

    std::ifstream file("input.txt");
    auto lines = common::splitlines(file);
    EXPECT_EQ(day16::run_part2(lines[0]), 2223947372407);
}
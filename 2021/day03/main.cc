#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day03::run_part1(lines) << std::endl;
    std::cout << "Part 2: " << day03::run_part2(lines) << std::endl;
}
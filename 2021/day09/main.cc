#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    if (argc != 2)
    {
        std::cerr << "Input file required" << "\n";
        return 1;
    }

    std::ifstream file(argv[1]);
    if (!file.is_open())
    {
        std::cerr << "Could not open file: " << argv[1] << "\n";
        return 1;
    }

    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day09::run_part1(lines) << "\n";
    std::cout << "Part 2: " << day09::run_part2(lines) << "\n";
}
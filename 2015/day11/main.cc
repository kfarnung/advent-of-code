#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    auto part1 = day11::next_valid_password(lines[0]);
    std::cout << "Part 1: " << part1 << std::endl;
    std::cout << "Part 2: " << day11::next_valid_password(part1) << std::endl;
}
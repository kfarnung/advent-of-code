#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day21::find_first_house(lines[0]) << std::endl;
    std::cout << "Part 2: " << day21::find_first_house_part2(lines[0]) << std::endl;
}
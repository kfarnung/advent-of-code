#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day14::find_longest_distance(lines, 2503) << std::endl;
    std::cout << "Part 2: " << day14::find_part2_winner(lines, 2503) << std::endl;
}
#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day09::calculate_shortest_distance(lines) << std::endl;
    std::cout << "Part 2: " << day09::calculate_longest_distance(lines) << std::endl;
}
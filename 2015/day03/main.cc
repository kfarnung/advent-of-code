#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day03::count_visited_houses(lines[0], 1) << std::endl;
    std::cout << "Part 2: " << day03::count_visited_houses(lines[0], 2) << std::endl;
}
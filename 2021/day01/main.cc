#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day01::count_increases(lines, 1) << std::endl;
    std::cout << "Part 2: " << day01::count_increases(lines, 3) << std::endl;
}
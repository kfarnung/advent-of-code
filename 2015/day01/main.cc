#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day01::find_floor(lines[0]) << std::endl;
    std::cout << "Part 2: " << day01::find_basement(lines[0]) << std::endl;
}
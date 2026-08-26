#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day19::count_molecules(lines) << std::endl;
    std::cout << "Part 2: " << day19::find_minimum_replacements(lines) << std::endl;
}
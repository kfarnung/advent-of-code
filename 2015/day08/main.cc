#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day08::calculate_string_overhead(lines) << std::endl;
    std::cout << "Part 2: " << day08::calculate_encoding_overhead(lines) << std::endl;
}
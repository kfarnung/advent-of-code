#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day06::count_lit_lights(lines) << std::endl;
    std::cout << "Part 2: " << day06::total_brightness(lines) << std::endl;
}
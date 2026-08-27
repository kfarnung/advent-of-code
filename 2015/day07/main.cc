#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day07::get_wire_output(lines, "a") << "\n";
    std::cout << "Part 2: " << day07::get_wire_output_part2(lines, "a") << "\n";
}
#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day04::mine_adventcoin(lines[0], 5) << std::endl;
    std::cout << "Part 2: " << day04::mine_adventcoin(lines[0], 6) << std::endl;
}
#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day15::find_best_cookie(lines, -1) << std::endl;
    std::cout << "Part 2: " << day15::find_best_cookie(lines, 500) << std::endl;
}
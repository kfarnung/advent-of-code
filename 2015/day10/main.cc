#include "lib.h"

#include <common/input_parser.h>

#include <fstream>
#include <iostream>

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    auto lines = common::splitlines(file);

    std::cout << "Part 1: " << day10::look_and_say(lines[0], 40).size() << std::endl;
    std::cout << "Part 2: " << day10::look_and_say(lines[0], 50).size() << std::endl;
}
#include <common/input_parser.h>

#include <istream>
#include <sstream>

std::vector<std::string> common::splitlines(std::istream &input)
{
    std::vector<std::string> lines;
    for (std::string line; std::getline(input, line);)
    {
        lines.push_back(line);
    }

    return lines;
}

std::vector<std::string> common::splitlines(const char *input)
{
    std::stringstream ss(input);
    return splitlines(ss);
}

std::vector<std::string> common::splitstr(const std::string &input, char ch)
{
    std::vector<std::string> tokens;
    std::istringstream ss(input);

    std::string current;
    while (std::getline(ss, current, ch))
    {
        tokens.emplace_back(std::move(current));
    }

    return tokens;
}
#include "lib.h"

#include <common/string_convert.h>

#include <regex>

namespace
{
    struct Dimensions
    {
        int length;
        int width;
        int height;
    };

    Dimensions parse_dimensions(const std::string &input)
    {
        std::regex dimensions_regex("^(\\d+)x(\\d+)x(\\d+)$");
        std::smatch sm;
        if (std::regex_match(input, sm, dimensions_regex))
        {
            return Dimensions{
                common::string_to_int(sm[1].str()),
                common::string_to_int(sm[2].str()),
                common::string_to_int(sm[3].str())};
        }

        return Dimensions{0, 0, 0};
    }
}

int64_t day02::calculate_wrapping_paper(const std::string &input)
{
    auto dimensions = parse_dimensions(input);
    int min_side = 0;

    int side1 = dimensions.length * dimensions.width;
    min_side = side1;

    int side2 = dimensions.width * dimensions.height;
    min_side = std::min(min_side, side2);

    int side3 = dimensions.height * dimensions.length;
    min_side = std::min(min_side, side3);

    return 2 * (side1 + side2 + side3) + min_side;
}

int64_t day02::calculate_wrapping_paper(const std::vector<std::string> &input)
{
    int64_t total = 0;

    for (const auto &line : input)
    {
        total += calculate_wrapping_paper(line);
    }

    return total;
}

int64_t day02::calculate_ribbon(const std::string &input)
{
    auto dimensions = parse_dimensions(input);
    int min_side = 2 * dimensions.length + 2 * dimensions.width;
    min_side = std::min(min_side, 2 * dimensions.width + 2 * dimensions.height);
    min_side = std::min(min_side, 2 * dimensions.height + 2 * dimensions.length);

    return min_side + (dimensions.length * dimensions.width * dimensions.height);
}

int64_t day02::calculate_ribbon(const std::vector<std::string> &input)
{
    int64_t total = 0;

    for (const auto &line : input)
    {
        total += calculate_ribbon(line);
    }

    return total;
}
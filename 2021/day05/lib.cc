#include "lib.h"

#include <common/point2d.h>

#include <unordered_map>
#include <regex>

namespace
{
    using LineSegment = std::pair<common::Point2D, common::Point2D>;

    std::vector<LineSegment> parse_lines(const std::vector<std::string> &input)
    {
        std::vector<LineSegment> segments;
        std::regex re(R"(^(\d+),(\d+) -> (\d+),(\d+)$)");

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, re))
            {
                segments.emplace_back(std::make_pair(
                    common::Point2D{
                        std::stoi(sm[1].str()),
                        std::stoi(sm[2].str()),
                    },
                    common::Point2D{
                        std::stoi(sm[3].str()),
                        std::stoi(sm[4].str()),
                    }));
            }
        }

        return segments;
    }

    common::Point2D find_slope(common::Point2D p1, common::Point2D p2)
    {
        int32_t x = p2.x - p1.x;
        int32_t y = p2.y - p1.y;
        int32_t divisor = std::max(std::abs(x), std::abs(y));

        return common::Point2D{x / divisor, y / divisor};
    }
}

size_t day05::count_overlaps(const std::vector<std::string> &input, bool include_diagonals)
{
    auto segments = parse_lines(input);
    std::unordered_map<common::Point2D, size_t> hits;

    for (const auto &segment : segments)
    {
        auto slope = find_slope(segment.first, segment.second);
        if (include_diagonals || slope.x == 0 || slope.y == 0)
        {
            for (auto i = segment.first; i != segment.second; i += slope)
            {
                hits[i] += 1;
            }

            hits[segment.second] += 1;
        }
    }

    size_t count = 0;
    for (const auto &pair : hits)
    {
        if (pair.second > 1)
        {
            ++count;
        }
    }

    return count;
}

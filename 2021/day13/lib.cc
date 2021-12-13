#include "lib.h"

#include <common/input_parser.h>
#include <common/point2d.h>

#include <limits>
#include <unordered_set>
#include <tuple>

namespace
{
    using grid_t = std::unordered_set<common::Point2D>;
    using foldlist_t = std::vector<common::Point2D>;

    std::tuple<grid_t, foldlist_t> parse_input(const std::vector<std::string> &input)
    {
        grid_t grid;
        foldlist_t foldlist;

        for (const auto &line : input)
        {
            if (line.empty())
            {
                continue;
            }

            if (line.find("fold") == 0)
            {
                auto parts = common::splitstr(line, ' ');
                auto fold = common::splitstr(parts[2], '=');
                auto value = std::stoll(fold[1]);

                if (fold[0] == "x")
                {
                    foldlist.emplace_back(value, 0);
                }
                else
                {
                    foldlist.emplace_back(0, value);
                }
            }
            else
            {
                auto parts = common::splitstr(line, ',');
                grid.emplace(std::stoll(parts[0]), std::stoll(parts[1]));
            }
        }

        return std::tuple<grid_t, foldlist_t>{grid, foldlist};
    }

    void process_fold(grid_t &grid, const common::Point2D &fold)
    {
        grid_t new_grid;

        for (const auto &point : grid)
        {
            if (fold.x > 0 && point.x > fold.x)
            {
                auto delta = point.x - fold.x;
                new_grid.emplace(fold.x - delta, point.y);
            }
            else if (fold.y > 0 && point.y > fold.y)
            {
                auto delta = point.y - fold.y;
                new_grid.emplace(point.x, fold.y - delta);
            }
            else
            {
                new_grid.emplace(point.x, point.y);
            }
        }

        std::swap(grid, new_grid);
    }

    std::string print_grid(const grid_t &grid)
    {
        int64_t min_x = std::numeric_limits<int64_t>::max();
        int64_t max_x = std::numeric_limits<int64_t>::min();
        int64_t min_y = std::numeric_limits<int64_t>::max();
        int64_t max_y = std::numeric_limits<int64_t>::min();

        for (const auto &point : grid)
        {
            min_x = std::min(min_x, point.x);
            max_x = std::max(max_x, point.x);
            min_y = std::min(min_y, point.y);
            max_y = std::max(max_y, point.y);
        }

        std::vector<char> rendered;

        for (int64_t y = min_y; y <= max_y; ++y)
        {
            for (int64_t x = min_x; x <= max_x; ++x)
            {
                auto result = grid.find(common::Point2D{x, y});
                if (result != grid.end())
                {
                    rendered.emplace_back('#');
                }
                else
                {
                    rendered.emplace_back(' ');
                }
            }

            rendered.emplace_back('\n');
        }

        rendered.pop_back();
        return std::string(begin(rendered), end(rendered));
    }
}

int64_t day13::run_part1(const std::vector<std::string> &input)
{
    grid_t grid;
    foldlist_t foldlist;
    std::tie(grid, foldlist) = parse_input(input);
    process_fold(grid, foldlist[0]);

    return grid.size();
}

std::string day13::run_part2(const std::vector<std::string> &input)
{
    grid_t grid;
    foldlist_t foldlist;
    std::tie(grid, foldlist) = parse_input(input);

    for (const auto &fold : foldlist)
    {
        process_fold(grid, fold);
    }

    return print_grid(grid);
}

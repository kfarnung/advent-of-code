#include "lib.h"

#include <algorithm>
#include <deque>
#include <limits>
#include <numeric>
#include <set>

namespace
{
    std::vector<std::vector<uint8_t>> parse_grid(const std::vector<std::string> &input)
    {
        std::vector<std::vector<uint8_t>> grid;
        for (const auto &line : input)
        {
            std::vector<uint8_t> row;
            for (const auto &ch : line)
            {
                row.emplace_back(static_cast<uint8_t>(ch - '0'));
            }

            grid.emplace_back(std::move(row));
        }

        return grid;
    }

    std::vector<std::pair<size_t, size_t>> find_low_points(const std::vector<std::vector<uint8_t>> &grid)
    {
        std::vector<std::pair<size_t, size_t>> low_points;

        for (size_t i = 0; i < grid.size(); ++i)
        {
            for (size_t j = 0; j < grid[i].size(); ++j)
            {
                uint8_t min_adjacent = std::numeric_limits<uint8_t>::max();

                if (i > 0)
                {
                    min_adjacent = std::min(min_adjacent, grid[i - 1][j]);
                }

                if (j > 0)
                {
                    min_adjacent = std::min(min_adjacent, grid[i][j - 1]);
                }

                if (i < grid.size() - 1)
                {
                    min_adjacent = std::min(min_adjacent, grid[i + 1][j]);
                }

                if (j < grid[i].size() - 1)
                {
                    min_adjacent = std::min(min_adjacent, grid[i][j + 1]);
                }

                if (grid[i][j] < min_adjacent)
                {
                    low_points.emplace_back(i, j);
                }
            }
        }

        return low_points;
    }
}

uint32_t day09::run_part1(const std::vector<std::string> &input)
{
    auto grid = parse_grid(input);
    auto low_points = find_low_points(grid);

    uint32_t risk_levels = 0;
    for (const auto &point : low_points)
    {
        risk_levels += grid[point.first][point.second] + 1;
    }

    return risk_levels;
}

uint32_t day09::run_part2(const std::vector<std::string> &input)
{
    auto grid = parse_grid(input);
    auto low_points = find_low_points(grid);

    std::vector<uint32_t> basin_sizes;

    for (const auto &point : low_points)
    {
        std::set<std::pair<size_t, size_t>> visited;
        std::deque<std::pair<size_t, size_t>> queue;

        uint32_t basin_size = 0;
        queue.emplace_back(point);

        while (!queue.empty())
        {
            auto current = queue.front();
            queue.pop_front();

            if (visited.find(current) != visited.end())
            {
                continue;
            }

            visited.emplace(current);
            basin_size += 1;

            if (current.first > 0 &&
                grid[current.first - 1][current.second] < 9 &&
                visited.find(std::make_pair(current.first - 1, current.second)) == visited.end())
            {
                queue.emplace_back(current.first - 1, current.second);
            }

            if (current.second > 0 &&
                grid[current.first][current.second - 1] < 9 &&
                visited.find(std::make_pair(current.first, current.second - 1)) == visited.end())
            {
                queue.emplace_back(current.first, current.second - 1);
            }

            if (current.first < grid.size() - 1 &&
                grid[current.first + 1][current.second] < 9 &&
                visited.find(std::make_pair(current.first + 1, current.second)) == visited.end())
            {
                queue.emplace_back(current.first + 1, current.second);
            }

            if (current.second < grid[current.first].size() - 1 &&
                grid[current.first][current.second + 1] < 9 &&
                visited.find(std::make_pair(current.first, current.second + 1)) == visited.end())
            {
                queue.emplace_back(current.first, current.second + 1);
            }
        }

        basin_sizes.emplace_back(basin_size);
    }

    std::sort(begin(basin_sizes), end(basin_sizes));
    
    uint32_t basins_score = 1;
    for (size_t i = 0; i < 3; ++i)
    {
        basins_score *= basin_sizes[basin_sizes.size() - 1 - i];
    }

    return basins_score;
}

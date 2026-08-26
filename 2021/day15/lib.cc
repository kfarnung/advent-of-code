#include "lib.h"

#include <common/point2d.h>

#include <array>
#include <cstdint>
#include <queue>
#include <map>

namespace
{
    using grid_t = std::vector<std::vector<int8_t>>;
    using queue_item_t = std::pair<int64_t, common::Point2D>;

    const std::vector<common::Point2D> directions{
        {0, 1},
        {1, 0},
        {0, -1},
        {-1, 0},
    };

    grid_t parse_grid(const std::vector<std::string> &input)
    {
        grid_t grid;

        for (const auto &line : input)
        {
            std::vector<int8_t> row;

            for (const auto &ch : line)
            {
                row.emplace_back(static_cast<uint8_t>(ch - '0'));
            }

            grid.emplace_back(row);
        }

        return grid;
    }

    int64_t calculate_risk_level(const grid_t &grid, size_t boards)
    {
        auto size_x = static_cast<int64_t>(grid.size());
        auto size_y = static_cast<int64_t>(grid[0].size());
        int64_t bound_x = size_x * boards;
        int64_t bound_y = size_y * boards;
        common::Point2D dest{bound_x - 1, bound_y - 1};

        std::priority_queue<queue_item_t, std::vector<queue_item_t>, std::greater<queue_item_t>> queue;
        queue.emplace(0, common::Point2D(0, 0));

        std::map<common::Point2D, int64_t> visited;

        while (!queue.empty())
        {
            auto current = queue.top();
            queue.pop();

            if (current.second == dest)
            {
                return current.first;
            }

            visited.emplace(current.second, current.first);

            for (const auto &direction : directions)
            {
                auto x = current.second.x + direction.x;
                auto y = current.second.y + direction.y;

                if (x >= 0 && y >= 0 && x < bound_x && y < bound_y)
                {
                    common::Point2D next{x, y};
                    auto risk_level = (grid[x % size_x][y % size_y] - 1 + x / size_x + y / size_y) % 9 + 1;

                    auto best = visited.find(next);
                    if (best == visited.end() || risk_level < best->second)
                    {
                        visited[next] = risk_level;
                        queue.emplace(current.first + risk_level, next);
                    }
                }
            }
        }

        return 0;
    }
}

int64_t day15::run_part1(const std::vector<std::string> &input)
{
    auto grid = parse_grid(input);
    return calculate_risk_level(grid, 1);
}

int64_t day15::run_part2(const std::vector<std::string> &input)
{
    auto grid = parse_grid(input);
    return calculate_risk_level(grid, 5);
}

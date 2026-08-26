#include "lib.h"

#include <cstdint>
#include <deque>

namespace
{
    using Grid = std::vector<std::vector<uint8_t>>;

    const std::vector<std::pair<int32_t, int32_t>> directions{
        {-1, -1},
        {-1, 0},
        {-1, 1},
        {0, -1},
        {0, 1},
        {1, -1},
        {1, 0},
        {1, 1},
    };

    Grid parse_grid(const std::vector<std::string> &input)
    {
        Grid grid;

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

    int64_t count_grid_cells(const Grid &grid)
    {
        size_t count = 0;
        for (const auto &row : grid)
        {
            count += row.size();
        }

        return count;
    }

    int64_t do_step(Grid &grid)
    {
        std::deque<std::pair<int64_t, int64_t>> queue;

        for (size_t i = 0; i < grid.size(); ++i)
        {
            for (size_t j = 0; j < grid[i].size(); ++j)
            {
                grid[i][j] += 1;
                if (grid[i][j] > 9)
                {
                    queue.emplace_back(i, j);
                }
            }
        }

        while (!queue.empty())
        {
            auto current = queue.front();
            queue.pop_front();

            for (const auto &direction : directions)
            {
                auto i = current.first + direction.first;
                auto j = current.second + direction.second;

                if (i >= 0 && i < static_cast<int64_t>(grid.size()) &&
                    j >= 0 && j < static_cast<int64_t>(grid[i].size()) &&
                    grid[i][j] <= 9)
                {
                    grid[i][j] += 1;
                    if (grid[i][j] > 9)
                    {
                        queue.emplace_back(i, j);
                    }
                }
            }
        }

        int64_t flash_count = 0;
        for (size_t i = 0; i < grid.size(); ++i)
        {
            for (size_t j = 0; j < grid[i].size(); ++j)
            {
                if (grid[i][j] > 9)
                {
                    flash_count += 1;
                    grid[i][j] = 0;
                }
            }
        }

        return flash_count;
    }
}

int64_t day11::run_part1(const std::vector<std::string> &input)
{
    auto grid = parse_grid(input);
    int64_t flash_count = 0;

    for (size_t i = 0; i < 100; ++i)
    {
        flash_count += do_step(grid);
    }

    return flash_count;
}

int64_t day11::run_part2(const std::vector<std::string> &input)
{
    auto grid = parse_grid(input);
    auto cell_count = count_grid_cells(grid);

    for (int64_t step_count = 1;; ++step_count)
    {
        if (do_step(grid) == cell_count)
        {
            return step_count;
        }
    }
}

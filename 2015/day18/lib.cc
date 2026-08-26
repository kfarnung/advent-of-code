#include "lib.h"

#include <common/point2d.h>

#include <unordered_map>

namespace
{
    struct Grid
    {
        std::unordered_map<common::Point2D, bool> cells;
        int32_t height;
        int32_t width;
    };

    Grid parse_grid(const std::vector<std::string> &input)
    {
        Grid grid;

        int32_t row = 0;
        for (const auto &line : input)
        {
            int32_t col = 0;
            for (const auto &ch : line)
            {
                grid.cells.insert({common::Point2D{row, col}, ch == '#'});
                col++;
            }

            grid.width = col;
            row++;
        }

        grid.height = row;
        return grid;
    }

    size_t count_lit_neighbors(const Grid &grid, const common::Point2D &point)
    {
        size_t lit_count = 0;
        for (int32_t i = -1; i <= 1; i++)
        {
            for (int32_t j = -1; j <= 1; j++)
            {
                if (i == 0 && j == 0)
                {
                    continue;
                }

                auto result = grid.cells.find(common::Point2D{point.x + i, point.y + j});
                if (result != grid.cells.end() && result->second)
                {
                    lit_count++;
                }
            }
        }

        return lit_count;
    }

    size_t count_lit(const Grid &grid)
    {
        size_t count = 0;
        for (const auto &entry : grid.cells)
        {
            if (entry.second)
            {
                count++;
            }
        }

        return count;
    }

    void turn_on_corners(Grid &grid)
    {
        grid.cells[common::Point2D{0, 0}] = true;
        grid.cells[common::Point2D{0, grid.width - 1}] = true;
        grid.cells[common::Point2D{grid.height - 1, 0}] = true;
        grid.cells[common::Point2D{grid.height - 1, grid.width - 1}] = true;
    }
}

size_t day18::run_iterations(const std::vector<std::string> &input, size_t iteration_count, bool corners_stuck)
{
    Grid current_grid = parse_grid(input);
    Grid next_grid{{}, current_grid.height, current_grid.width};

    for (size_t i = 0; i < iteration_count; i++)
    {
        if (corners_stuck)
        {
            turn_on_corners(current_grid);
        }

        for (const auto &entry : current_grid.cells)
        {
            size_t lit_neighbors = count_lit_neighbors(current_grid, entry.first);
            next_grid.cells[entry.first] = lit_neighbors == 3 || (entry.second && lit_neighbors == 2);
        }

        std::swap(current_grid, next_grid);
    }

    if (corners_stuck)
    {
        turn_on_corners(current_grid);
    }

    return count_lit(current_grid);
}
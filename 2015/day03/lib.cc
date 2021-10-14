#include "lib.h"

#include <common/point2d.h>

#include <unordered_set>
#include <vector>

int64_t day03::count_visited_houses(const std::string &input, size_t actor_count)
{
    std::unordered_set<common::Point2D> visited{};
    std::vector<common::Point2D> actor_positions{actor_count, common::Point2D{0, 0}};

    visited.insert(common::Point2D{0, 0});
    size_t step_count = 0;

    for (const char &ch : input)
    {
        auto actor_index = step_count % actor_count;
        auto current_position = actor_positions[actor_index];

        switch (ch)
        {
        case '^':
            current_position = current_position + common::Point2D{0, 1};
            break;

        case 'v':
            current_position = current_position + common::Point2D{0, -1};
            break;

        case '>':
            current_position = current_position + common::Point2D{1, 0};
            break;

        case '<':
            current_position = current_position + common::Point2D{-1, 0};
            break;

        default:
            return -1;
        }

        visited.insert(current_position);
        actor_positions[actor_index] = current_position;
        step_count++;
    }

    return visited.size();
}
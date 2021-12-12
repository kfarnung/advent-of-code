#include "lib.h"

#include <common/input_parser.h>

#include <deque>
#include <unordered_map>
#include <unordered_set>

namespace
{
    using Graph = std::unordered_map<std::string, std::vector<std::string>>;

    struct GraphState
    {
        const std::string current;
        const std::unordered_set<std::string> visited;
        const bool visited_small_twice;
    };

    Graph parse_graph(const std::vector<std::string> &input)
    {
        Graph graph;

        for (const auto &line : input)
        {
            auto parts = common::splitstr(line, '-');
            graph[parts[0]].push_back(parts[1]);
            graph[parts[1]].push_back(parts[0]);
        }

        return graph;
    }

    int64_t count_paths(const Graph &graph, bool allow_second_visit)
    {
        int64_t path_count = 0;

        std::deque<GraphState> queue;
        queue.emplace_back(GraphState{"start", std::unordered_set<std::string>{}, false});

        while (!queue.empty())
        {
            auto current = queue.front();
            queue.pop_front();

            if (current.current == "end")
            {
                path_count += 1;
                continue;
            }

            std::unordered_set<std::string> next_visited(current.visited);
            next_visited.insert(current.current);

            for (const auto &neighbor : graph.at(current.current))
            {
                if (neighbor == "start")
                {
                    continue;
                }

                auto is_small_cave = !std::isupper(neighbor[0]);
                auto already_visited = current.visited.find(neighbor) != current.visited.end();
                if (is_small_cave &&
                    already_visited &&
                    (!allow_second_visit || current.visited_small_twice))
                {
                    continue;
                }

                queue.emplace_back(
                    GraphState{
                        neighbor,
                        next_visited,
                        current.visited_small_twice || (is_small_cave && already_visited),
                    });
            }
        }

        return path_count;
    }
}

int64_t day12::run_part1(const std::vector<std::string> &input)
{
    auto graph = parse_graph(input);
    return count_paths(graph, false);
}

int64_t day12::run_part2(const std::vector<std::string> &input)
{
    auto graph = parse_graph(input);
    return count_paths(graph, true);
}

#include "lib.h"

#include <common/input_parser.h>

#include <deque>
#include <map>
#include <set>

namespace
{
    using Graph = std::map<std::string, std::vector<std::string>>;

    Graph parse_graph(const std::vector<std::string> &input)
    {
        Graph graph;

        for (const auto &line : input)
        {
            auto parts = common::splitstr(line, '-');

            if (parts[1] != "start" && parts[0] != "end")
            {
                graph[parts[0]].push_back(parts[1]);
            }

            if (parts[0] != "start" && parts[1] != "end")
            {
                graph[parts[1]].push_back(parts[0]);
            }
        }

        return graph;
    }

    int64_t count_paths(
        const Graph &graph,
        const bool allow_second_visit,
        const std::string &current,
        const std::set<std::string> &visited,
        const bool used_second_visit)
    {
        if (current == "end")
        {
            return 1;
        }

        std::set<std::string> next_visited(visited);

        if (!std::isupper(current[0]))
        {
            next_visited.insert(current);
        }

        int64_t path_count = 0;

        for (const auto &neighbor : graph.at(current))
        {
            auto already_visited = visited.find(neighbor) != visited.end();
            if (already_visited && (!allow_second_visit || used_second_visit))
            {
                continue;
            }

            path_count += count_paths(
                graph,
                allow_second_visit,
                neighbor,
                next_visited,
                used_second_visit || already_visited);
        }

        return path_count;
    }
}

int64_t day12::run_part1(const std::vector<std::string> &input)
{
    auto graph = parse_graph(input);
    return count_paths(graph, false, "start", {}, false);
}

int64_t day12::run_part2(const std::vector<std::string> &input)
{
    auto graph = parse_graph(input);
    return count_paths(graph, true, "start", {}, false);
}

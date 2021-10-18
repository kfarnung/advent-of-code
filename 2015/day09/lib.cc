#include "lib.h"

#include <common/string_convert.h>

#include <limits>
#include <regex>
#include <unordered_map>
#include <unordered_set>

namespace
{
    using GraphEntry = std::unordered_map<std::string, int>;
    using Graph = std::unordered_map<std::string, GraphEntry>;

    struct Distance
    {
        std::string start;
        std::string end;
        int distance;
    };

    Distance parse_line(const std::string &input)
    {
        std::regex re("^(.+) to (.+) = (\\d+)");

        std::smatch sm;
        if (std::regex_match(input, sm, re))
        {
            return Distance{
                sm[1].str(),
                sm[2].str(),
                common::string_to_int(sm[3].str()),
            };
        }

        return Distance{"", "", 0};
    }

    void add_graph_edge(Graph &graph, std::string start, std::string end, int distance)
    {
        auto search = graph.find(start);
        if (search == graph.end())
        {
            graph.emplace(start, GraphEntry{});
        }

        graph[start][end] = distance;
    }

    Graph create_graph(const std::vector<std::string> &input)
    {
        Graph graph;

        for (const std::string &line : input)
        {
            auto distance = parse_line(line);
            add_graph_edge(graph, distance.start, distance.end, distance.distance);
            add_graph_edge(graph, distance.end, distance.start, distance.distance);
        }

        return graph;
    }

    int find_shortest_distance(const Graph &graph, const std::string &location)
    {
        std::unordered_set<std::string> visited_locations;
        int current_distance = 0;
        std::string current_location = location;

        while (true)
        {
            int min_distance = std::numeric_limits<int>::max();
            std::string min_location;
            
            visited_locations.insert(current_location);

            for (const auto &path : graph.at(current_location))
            {
                auto search = visited_locations.find(path.first);
                if (search != visited_locations.end())
                {
                    continue;
                }

                if (path.second < min_distance)
                {
                    min_distance = path.second;
                    min_location = path.first;
                }
            }

            if (min_distance == std::numeric_limits<int>::max())
            {
                break;
            }

            current_distance += min_distance;
            current_location = min_location;
        }

        return current_distance;
    }

    int find_longest_distance(const Graph &graph, const std::string &location)
    {
        std::unordered_set<std::string> visited_locations;
        int current_distance = 0;
        std::string current_location = location;

        while (true)
        {
            int max_distance = std::numeric_limits<int>::min();
            std::string max_location;
            
            visited_locations.insert(current_location);

            for (const auto &path : graph.at(current_location))
            {
                auto search = visited_locations.find(path.first);
                if (search != visited_locations.end())
                {
                    continue;
                }

                if (path.second > max_distance)
                {
                    max_distance = path.second;
                    max_location = path.first;
                }
            }

            if (max_distance == std::numeric_limits<int>::min())
            {
                break;
            }

            current_distance += max_distance;
            current_location = max_location;
        }

        return current_distance;
    }
}

size_t day09::calculate_shortest_distance(const std::vector<std::string> &input)
{
    auto graph = create_graph(input);

    int min_distance = std::numeric_limits<int>::max();
    for (const auto &element : graph)
    {
        min_distance = std::min(min_distance, find_shortest_distance(graph, element.first));
    }

    return min_distance;
}

size_t day09::calculate_longest_distance(const std::vector<std::string> &input)
{
    auto graph = create_graph(input);

    int max_distance = std::numeric_limits<int>::min();
    for (const auto &element : graph)
    {
        max_distance = std::max(max_distance, find_longest_distance(graph, element.first));
    }

    return max_distance;
}
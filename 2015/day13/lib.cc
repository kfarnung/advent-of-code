#include "lib.h"

#include <common/string_convert.h>

#include <limits>
#include <regex>
#include <unordered_map>
#include <unordered_set>

namespace
{
    using HappinessMap = std::unordered_map<std::string, std::unordered_map<std::string, int32_t>>;

    HappinessMap build_map(const std::vector<std::string> &input, bool include_me)
    {
        HappinessMap map;
        std::regex re(R"(^(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).$)");

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, re))
            {
                int value = common::string_to_int(sm[3].str());
                if (sm[2].str() == "lose")
                {
                    value *= -1;
                }

                // Collect the total value for both parties.
                map[sm[1].str()][sm[4].str()] += value;
                map[sm[4].str()][sm[1].str()] += value;

                if (include_me)
                {
                    map[sm[1].str()]["Me"] = 0;
                    map["Me"][sm[1].str()] = 0;
                }
            }
        }

        return map;
    }

    int64_t happiness_recurse(
        HappinessMap &map,
        std::unordered_set<std::string> visited,
        std::string first_name,
        std::string current_name,
        int64_t current_total)
    {
        bool found = false;
        int64_t max_value = std::numeric_limits<int64_t>::min();

        visited.insert(current_name);

        for (const auto &entry : map[current_name])
        {
            auto search = visited.find(entry.first);
            if (search != visited.end())
            {
                continue;
            }

            found = true;
            std::unordered_set<std::string> next_visited{visited};
            max_value = std::max(
                max_value,
                happiness_recurse(map, next_visited, first_name, entry.first, current_total + entry.second));
        }

        if (!found)
        {
            return current_total + map[current_name][first_name];
        }

        return max_value;
    }
}

int64_t day13::find_highest_happiness(const std::vector<std::string> &input, bool include_me)
{
    auto map = build_map(input, include_me);
    std::unordered_set<std::string> visited;

    auto first = begin(map)->first;
    return happiness_recurse(map, visited, first, first, 0);
}
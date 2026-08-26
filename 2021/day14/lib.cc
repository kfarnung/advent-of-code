#include "lib.h"

#include <algorithm>
#include <list>
#include <map>
#include <tuple>

namespace
{
    using pair_insertions_t = std::map<std::pair<char, char>, char>;
    std::tuple<std::string, pair_insertions_t> parse_input(const std::vector<std::string> &input)
    {
        std::string polymer_template;
        pair_insertions_t pair_insertions;

        for (const auto &line : input)
        {
            if (line.empty())
            {
                continue;
            }

            auto location = line.find("->");
            if (location != std::string::npos)
            {
                pair_insertions.emplace(std::make_pair(line[0], line[1]), line[line.size() - 1]);
            }
            else
            {
                polymer_template = line;
            }
        }

        return std::tuple<std::string, pair_insertions_t>{
            polymer_template,
            pair_insertions,
        };
    }
}

int64_t day14::run_part1(const std::vector<std::string> &input)
{
    std::string polymer_template;
    pair_insertions_t pair_insertions;
    std::tie(polymer_template, pair_insertions) = parse_input(input);

    std::list<char> current_polymer(begin(polymer_template), end(polymer_template));

    for (size_t i = 0; i < 10; ++i)
    {
        char previous = '\0';
        for (auto it = current_polymer.begin(); it != current_polymer.end(); ++it)
        {
            char current = *it;
            auto lookup = pair_insertions.find(std::make_pair(previous, current));
            if (lookup != pair_insertions.end())
            {
                current_polymer.insert(it, lookup->second);
            }

            previous = current;
        }
    }

    std::map<char, int64_t> counts;
    for (const auto &ch : current_polymer)
    {
        counts[ch] += 1;
    }

    auto smallest = std::min_element(
        begin(counts), end(counts),
        [](const std::pair<char, int64_t> &a, const std::pair<char, int64_t> &b)
        {
            return a.second < b.second;
        });

    auto largest = std::max_element(
        begin(counts), end(counts),
        [](const std::pair<char, int64_t> &a, const std::pair<char, int64_t> &b)
        {
            return a.second < b.second;
        });

    return largest->second - smallest->second;
}

int64_t day14::run_part2(const std::vector<std::string> &input)
{
    std::string polymer_template;
    pair_insertions_t pair_insertions;
    std::tie(polymer_template, pair_insertions) = parse_input(input);

    std::map<char, int64_t> char_counts;
    std::map<std::pair<char, char>, int64_t> pair_counts;

    char previous = '\0';
    for (const auto &ch : polymer_template)
    {
        if (previous != '\0')
        {
            pair_counts[std::make_pair(previous, ch)] += 1;
        }

        char_counts[ch] += 1;
        previous = ch;
    }

    for (size_t i = 0; i < 40; ++i)
    {
        auto pair_counts_copy = pair_counts;
        for (const auto &entry : pair_counts_copy)
        {
            auto result = pair_insertions.find(entry.first);
            if (result != pair_insertions.end())
            {
                char_counts[result->second] += entry.second;

                pair_counts[entry.first] -= entry.second;
                pair_counts[std::make_pair(entry.first.first, result->second)] += entry.second;
                pair_counts[std::make_pair(result->second, entry.first.second)] += entry.second;
            }
        }
    }

    auto smallest = std::min_element(
        begin(char_counts), end(char_counts),
        [](const std::pair<char, int64_t> &a, const std::pair<char, int64_t> &b)
        {
            return a.second < b.second;
        });

    auto largest = std::max_element(
        begin(char_counts), end(char_counts),
        [](const std::pair<char, int64_t> &a, const std::pair<char, int64_t> &b)
        {
            return a.second < b.second;
        });

    return largest->second - smallest->second;
}

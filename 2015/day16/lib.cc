#include "lib.h"

#include <common/string_convert.h>

#include <limits>
#include <regex>
#include <unordered_map>

namespace
{
    const std::unordered_map<std::string, int32_t> analysis_results{
        {"children", 3},
        {"cats", 7},
        {"samoyeds", 2},
        {"pomeranians", 3},
        {"akitas", 0},
        {"vizslas", 0},
        {"goldfish", 5},
        {"trees", 3},
        {"cars", 2},
        {"perfumes", 1},
    };

    struct AuntSue
    {
        int32_t number;
        std::unordered_map<std::string, int32_t> traits;
    };

    std::vector<AuntSue> parse_list(const std::vector<std::string> &input)
    {
        std::regex re(R"(^Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)$)");
        std::vector<AuntSue> result;

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, re))
            {
                auto sue = AuntSue{common::string_to_int(sm[1].str())};
                sue.traits.insert({sm[2].str(), common::string_to_int(sm[3].str())});
                sue.traits.insert({sm[4].str(), common::string_to_int(sm[5].str())});
                sue.traits.insert({sm[6].str(), common::string_to_int(sm[7].str())});

                result.push_back(std::move(sue));
            }
        }

        return result;
    }

    bool is_match(const AuntSue &sue, bool use_ranges)
    {
        for (const auto &trait : sue.traits)
        {
            auto result = analysis_results.at(trait.first);
            if (use_ranges && (trait.first == "cats" || trait.first == "trees"))
            {
                if (trait.second <= result)
                {
                    return false;
                }
            }
            else if (use_ranges && (trait.first == "pomeranians" || trait.first == "goldfish"))
            {
                if (trait.second >= result)
                {
                    return false;
                }
            }
            else if (result != trait.second)
            {
                return false;
            }
        }

        return true;
    }
}

int32_t day16::find_aunt_sue(const std::vector<std::string> &input, bool use_ranges)
{
    auto sue_list = parse_list(input);
    for (const auto &sue : sue_list)
    {
        if (is_match(sue, use_ranges))
        {
            return sue.number;
        }
    }

    return 0;
}
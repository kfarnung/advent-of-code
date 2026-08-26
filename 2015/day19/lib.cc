#include "lib.h"

#include <regex>
#include <tuple>
#include <unordered_set>

namespace
{
    using Replacement = std::pair<std::string, std::string>;

    std::tuple<std::vector<Replacement>, std::string> parse_inputs(const std::vector<std::string> &input)
    {
        std::regex re(R"(^(\w+) => (\w+)$)");
        std::vector<Replacement> replacements;
        std::string molecule;

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, re))
            {
                replacements.emplace_back(sm[1].str(), sm[2].str());
            }
            else if (!line.empty())
            {
                molecule = line;
            }
        }

        std::sort(begin(replacements), end(replacements),
                  [](Replacement i, Replacement j)
                  { return i.second.size() > j.second.size(); });

        return std::make_tuple(replacements, molecule);
    }

    std::vector<std::tuple<size_t, size_t, std::string>> find_inverse_replacements(
        const std::string &current,
        const std::vector<Replacement> &replacements)
    {
        std::vector<std::tuple<size_t, size_t, std::string>> valid_replacements;

        for (const auto &replacement : replacements)
        {
            size_t found = current.rfind(replacement.second);
            if (found != std::string::npos)
            {
                valid_replacements.emplace_back(found, replacement.second.size(), replacement.first);
            }
        }

        // Sort so that the latest (and then longest) replacement is first.
        std::sort(
            begin(valid_replacements),
            end(valid_replacements),
            std::greater<std::tuple<size_t, size_t, std::string>>());

        return valid_replacements;
    }
}

size_t day19::count_molecules(const std::vector<std::string> &input)
{
    std::unordered_set<std::string> found_molecules;
    std::vector<Replacement> replacements;
    std::string starting_molecule;

    std::tie(replacements, starting_molecule) = parse_inputs(input);

    for (const auto &replacement : replacements)
    {
        size_t found = 0;

        while ((found = starting_molecule.find(replacement.first, found)) != std::string::npos)
        {
            found_molecules.insert(
                starting_molecule.substr(0, found) +
                replacement.second +
                starting_molecule.substr(found + replacement.first.size()));

            found++;
        }
    }

    return found_molecules.size();
}

size_t day19::find_minimum_replacements(const std::vector<std::string> &input)
{
    std::vector<std::pair<std::string, size_t>> stack;
    std::unordered_set<std::string> seen;
    std::vector<Replacement> replacements;
    std::string target_molecule;

    std::tie(replacements, target_molecule) = parse_inputs(input);
    stack.emplace_back(target_molecule, 0);

    while (!stack.empty())
    {
        auto &current = stack.back();
        seen.insert(current.first);

        auto valid_replacements = find_inverse_replacements(current.first, replacements);
        if (current.second < valid_replacements.size())
        {
            const auto &replacement = valid_replacements[current.second];
            current.second += 1;

            std::string new_molecule{current.first};
            new_molecule.replace(
                std::get<0>(replacement),
                std::get<1>(replacement),
                std::get<2>(replacement));

            if (new_molecule == "e")
            {
                return stack.size();
            }

            auto result = seen.find(new_molecule);
            if (result == seen.end())
            {
                stack.emplace_back(std::move(new_molecule), 0);
            }
        }
        else
        {
            stack.pop_back();
        }
    }

    return 0;
}
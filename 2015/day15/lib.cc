#include "lib.h"

#include <common/string_convert.h>

#include <limits>
#include <regex>

namespace
{
    struct Ingredient
    {
        std::string name;
        int32_t capacity;
        int32_t durability;
        int32_t flavor;
        int32_t texture;
        int32_t calories;
    };

    std::vector<Ingredient> parse_list(const std::vector<std::string> &input)
    {
        std::regex re(R"(^(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$)");
        std::vector<Ingredient> result;

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, re))
            {
                result.emplace_back(Ingredient{
                    sm[1].str(),
                    common::string_to_int(sm[2].str()),
                    common::string_to_int(sm[3].str()),
                    common::string_to_int(sm[4].str()),
                    common::string_to_int(sm[5].str()),
                    common::string_to_int(sm[6].str()),
                });
            }
        }

        return result;
    }

    int32_t sum_without_last(const std::vector<int32_t> &input)
    {
        int32_t sum = 0;

        for (size_t i = 0; i < input.size() - 1; i++)
        {
            sum += input[i];
        }

        return sum;
    }

    int32_t increment_ingredients(std::vector<int32_t> &input)
    {
        for (size_t i = 0; i < input.size() - 1; i++)
        {
            size_t index = input.size() - 2 - i;
            input[index]++;

            int32_t sum = sum_without_last(input);
            if (sum <= 100)
            {
                return sum;
            }

            input[index] = 0;
        }

        return -1;
    }

    int64_t calculate_score(const std::vector<Ingredient> &input, const std::vector<int32_t> &counts)
    {
        int32_t capacity = 0;
        for (size_t i = 0; i < input.size(); i++)
        {
            capacity += input[i].capacity * counts[i];
        }

        int32_t durability = 0;
        for (size_t i = 0; i < input.size(); i++)
        {
            durability += input[i].durability * counts[i];
        }

        int32_t flavor = 0;
        for (size_t i = 0; i < input.size(); i++)
        {
            flavor += input[i].flavor * counts[i];
        }

        int32_t texture = 0;
        for (size_t i = 0; i < input.size(); i++)
        {
            texture += input[i].texture * counts[i];
        }

        return std::max(capacity, 0) * std::max(durability, 0) * std::max(flavor, 0) * std::max(texture, 0);
    }

    int32_t calculate_calories(const std::vector<Ingredient> &input, const std::vector<int32_t> &counts)
    {
        int32_t calories = 0;
        for (size_t i = 0; i < input.size(); i++)
        {
            calories += input[i].calories * counts[i];
        }

        return calories;
    }
}

int64_t day15::find_best_cookie(const std::vector<std::string> &input, int32_t calorie_target)
{
    int32_t total_teaspoons = 100;
    auto ingredient_list = parse_list(input);
    std::vector<int32_t> counts(ingredient_list.size());
    int64_t max_score = std::numeric_limits<int64_t>::min();

    int32_t sum = 0;
    while (sum >= 0)
    {
        counts[counts.size() - 1] = total_teaspoons - sum;

        if (calorie_target <= 0 || calorie_target == calculate_calories(ingredient_list, counts))
        {
            max_score = std::max(max_score, calculate_score(ingredient_list, counts));
        }

        sum = increment_ingredients(counts);
    }

    return max_score;
}
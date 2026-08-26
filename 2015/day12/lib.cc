#include "lib.h"

#include <nlohmann/json.hpp>

namespace
{
    int64_t walk_object(nlohmann::json &j, bool skip_red)
    {
        int64_t sum = 0;

        for (const auto &el : j.items())
        {
            auto &value = el.value();

            if (skip_red && j.is_object() && value.is_string() && value.get<std::string>() == "red")
            {
                return 0;
            }
            else if (value.is_number())
            {
                sum += value.get<int64_t>();
            }
            else if (value.is_array() || value.is_object())
            {
                sum += walk_object(value, skip_red);
            }
        }

        return sum;
    }
}

int64_t day12::sum_all_numbers(const std::string &input, bool skip_red)
{
    int64_t sum = 0;
    auto j = nlohmann::json::parse(input);

    return walk_object(j, skip_red);
}
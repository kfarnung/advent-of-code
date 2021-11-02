#include "lib.h"

#include <common/string_convert.h>

#include <limits>
#include <regex>
#include <unordered_map>

namespace
{
    struct Reindeer
    {
        std::string name;
        int32_t velocity;
        int32_t move_duration;
        int32_t rest_duration;
    };

    std::vector<Reindeer> parse_list(const std::vector<std::string> &input)
    {
        std::regex re(R"(^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.$)");
        std::vector<Reindeer> result;

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, re))
            {
                result.emplace_back(Reindeer{
                    sm[1].str(),
                    common::string_to_int(sm[2].str()),
                    common::string_to_int(sm[3].str()),
                    common::string_to_int(sm[4].str()),
                });
            }
        }

        return result;
    }

    int32_t calculate_distance(Reindeer reindeer, int32_t current_time)
    {
        int32_t total_time = reindeer.move_duration + reindeer.rest_duration;
        int32_t iteration_count = current_time / total_time;
        int32_t remainder = current_time % total_time;

        return (reindeer.velocity * reindeer.move_duration * iteration_count) +
               (reindeer.velocity * std::min(reindeer.move_duration, remainder));
    }

    std::vector<std::string> pick_winners(const std::vector<Reindeer> &input, int32_t current_time)
    {
        int32_t max_distance = std::numeric_limits<int32_t>::min();
        std::vector<std::string> max_reindeer;

        for (const auto &reindeer : input)
        {
            int32_t distance_travelled = calculate_distance(reindeer, current_time);
            if (distance_travelled > max_distance)
            {
                max_distance = distance_travelled;
                max_reindeer.clear();
            }

            if (distance_travelled == max_distance)
            {
                max_reindeer.emplace_back(reindeer.name);
            }
        }

        return max_reindeer;
    }
}

int32_t day14::find_longest_distance(const std::vector<std::string> &input, int32_t target_time)
{
    auto reindeer_list = parse_list(input);
    int32_t max_distance = std::numeric_limits<int32_t>::min();

    for (const auto &reindeer : reindeer_list)
    {
        int32_t distance_travelled = calculate_distance(reindeer, target_time);
        max_distance = std::max(max_distance, distance_travelled);
    }

    return max_distance;
}

int32_t day14::find_part2_winner(const std::vector<std::string> &input, int32_t target_time)
{
    auto reindeer_list = parse_list(input);
    std::unordered_map<std::string, int32_t> scores;

    for (int32_t i = 1; i <= target_time; i++)
    {
        for (const auto &winner : pick_winners(reindeer_list, i))
        {
            scores[winner] += 1;
        }
    }

    int32_t max_score = std::numeric_limits<int32_t>::min();
    for (const auto &entry : scores)
    {
        max_score = std::max(max_score, entry.second);
    }

    return max_score;
}
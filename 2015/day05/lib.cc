#include "lib.h"

#include <unordered_set>

bool day05::is_nice_string_part1(const std::string &input)
{
    size_t vowel_count = 0;
    bool has_repeat_char = false;

    char last = '\0';

    for (const auto &ch : input)
    {
        if ((last == 'a' && ch == 'b') ||
            (last == 'c' && ch == 'd') ||
            (last == 'p' && ch == 'q') ||
            (last == 'x' && ch == 'y'))
        {
            return false;
        }

        if (last == ch)
        {
            has_repeat_char = true;
        }

        if (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u')
        {
            vowel_count++;
        }

        last = ch;
    }

    return vowel_count >= 3 && has_repeat_char;
}

size_t day05::count_nice_strings_part1(const std::vector<std::string> &input)
{
    size_t count = 0;

    for (const auto &str : input)
    {
        if (is_nice_string_part1(str))
        {
            count++;
        }
    }

    return count;
}

bool day05::is_nice_string_part2(const std::string &input)
{
    std::unordered_set<std::string> pairs;
    bool has_repeated_pairs = false;
    bool has_separated_pair = false;

    char minus2 = '\0';
    char minus1 = '\0';

    for (const auto &ch : input)
    {
        if (minus2 == ch)
        {
            has_separated_pair = true;
        }

        if (minus1 != ch || minus2 != ch)
        {
            auto pair = std::to_string(minus1) + std::to_string(ch);
            if (pairs.find(pair) != pairs.end())
            {
                has_repeated_pairs = true;
            }

            pairs.insert(pair);
        }

        if (has_repeated_pairs && has_separated_pair)
        {
            return true;
        }

        minus2 = minus1;
        minus1 = ch;
    }

    return false;
}

size_t day05::count_nice_strings_part2(const std::vector<std::string> &input)
{
    size_t count = 0;

    for (const auto &str : input)
    {
        if (is_nice_string_part2(str))
        {
            count++;
        }
    }

    return count;
}
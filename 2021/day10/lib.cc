#include "lib.h"

#include <algorithm>

namespace
{
    struct parser_state
    {
        bool valid;
        char current_char;
        std::vector<char> remaining_stack;
    };

    parser_state parse_line(const std::string &line)
    {
        std::vector<char> stack;
        for (const auto &ch : line)
        {
            switch (ch)
            {
            case '(':
                stack.push_back(')');
                break;

            case '[':
                stack.push_back(']');
                break;

            case '{':
                stack.push_back('}');
                break;

            case '<':
                stack.push_back('>');
                break;

            default:
                if (stack.back() == ch)
                {
                    stack.pop_back();
                }
                else
                {
                    return parser_state{false, ch};
                }
            }
        }

        return parser_state{true, '\0', stack};
    }

    uint32_t score_illegal_character(char ch)
    {
        switch (ch)
        {
        case ')':
            return 3;

        case ']':
            return 57;

        case '}':
            return 1197;

        case '>':
            return 25137;
        }

        return 0;
    }

    uint32_t score_missing_character(char ch)
    {
        switch (ch)
        {
        case ')':
            return 1;

        case ']':
            return 2;

        case '}':
            return 3;

        case '>':
            return 4;
        }

        return 0;
    }
}

uint64_t day10::run_part1(const std::vector<std::string> &input)
{
    uint64_t total_score = 0;

    for (const auto &line : input)
    {
        auto parsed = parse_line(line);
        if (!parsed.valid)
        {
            total_score += score_illegal_character(parsed.current_char);
        }
    }

    return total_score;
}

uint64_t day10::run_part2(const std::vector<std::string> &input)
{
    std::vector<uint64_t> scores;

    for (const auto &line : input)
    {
        auto parsed = parse_line(line);
        if (parsed.valid)
        {
            uint64_t score = 0;
            while (!parsed.remaining_stack.empty())
            {
                score *= 5;
                score += score_missing_character(parsed.remaining_stack.back());
                parsed.remaining_stack.pop_back();
            }

            scores.push_back(score);
        }
    }

    std::sort(begin(scores), end(scores));
    return scores[scores.size() / 2];
}

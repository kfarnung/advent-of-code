#include "lib.h"

#include <vector>

std::string day10::look_and_say(const std::string &input, size_t iterations)
{
    std::string current{input};

    for (size_t i = 0; i < iterations; i++)
    {
        std::vector<char> next;
        char current_char = '\0';
        size_t current_count = 0;

        for (const char &ch : current)
        {
            if (ch != current_char)
            {
                if (current_count > 0)
                {
                    next.push_back(static_cast<char>('0' + current_count));
                    next.push_back(current_char);
                }

                current_char = ch;
                current_count = 0;
            }

            current_count++;
        }

        if (current_count > 0)
        {
            next.push_back(static_cast<char>('0' + current_count));
            next.push_back(current_char);
        }

        current = std::string{begin(next), end(next)};
    }

    return current;
}
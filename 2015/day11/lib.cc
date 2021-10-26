#include "lib.h"

void day11::increment_password(std::vector<char> &input)
{
    bool carry = true;

    for (size_t i = 0, size = input.size(); i < size && carry; i++)
    {
        size_t current_index = size - i - 1;
        char ch = input[current_index] + 1;

        if (ch == 'i' || ch == 'l' || ch == 'o')
        {
            // Skip invalid characters.
            ch++;
        }

        carry = ch > 'z';
        input[current_index] = carry ? 'a' : ch;
    }
}

bool day11::validate_password(const std::vector<char> &input)
{
    bool has_increasing_straight = false;
    bool has_different_pairs = false;
    char previous_pair = '\0';

    char minus2 = '\0';
    char minus1 = '\0';

    for (const auto &ch : input)
    {
        if (ch == 'i' || ch == 'l' || ch == 'o')
        {
            return false;
        }

        if (ch == (minus1 + 1) && minus1 == (minus2 + 1))
        {
            has_increasing_straight = true;
        }

        if (ch == minus1 && ch != minus2)
        {
            if (previous_pair != '\0' && previous_pair != ch)
            {
                has_different_pairs = true;
            }

            previous_pair = ch;
        }

        minus2 = minus1;
        minus1 = ch;
    }

    return has_increasing_straight && has_different_pairs;
}

std::string day11::next_valid_password(const std::string &input)
{
    std::vector<char> current{begin(input), end(input)};

    do
    {
        increment_password(current);
    } while (!validate_password(current));

    return std::string{begin(current), end(current)};
}
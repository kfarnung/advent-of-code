#include "lib.h"

int64_t day01::find_floor(const std::string &input)
{
    int64_t floor = 0;

    for (const auto ch : input)
    {
        switch (ch)
        {
        case '(':
            floor++;
            break;

        case ')':
            floor--;
            break;
        }
    }

    return floor;
}

int64_t day01::find_basement(const std::string &input)
{
    int64_t position = 1;
    int64_t floor = 0;

    for (const auto ch : input)
    {
        switch (ch)
        {
        case '(':
            floor++;
            break;

        case ')':
            floor--;
            break;
        }

        if (floor < 0)
        {
            return position;
        }

        position++;
    }

    return -1;
}
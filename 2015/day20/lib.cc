#include "lib.h"

#include <cmath>
#include <cstdint>
#include <limits>

namespace
{    
    uint32_t sum_of_factors(uint32_t num, uint32_t multiplier, uint32_t max_visits)
    {
        if (num == 1)
        {
            return multiplier;
        }

        uint32_t sum = num * multiplier;

        if (num <= max_visits)
        {
            sum += multiplier;
        }

        uint32_t range_end = static_cast<uint32_t>(std::sqrt(static_cast<double>(num)));

        for (uint32_t i = 2; i <= range_end; i++)
        {
            if (num % i == 0)
            {
                auto other = num / i;

                if (other <= max_visits)
                {
                    sum += i * multiplier;
                }
                
                if (other != i && i <= max_visits)
                {
                    sum += other * multiplier;
                }
            }
        }

        return sum;
    }
}

uint32_t day20::find_first_house(const std::string &input)
{
    uint32_t target = std::stoul(input);

    for (uint32_t i = 1;; i++)
    {
        if (sum_of_factors(i, 10, std::numeric_limits<uint32_t>::max()) >= target)
        {
            return i;
        }
    }

    return 0;
}

uint32_t day20::find_first_house_part2(const std::string &input)
{
    uint32_t target = std::stoul(input);

    for (uint32_t i = 1;; i++)
    {
        if (sum_of_factors(i, 11, 50) >= target)
        {
            return i;
        }
    }

    return 0;
}
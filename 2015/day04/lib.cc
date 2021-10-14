#include "lib.h"

#include <common/md5.h>

int64_t day04::mine_adventcoin(const std::string &input, size_t zero_count)
{
    int64_t count = 0;
    std::string zeros(zero_count, '0');

    while (true) {
        auto current = input + std::to_string(count);
        auto hash = common::md5(current);

        if (hash.find(zeros) == 0)
        {
            break;
        }

        count++;
    }

    return count;
}
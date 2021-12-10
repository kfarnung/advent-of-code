#include "lib.h"

#include <common/input_parser.h>

#include <unordered_map>

namespace
{
    uint8_t segments_to_byte(const std::string &input)
    {
        uint8_t byte = 0;
        for (const auto &ch : input)
        {
            uint8_t mask = 1;
            mask <<= ch - 'a';
            byte |= mask;
        }

        return byte;
    }
}

uint32_t day08::run_part1(const std::vector<std::string> &input)
{
    uint32_t count = 0;
    for (const auto &line : input)
    {
        size_t index = line.find("| ") + 2;
        auto entries = common::splitstr(line.substr(index), ' ');
        for (const auto &entry : entries)
        {
            switch (entry.size())
            {
            case 2:
            case 3:
            case 4:
            case 7:
                ++count;
            }
        }
    }

    return count;
}

uint32_t day08::run_part2(const std::vector<std::string> &input)
{
    uint32_t sum = 0;

    for (const auto &line : input)
    {
        size_t split_index = line.find('|');
        auto inputs = common::splitstr(line.substr(0, split_index - 1), ' ');
        auto outputs = common::splitstr(line.substr(split_index + 2), ' ');

        std::unordered_map<uint8_t, uint8_t> byte_to_digit;
        std::unordered_map<uint8_t, uint8_t> digit_to_byte;

        while (!inputs.empty())
        {
            std::vector<std::string> remaining;

            for (const auto &entry : inputs)
            {
                uint8_t byte = segments_to_byte(entry);

                switch (entry.size())
                {
                case 2:
                    // 1
                    byte_to_digit[byte] = 1;
                    digit_to_byte[1] = byte;
                    continue;

                case 3:
                    // 7
                    byte_to_digit[byte] = 7;
                    digit_to_byte[7] = byte;
                    continue;

                case 4:
                    // 4
                    byte_to_digit[byte] = 4;
                    digit_to_byte[4] = byte;
                    continue;

                case 5:
                {
                    // 2, 3, 5
                    auto one = digit_to_byte.find(1);
                    auto six = digit_to_byte.find(6);
                    if (one != digit_to_byte.end() && six != digit_to_byte.end())
                    {
                        if ((byte & one->second) == one->second)
                        {
                            // 3
                            byte_to_digit[byte] = 3;
                            digit_to_byte[3] = byte;
                            continue;
                        }
                        else if ((byte | six->second) == six->second)
                        {
                            // 5
                            byte_to_digit[byte] = 5;
                            digit_to_byte[5] = byte;
                            continue;
                        }
                        else
                        {
                            // 2
                            byte_to_digit[byte] = 2;
                            digit_to_byte[2] = byte;
                            continue;
                        }
                    }

                    break;
                }

                case 6:
                {
                    // 0, 6, 9
                    auto one = digit_to_byte.find(1);
                    auto four = digit_to_byte.find(4);
                    if (one != digit_to_byte.end() && four != digit_to_byte.end())
                    {
                        if ((byte & four->second) == four->second)
                        {
                            // 9
                            byte_to_digit[byte] = 9;
                            digit_to_byte[9] = byte;
                            continue;
                        }
                        else if ((byte & one->second) == one->second && (byte | four->second) != byte)
                        {
                            // 0
                            byte_to_digit[byte] = 0;
                            digit_to_byte[0] = byte;
                            continue;
                        }
                        else if ((byte | one->second) != byte)
                        {
                            // 6
                            byte_to_digit[byte] = 6;
                            digit_to_byte[6] = byte;
                            continue;
                        }
                    }

                    break;
                }

                case 7:
                    // 8
                    byte_to_digit[byte] = 8;
                    digit_to_byte[8] = byte;
                    continue;
                }

                remaining.push_back(entry);
            }

            std::swap(inputs, remaining);
        }

        uint32_t decoded = 0;
        for (const auto &output : outputs)
        {
            decoded *= 10;
            uint8_t byte = segments_to_byte(output);
            decoded += byte_to_digit[byte];
        }

        sum += decoded;
    }

    return sum;
}

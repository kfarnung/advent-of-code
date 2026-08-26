#include "lib.h"

#include <cstdint>
#include <limits>
#include <numeric>
#include <tuple>
#include <vector>

namespace
{
    struct bitstream
    {
        size_t offset;
        std::vector<uint8_t> bits;

        bitstream() : offset(0)
        {
        }

        static bitstream from_hex(const std::string &input)
        {
            bitstream bs;
            for (size_t i = 0; i < input.size(); i += 2)
            {
                auto bytestr = input.substr(i, 2);
                auto byteval = static_cast<uint8_t>(std::stoul(bytestr, nullptr, 16));
                bs.bits.emplace_back(byteval);
            }

            return bs;
        }

        uint64_t get_bits(size_t count)
        {
            size_t upper_bound = offset + count;
            uint64_t bit_value = 0;
            for (; offset < upper_bound; ++offset)
            {
                bit_value <<= 1;
                auto byte_index = offset / 8;
                auto shift_amount = 8 - (offset % 8) - 1;
                bit_value += (bits[byte_index] >> shift_amount) & 1;
            }

            return bit_value;
        }
    };

    std::tuple<int64_t, int64_t> parse_packet(bitstream &bs)
    {
        int64_t version_sum = 0;

        auto version = bs.get_bits(3);
        version_sum += version;

        auto type_id = bs.get_bits(3);
        if (type_id == 4)
        {
            // literal value
            uint64_t current_group = 0;
            int64_t value = 0;

            do
            {
                value <<= 4;
                current_group = bs.get_bits(5);
                value += current_group & 0x0F;
            } while ((current_group & 0x10) != 0);

            return std::tuple<int64_t, int64_t>(version_sum, value);
        }
        else
        {
            std::vector<int64_t> subpacket_results;

            // operators
            auto length_type_id = bs.get_bits(1);
            if (length_type_id == 0)
            {
                // length in bits
                auto total_length = bs.get_bits(15);
                auto end_offset = bs.offset + total_length;
                while (bs.offset < end_offset)
                {
                    auto sub_result = parse_packet(bs);
                    version_sum += std::get<0>(sub_result);
                    subpacket_results.emplace_back(std::get<1>(sub_result));
                }
            }
            else
            {
                // number of sub-packets
                auto packet_count = bs.get_bits(11);
                for (uint64_t i = 0; i < packet_count; ++i)
                {
                    auto sub_result = parse_packet(bs);
                    version_sum += std::get<0>(sub_result);
                    subpacket_results.emplace_back(std::get<1>(sub_result));
                }
            }

            int64_t result = 0;

            switch (type_id)
            {
            case 0:
                // sum
                result = std::accumulate(
                    begin(subpacket_results),
                    end(subpacket_results),
                    0LL,
                    [](int64_t total, int64_t current)
                    { return total + current; });
                break;

            case 1:
                // product
                result = std::accumulate(
                    begin(subpacket_results),
                    end(subpacket_results),
                    1LL,
                    [](int64_t total, int64_t current)
                    { return total * current; });
                break;

            case 2:
                // minimum
                result = std::accumulate(
                    begin(subpacket_results),
                    end(subpacket_results),
                    std::numeric_limits<int64_t>::max(),
                    [](int64_t total, int64_t current)
                    { return std::min(total, current); });
                break;

            case 3:
                // maximum
                result = std::accumulate(
                    begin(subpacket_results),
                    end(subpacket_results),
                    std::numeric_limits<int64_t>::min(),
                    [](int64_t total, int64_t current)
                    { return std::max(total, current); });
                break;

            case 5:
                // greater than
                result = subpacket_results[0] > subpacket_results[1] ? 1 : 0;
                break;

            case 6:
                // less than
                result = subpacket_results[0] < subpacket_results[1] ? 1 : 0;
                break;

            case 7:
                // equal to
                result = subpacket_results[0] == subpacket_results[1] ? 1 : 0;
                break;
            }

            return std::tuple<int64_t, int64_t>(version_sum, result);
        }
    }
}

int64_t day16::run_part1(const std::string &input)
{
    auto bs = bitstream::from_hex(input);
    return std::get<0>(parse_packet(bs));
}

int64_t day16::run_part2(const std::string &input)
{
    auto bs = bitstream::from_hex(input);
    return std::get<1>(parse_packet(bs));
}

#include "lib.h"

#include <common/input_parser.h>

#include <set>
#include <sstream>
#include <tuple>

namespace
{
    using BingoBoard = std::vector<std::vector<uint32_t>>;

    std::vector<uint32_t> parse_numbers(const std::string &input)
    {
        std::vector<uint32_t> numbers;

        auto tokens = common::splitstr(input, ',');
        for (const auto &token : tokens)
        {
            numbers.emplace_back(std::stoul(token));
        }

        return numbers;
    }

    std::tuple<std::vector<uint32_t>, std::vector<BingoBoard>> parse_input(const std::vector<std::string> &input)
    {
        std::vector<BingoBoard> boards;
        size_t index = 0;
        auto numbers = parse_numbers(input[index]);
        index += 2;

        BingoBoard current;

        while (index < input.size())
        {
            auto line = input[index];
            if (line.empty())
            {
                boards.emplace_back(std::move(current));
            }
            else
            {
                std::istringstream ss(line);
                std::vector<uint32_t> row;
                uint32_t col;
                while (ss >> col)
                {
                    row.emplace_back(col);
                }

                current.emplace_back(std::move(row));
            }

            ++index;
        }

        boards.emplace_back(std::move(current));

        return std::tuple<std::vector<uint32_t>, std::vector<BingoBoard>>{numbers, boards};
    }

    bool is_winner(BingoBoard board, const std::set<uint32_t> &chosen_numbers)
    {
        for (size_t i = 0; i < board.size(); ++i)
        {
            auto row = board[i];
            size_t j = 0;
            for (; j < row.size(); j++)
            {
                if (chosen_numbers.find(row[j]) == chosen_numbers.end())
                {
                    break;
                }
            }

            if (j == row.size())
            {
                return true;
            }
        }

        for (size_t i = 0; i < board[0].size(); ++i)
        {
            size_t j = 0;
            for (; j < board.size(); j++)
            {
                if (chosen_numbers.find(board[j][i]) == chosen_numbers.end())
                {
                    break;
                }
            }

            if (j == board.size())
            {
                return true;
            }
        }

        return false;
    }

    uint32_t score_board(BingoBoard board, std::set<uint32_t> chosen_numbers, uint32_t last_number)
    {
        uint32_t score = 0;

        for (const auto &row : board)
        {
            for (const auto &col : row)
            {
                if (chosen_numbers.find(col) == chosen_numbers.end())
                {
                    score += col;
                }
            }
        }

        return score * last_number;
    }
}

size_t day04::run_part1(const std::vector<std::string> &input)
{
    std::vector<uint32_t> numbers;
    std::vector<BingoBoard> boards;
    std::tie(numbers, boards) = parse_input(input);

    std::set<uint32_t> chosen_numbers;

    for (const auto &number : numbers)
    {
        chosen_numbers.insert(number);
        if (chosen_numbers.size() >= 5)
        {
            for (const auto &board : boards)
            {
                if (is_winner(board, chosen_numbers))
                {
                    return score_board(board, chosen_numbers, number);
                }
            }
        }
    }

    return 0;
}

size_t day04::run_part2(const std::vector<std::string> &input)
{
    std::vector<uint32_t> numbers;
    std::vector<BingoBoard> boards;
    std::tie(numbers, boards) = parse_input(input);

    std::set<uint32_t> chosen_numbers;

    for (const auto &number : numbers)
    {
        chosen_numbers.insert(number);
        if (chosen_numbers.size() >= 5)
        {
            std::vector<BingoBoard> remaining;

            for (const auto &board : boards)
            {
                if (!is_winner(board, chosen_numbers))
                {
                    remaining.emplace_back(std::move(board));
                }
                else if (boards.size() == 1)
                {
                    return score_board(board, chosen_numbers, number);
                }
            }

            std::swap(boards, remaining);
        }
    }

    return 0;
}

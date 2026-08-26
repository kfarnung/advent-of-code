#include "lib.h"

namespace
{
    int hex_digit_value(char ch)
    {
        if (ch >= 'a' && ch <= 'f')
        {
            return ch - 'a';
        }
        else if (ch >= 'A' && ch <= 'F')
        {
            return ch - 'A';
        }
        else
        {
            return ch - '0';
        }
    }

    std::string parse_string(const std::string &input)
    {
        bool found_parens = false;
        std::vector<char> parsed_characters;

        for (size_t i = 0; i < input.size(); i++)
        {
            if (input[i] == '\"')
            {
                if (!found_parens)
                {
                    found_parens = true;
                }
                else
                {
                    // We found the end
                    found_parens = false;
                    break;
                }
            }
            else if (found_parens)
            {
                if (input[i] == '\\')
                {
                    // Consume the backslash
                    i++;

                    if (input[i] == 'x')
                    {
                        // Consume the 'x' and capture the first digit.
                        i++;
                        char ch = hex_digit_value(input[i]) * 16;

                        // Consume the first digit and capture the second.
                        i++;
                        ch += hex_digit_value(input[i]);

                        // Capture the calculated character.
                        parsed_characters.emplace_back(ch);
                    }
                    else
                    {
                        // Take the character literally.
                        parsed_characters.emplace_back(input[i]);
                    }
                }
                else
                {
                    // Pass the read character through.
                    parsed_characters.emplace_back(input[i]);
                }
            }
        }

        return std::string{begin(parsed_characters), end(parsed_characters)};
    }

    std::string encode_string(const std::string &input)
    {
        std::vector<char> encoded_characters;
        encoded_characters.emplace_back('"');

        for (const auto &ch : input)
        {
            if (ch == '"' || ch == '\\')
            {
                // Need to insert a backslash.
                encoded_characters.emplace_back('\\');
            }

            // Copy the character unchanged.
            encoded_characters.emplace_back(ch);
        }
        
        encoded_characters.emplace_back('"');
        return std::string{begin(encoded_characters), end(encoded_characters)};
    }
}

size_t day08::calculate_string_overhead(const std::vector<std::string> &input)
{
    size_t code_size = 0;
    size_t memory_size = 0;

    for (const auto &line : input)
    {
        code_size += line.size();
        auto parsed = parse_string(line);
        memory_size += parsed.size();
    }

    return code_size - memory_size;
}

size_t day08::calculate_encoding_overhead(const std::vector<std::string> &input)
{
    size_t code_size = 0;
    size_t encoded_size = 0;

    for (const auto &line : input)
    {
        code_size += line.size();
        auto parsed = encode_string(line);
        encoded_size += parsed.size();
    }

    return encoded_size - code_size;
}
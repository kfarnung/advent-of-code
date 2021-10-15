#include "lib.h"

#include <regex>
#include <unordered_map>

namespace
{
    struct Connection
    {
        std::string a;
        std::string operand;
        std::string b;
    };

    std::unordered_map<std::string, Connection> parse_connections(const std::vector<std::string> &input)
    {
        std::unordered_map<std::string, Connection> connections;
        std::regex re("^(?:([a-z0-9]+) )\?\?(?:(AND|LSHIFT|NOT|OR|RSHIFT) )?([a-z0-9]+) -> ([a-z]+)$");

        for (const auto &line : input)
        {
            std::smatch sm;
            if (std::regex_match(line, sm, re))
            {
                connections[sm[4].str()] = Connection{sm[1].str(), sm[2].str(), sm[3].str()};
            }
        }

        return connections;
    }

    uint16_t get_wire_output(
        std::unordered_map<std::string, uint16_t> &memo,
        const std::unordered_map<std::string, Connection> &connections,
        const std::string &wire)
    {
        if (std::isdigit(wire[0]))
        {
            return static_cast<uint16_t>(std::atoi(wire.c_str()));
        }

        auto search = memo.find(wire);
        if (search != memo.end())
        {
            return search->second;
        }

        uint16_t value = 0;
        auto connection = connections.at(wire);
        if (connection.operand.empty())
        {
            value = get_wire_output(memo, connections, connection.b);
        }
        else if (connection.operand == "AND")
        {
            value = get_wire_output(memo, connections, connection.a) & get_wire_output(memo, connections, connection.b);
        }
        else if (connection.operand == "LSHIFT")
        {
            value = get_wire_output(memo, connections, connection.a) << get_wire_output(memo, connections, connection.b);
        }
        else if (connection.operand == "NOT")
        {
            value = ~get_wire_output(memo, connections, connection.b);
        }
        else if (connection.operand == "OR")
        {
            value = get_wire_output(memo, connections, connection.a) | get_wire_output(memo, connections, connection.b);
        }
        else if (connection.operand == "RSHIFT")
        {
            value = get_wire_output(memo, connections, connection.a) >> get_wire_output(memo, connections, connection.b);
        }

        memo[wire] = value;

        return value;
    }
}

uint16_t day07::get_wire_output(const std::vector<std::string> &input, const std::string &wire)
{
    std::unordered_map<std::string, uint16_t> memo;
    auto connections = parse_connections(input);
    return get_wire_output(memo, connections, wire);
}

uint16_t day07::get_wire_output_part2(const std::vector<std::string> &input, const std::string &wire)
{
    std::unordered_map<std::string, uint16_t> memo;
    auto connections = parse_connections(input);
    auto part1 = get_wire_output(memo, connections, wire);

    memo.clear();
    connections["b"] = Connection{"", "", std::to_string(part1)};
    return get_wire_output(memo, connections, wire);
}
#ifndef COMMON_POINT2D_H
#define COMMON_POINT2D_H

#include <functional>

namespace common
{
    struct Point2D
    {
        int x;
        int y;
    };

    bool operator==(const Point2D &lhs, const Point2D &rhs);
    Point2D operator+(const Point2D &lhs, const Point2D &rhs);
}

namespace std
{
    template <>
    struct hash<common::Point2D>
    {
        // Using the Apache.Commons prime number multiplication algorithm as
        // demonstrated by
        // http://myeyesareblind.com/2017/02/06/Combine-hash-values/
        std::size_t operator()(const common::Point2D &s) const noexcept
        {
            std::size_t hash = 17;
            hash = hash * 37 + std::hash<int>{}(s.x);
            hash = hash * 37 + std::hash<int>{}(s.y);
            return hash;
        }
    };
}

#endif
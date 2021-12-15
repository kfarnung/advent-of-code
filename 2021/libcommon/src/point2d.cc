#include <common/point2d.h>

common::Point2D::Point2D(int64_t x, int64_t y) : x(x), y(y)
{
}

bool common::operator==(const Point2D &lhs, const Point2D &rhs)
{
    return lhs.x == rhs.x && lhs.y == rhs.y;
}

bool common::operator!=(const Point2D &lhs, const Point2D &rhs)
{
    return !(lhs == rhs);
}

bool common::operator<(const Point2D &lhs, const Point2D &rhs)
{
    if (lhs.x == rhs.x)
    {
        return lhs.y < rhs.y;
    }

    return lhs.x < rhs.x;
}

common::Point2D common::operator+(const Point2D &lhs, const Point2D &rhs)
{
    return Point2D{lhs.x + rhs.x, lhs.y + rhs.y};
}

common::Point2D &common::operator+=(Point2D &lhs, const Point2D &rhs)
{
    lhs.x += rhs.x;
    lhs.y += rhs.y;
    return lhs;
}
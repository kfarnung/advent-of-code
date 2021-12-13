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
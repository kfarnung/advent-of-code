#include <common/point2d.h>

bool common::operator==(const Point2D &lhs, const Point2D &rhs)
{
    return lhs.x == rhs.x && lhs.y == rhs.y;
}

common::Point2D common::operator+(const Point2D &lhs, const Point2D &rhs)
{
    return Point2D{lhs.x + rhs.x, lhs.y + rhs.y};
}
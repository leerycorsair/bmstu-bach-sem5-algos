#include "point.h"

#include <cmath>

double degrees_to_rad(double degrees)
{
    return degrees / 180 * M_PI;
}

void point_read(point_t &point)
{
    cin >> point.x >> point.y;
}

void point_rand(point_t &point)
{
    point.x = (double)(rand() % 100) + 0.01 * (double)(rand() % 100);
    point.y = (double)(rand() % 100) + 0.01 * (double)(rand() % 100);
}

void point_print(const point_t &point)
{
    cout << point.x << " " << point.y;
}

void point_rotate(point_t &point, double angle)
{
    double tmp_x = point.x;
    double angle_pi = degrees_to_rad(angle);

    point.x = point.x * cos(angle_pi) - point.y * sin(angle_pi);
    point.y = tmp_x * sin(angle_pi) + point.y * cos(angle_pi);
}

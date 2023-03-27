#pragma once

using namespace std;

#include <iostream>
#include <iomanip>


typedef struct
{
    double x;
    double y;
} point_t;

void point_read(point_t &point);

void point_rand(point_t &point);

void point_print(const point_t &point);

void point_rotate(point_t &point, double angle);

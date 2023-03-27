#pragma once

#include "point.h"
#include <thread>
#include <vector>

typedef vector<point_t> figure_t;

figure_t figure_read(size_t cnt);

figure_t figure_generate(size_t cnt);

void figure_print(const figure_t &figure);

void figure_rotate(figure_t &figure, double angle, size_t thrd_cnt = 1);

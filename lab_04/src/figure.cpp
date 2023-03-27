#include "figure.h"

figure_t figure_read(size_t cnt)
{
    figure_t figure(cnt);
    for (size_t i = 0; i < cnt; i++)
        point_read(figure[i]);
    return figure;
}

figure_t figure_generate(size_t cnt)
{
    figure_t figure(cnt);
    for (size_t i = 0; i < cnt; i++)
        point_rand(figure[i]);
    return figure;
}

void figure_print(const figure_t &figure)
{   
    for (auto point : figure)
    {   
        point_print(point);
        cout << "\n";
    }
}

void thread_rotate(figure_t &figure, double angle, size_t start, size_t end)
{
    for (size_t i = start; i < end; i++)
        point_rotate(figure[i], angle);
}

void figure_rotate(figure_t &figure, double angle, size_t thrd_cnt)
{
    size_t elem_per_thrd = figure.size() / thrd_cnt;
    thread *thrds = new thread[thrd_cnt];

    for (size_t i = 0; i < thrd_cnt - 1; i++)
    {
        size_t start = i * elem_per_thrd;
        size_t finish = (i + 1) * elem_per_thrd;

        thrds[i] = thread(thread_rotate, ref(figure), angle, start, finish);
    }
    thrds[thrd_cnt - 1] = thread(thread_rotate, ref(figure), angle, (thrd_cnt - 1) * elem_per_thrd, figure.size());

    for (size_t i = 0; i < thrd_cnt; i++)
        thrds[i].join();
    delete[] thrds;
}

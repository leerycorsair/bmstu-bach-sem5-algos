
#include "figure.h"
#include "utils.h"
using namespace std;

#define ITERATIONS 100
#define MIN_THREADS 1
#define MAX_THREAD 32
#define RESOLUTION 1920 * 1080
#define CPU_THREADS 12
#define MIN_POINTS 1000
#define MAX_POINTS 2048000

void menu_print()
{
    cout << "Menu:\n";
    cout << "1. Read and rotate a figure.\n";
    cout << "2. Randomise and rotate a figure.\n";
    cout << "3. Comparative analysis.\n";
    cout << "0. Exit.\n";
}

int menu_option()
{
    int cmd;
    cout << "Option:";
    cin >> cmd;
    return cmd;
}

int main()
{
    srand(time(0));
    figure_t figure;
    int cmd = -1;
    while (cmd)
    {
        menu_print();
        cmd = menu_option();
        size_t cnt, thrds;
        double angle;
        uint64_t time1, time2;
        figure_t tmp_figure;

        switch (cmd)
        {
        case 1:
            cout << "Points:";
            cin >> cnt;
            cout << "Enter points' coordinates:";
            figure = figure_read(cnt);
            cout << "The original figure\n";
            figure_print(figure);

            cout << "Threads amount:";
            cin >> thrds;
            cout << "\n";

            cout << "Angle:";
            cin >> angle;
            cout << "\n";

            tmp_figure = figure;

            time1 = tick();
            figure_rotate(figure, angle);
            time2 = tick();
            cout << "Single thread:" << (time2 - time1) << " ticks\n";

            time1 = tick();
            figure_rotate(tmp_figure, angle, thrds);
            time2 = tick();
            cout << thrds << " threads:" << (time2 - time1) << " ticks\n";
            cout << "\n";

            cout << "The rotated figure\n";
            figure_print(figure);

            break;
        case 2:
            cout << "Points:";
            cin >> cnt;
            figure = figure_generate(cnt);

            cout << "Threads amount:";
            cin >> thrds;
            cout << "\n";

            cout << "Angle:";
            cin >> angle;
            cout << "\n";

            tmp_figure = figure;

            time1 = tick();
            figure_rotate(figure, angle);
            time2 = tick();
            cout << "Single thread:" << (time2 - time1) << " ticks\n";

            time1 = tick();
            figure_rotate(tmp_figure, angle, thrds);
            time2 = tick();
            cout << thrds << " threads:" << (time2 - time1) << " ticks\n";
            cout << "\n";

            break;
        case 3:
            cout << "\nPoints = " << RESOLUTION << " (default screen resolution 1920*1080)\n";
            for (size_t thrds = MIN_THREADS; thrds <= MAX_THREAD; thrds *= 2)
            {
                uint64_t sum_time = 0;
                figure = figure_generate(RESOLUTION);
                for (size_t i = 0; i < ITERATIONS; i++)
                {
                    tmp_figure = figure;
                    time1 = tick();
                    figure_rotate(tmp_figure, angle, thrds);
                    time2 = tick();
                    sum_time += time2 - time1;
                }
                cout << thrds << " thread(s) - " << sum_time / ITERATIONS << " ticks\n";
            }

            cout << "\nThreads = " <<  CPU_THREADS << "\n";
            for (size_t points = MIN_POINTS; points <= MAX_POINTS; points *= 2)
            {
                uint64_t sum_time_thrds = 0;
                uint64_t sum_time_single = 0;
                figure = figure_generate(points);
                for (size_t i = 0; i < ITERATIONS; i++)
                {
                    tmp_figure = figure;
                    time1 = tick();
                    figure_rotate(tmp_figure, angle);
                    time2 = tick();
                    sum_time_single += time2 - time1;

                    tmp_figure = figure;
                    time1 = tick();
                    figure_rotate(tmp_figure, angle, CPU_THREADS);
                    time2 = tick();
                    sum_time_thrds += time2 - time1;
                }
                cout << points << " points - " << sum_time_single / ITERATIONS << " ticks (Single)\n";
                cout << points << " points - " << sum_time_thrds / ITERATIONS << " ticks (Multithreading)\n";
            }
            break;
        }
    }
    return 0;
}
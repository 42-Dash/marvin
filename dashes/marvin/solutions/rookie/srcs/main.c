#include "dash.h"

int main(int argc, char **argv) {
    t_map map = {NULL, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL};

    if (argc != 2) {
        return 1;
    }
    init_map(&map, argv[1]);
    if (map.map == NULL) {
        return 1;
    }

    print_dummy_solution(&map);
    print_djikstra_solution(&map);
    return 0;
}

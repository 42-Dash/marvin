#include "openl.h"

int main(int argc, char **argv) {
    t_map map = {NULL, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, {0, 0, 0}};

    if (argc != 2) {
        return 1;
    }
    init_map(&map, argv[1]);
    if (map.map == NULL) {
        return 1;
    }
    print_dummy_solution(&map);
    loop_djikstra(&map);

    return 0;
}

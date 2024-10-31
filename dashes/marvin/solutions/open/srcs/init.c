#include "openl.h"

static inline int idx(const t_map *map, int row, int col) {
    return (map->width * row) + col;
}

static inline bool is_valid(t_map *map, t_point dst) {
    if (dst.row < 0 || dst.row >= map->height) {
        return false;
    }
    if (dst.col < 0 || dst.col >= map->width) {
        return false;
    }
    return true;
}

void init_graph(t_map *map) {
    const int DIR[4][2] = {{1, 0}, {0, 1}, {-1, 0}, {0, -1}};
    t_point dest_p = {0, 0};
    int from = 0, dest = 0;

    for (int row = 0; row < map->height; row++) {
        for (int col = 0; col < map->width; col++) {
            for (int iter = 0; iter < 4; iter++) {
                dest_p = (t_point){row + DIR[iter][0], col + DIR[iter][1]};

                if (is_valid(map, dest_p) == false) {
                    continue;
                }

                from = idx(map, row, col);
                dest = idx(map, dest_p.row, dest_p.col);

                if (ft_strchr(NOT_NUMERIC, map->map[dest_p.row][dest_p.col])) {
                    map->graph[from][dest] = 0;
                } else {
                    map->graph[from][dest] = map->map[dest_p.row][dest_p.col] - '0';
                }
            }
        }
    }
}

void setup_structure(t_map *map) {
    map->solutions = malloc(map->nodes * sizeof (char *));
    map->visited = malloc(map->nodes * sizeof (bool));
    map->graph = malloc(map->nodes * sizeof (int *));
    map->costs = malloc(map->nodes * sizeof (int));

    if (!map->costs || !map->solutions || !map->visited || !map->graph) {
        exit(1);
    }

    // inits 2D array with INT_MAX value
    for (int i = 0; i < map->nodes; i++) {
        map->graph[i] = malloc(map->nodes * sizeof (int));
        if (!map->graph[i]) {
            exit(1);
        }
        for (int step = 0; step < map->nodes; step++) {
            map->graph[i][step] = INT_MAX;
        }
    }
}

void preinit(t_map *map) {
    ft_bzero(map->solutions, map->nodes * sizeof (char *));
    ft_bzero(map->visited, map->nodes * sizeof (bool));

    for (int i = 0; i < map->nodes; i++) {
        map->costs[i] = INT_MAX;
    }

    map->visited[map->start] = true;
    map->costs[map->start] = 0;
}

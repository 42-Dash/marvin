#include "dash.h"

static inline int idx(t_map *map, int row, int col) {
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

static void init_graph(t_map *map) {
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

static void  setup_structure(t_map *map) {
    map->solutions = calloc(map->nodes, sizeof (char *));
    map->visited = calloc(map->nodes, sizeof (bool));
    map->graph = malloc(map->nodes * sizeof (int *));
    map->costs = malloc(map->nodes * sizeof (int));

    if (!map->costs || !map->solutions || !map->visited || !map->graph) {
        exit(1);
    }

    for (int i = 0; i < map->nodes; i++) {
        map->costs[i] = INT_MAX;
    }
    map->visited[map->start] = true;
    map->costs[map->start] = 0;

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

static int find_cheapest(t_map *map) {
    int min = INT_MAX;
    int idx = 0;

    for (int i = 0; i < map->nodes; i++) {
        if (map->visited[i] == false && map->costs[i] < min) {
            min = map->costs[i];
            idx = i;
        }
    }
    return idx;
}

// Good luck with understanding
static inline char find_direction(int from, int to) {
    if (from + 1 == to) {
        return RIGHT;
    } else if (from - 1 == to) {
        return LEFT;
    } else if (from < to) {
        return DOWN;
    } else {
        return UP;
    }
}

static char *append(char *str, char chr) {
    const int len = ft_strlen(str);
    char *new;

    new = malloc(len + 2);
    ft_memcpy(new, str, len);
    new[len] = chr;
    new[len + 1] = 0;
    return new;
}

static void recalculate_neighbors(t_map *map, int cheapest) {
    for (int dest = 0; dest < map->nodes; dest++) {
        if (map->graph[cheapest][dest] != INT_MAX && map->visited[dest] == false) {
            if (map->costs[dest] == INT_MAX) {
                map->costs[dest] = map->graph[cheapest][dest]
                    + map->costs[cheapest];
                map->solutions[dest] = append(
                    map->solutions[cheapest],
                    find_direction(cheapest, dest)
                );
            } else if (map->graph[cheapest][dest] + map->costs[cheapest] < map->costs[dest]) {
                map->costs[dest] = map->graph[cheapest][dest]
                    + map->costs[cheapest];
                map->solutions[dest] = append(
                    map->solutions[cheapest],
                    find_direction(cheapest, dest)
                );
            }
        }
    }
}

void print_djikstra_solution(t_map *map) {
    setup_structure(map);
    init_graph(map);
    int cheapest = map->start;

    do {
        recalculate_neighbors(map, cheapest);
        cheapest = find_cheapest(map);
        map->visited[cheapest] = true;
    } while (map->visited[map->end] == false);

    printf("%s\n", map->solutions[map->end]);
}

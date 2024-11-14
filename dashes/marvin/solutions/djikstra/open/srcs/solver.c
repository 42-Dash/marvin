#include "openl.h"

static inline char surf_mult(int surface) {
    if (surface > 5 || surface < 0) {
        exit(1);
    }
    return SURF_COST[surface] - '0';
}
static inline int step_cost(t_map *map, int src, int dest) {
    switch (map->surf[dest]) {
        case (WATER):
            return surf_mult(map->aff.water) * map->graph[src][dest];
        case (EARTH):
            return surf_mult(map->aff.earth) * map->graph[src][dest];
        case (AIR):
            return surf_mult(map->aff.air) * map->graph[src][dest];
        default:
            return 0;
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
                map->costs[dest] = step_cost(map, cheapest, dest)
                    + map->costs[cheapest];
                map->solutions[dest] = append(
                    map->solutions[cheapest],
                    find_direction(cheapest, dest)
                );
            } else if (step_cost(map, cheapest, dest) + map->costs[cheapest] < map->costs[dest]) {
                map->costs[dest] = step_cost(map, cheapest, dest)
                    + map->costs[cheapest];
                map->solutions[dest] = append(
                    map->solutions[cheapest],
                    find_direction(cheapest, dest)
                );
            }
        }
    }
}

static void solve_djikstra(t_map *map) {
    preinit(map);
    int cheapest = map->start;

    do {
        recalculate_neighbors(map, cheapest);
        cheapest = find_cheapest(map);
        map->visited[cheapest] = true;
    } while (map->visited[map->end] == false);
}

void loop_djikstra(t_map *map) {
    int best_cost = INT_MAX;
    // one time init
    setup_structure(map);
    init_graph(map);

    for (int w = 0; w <= 5; w++) {
        for (int e = 0; e <= 5; e++) {
            if (w + e < 5) {
                continue;
            }
            map->aff = (t_affects){w, e, 10 - w - e};
            solve_djikstra(map);

            if (map->costs[map->end] < best_cost) {
                best_cost = map->costs[map->end];
                printf(
                    "%d%d%d%s\n",
                    map->aff.water,
                    map->aff.air,
                    map->aff.earth,
                    map->solutions[map->end]
                );
            }
        }
    }
}

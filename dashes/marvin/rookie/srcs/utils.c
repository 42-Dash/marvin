#include "dash.h"

void debug(t_map *map) {
    printf("costs = {");
    for (int i = 0; i < map->nodes; i++) {
        printf("%i", map->costs[i]);
        if (i + 1 != map->nodes) {printf(",");}
    }
    printf("}\n");

    printf("vistd = {");
    for (int i = 0; i < map->nodes; i++) {
        printf("%i", map->visited[i]);
        if (i + 1 != map->nodes) {printf(",");}
    }
    printf("}\n");

    printf("solutions = {\n");
    for (int i = 0; i < map->nodes; i++) {
        if (map->solutions[i]) {
            printf("solution[%i] = %s\n",i, map->solutions[i]);
        }
    }
    printf("}\n");
}

void debug_all(t_map *map) {
    debug(map);

    printf("pathes[%d][%d] = {\n", map->nodes, map->nodes);
    for (int row = 0; row < map->nodes; row++) {
        printf("    ");
        for (int col = 0; col < map->nodes; col++) {
            if (map->graph[row][col] == INT_MAX) {
                printf("-");
            } else {
                printf("%i", map->graph[row][col]);
            }
            if (col + 1 != map->nodes) {
                printf(",");
            }
        }
        printf("\n");
    }
    printf("}\n");
    printf("begin = %d\n", map->start);
    printf("finish = %d\n", map->end);
}

inline int split_len(char **map) {
    int row = 0;

    if (map == NULL) {
        return 0;
    }
    while (map[row]) {
        row++;
    }
    return row;
}

static int map_search(t_map *map, char to_find) {
    for (int row = 0; row < map->height; row++) {
        for (int col = 0; col < map->width; col++) {
            if (map->map[row][col] == to_find) {
                return (map->width * row) + col;
            }
        }
    }
    exit(1);
}

void init_map(t_map *map, char *filename) {
    const int fd = open(filename, O_RDONLY);

    if (fd < 0) {
        exit(1);
    }
    char **lines = read_file(fd);
    close(fd);
    map->map = lines;
    map->width = ft_strlen(lines[0]) - 1;
    map->height = split_len(lines);
    map->start = map_search(map, START_CHAR);
    map->end = map_search(map, FINISH_CHAR);
    map->nodes = map->height * map->width;
    if (map->start == -1 || map->end == -1) {
        exit(1);
    }
}

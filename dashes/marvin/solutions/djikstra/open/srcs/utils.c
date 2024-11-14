#include "openl.h"

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
    read_file(fd, map);
    close(fd);
    map->width = ft_strlen(map->map[0]);
    map->height = split_len(map->map);
    map->start = map_search(map, START_CHAR);
    map->end = map_search(map, FINISH_CHAR);
    map->nodes = map->height * map->width;
    if (map->start == -1 || map->end == -1) {
        exit(1);
    }
}

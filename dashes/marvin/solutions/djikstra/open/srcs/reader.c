#include "openl.h"

static char **append(char **map, char *line) {
    char **new;
    int len;

    len = split_len(map);
    new = (char **)malloc((len + 2) * sizeof (char *));
    if (new == NULL) {
        exit(1);
    }
    memcpy(new, map, sizeof (char *) * len);
    new[len] = line;
    new[len + 1] = NULL;

    return new;
}

static inline char *extract(char *line, int offset) {
    static int cached_len = INT_MAX;
    if (cached_len == INT_MAX) {cached_len = ft_strlen(line) / 2;}
    char *new;

    new = (char *)malloc((cached_len + 1) * sizeof (char));
    if (new == NULL) {
        exit(1);
    }
    for (int i = 0; i < cached_len; i++) {
        new[i] = line[i * 2 + offset];
    }
    return new;
}

void read_file(int fd, t_map *data) {
    char **map = NULL;
    char *surf = NULL;
    char *line = NULL;

    while (( line = get_next_line(fd, O_RDONLY))) {
        if (line == NULL) {
            break;
        }
        map = append(map, extract(line, MAP_OFFSET));
        surf = ft_strcat(surf, extract(line, SURF_OFFSET));
    }
    data->map = map;
    data->surf = surf;
}

#ifndef OPENL_H
# define OPENL_H

# include <stdbool.h>
# include <stdlib.h>
# include <limits.h>
# include <string.h>
# include <stdio.h>
# include <fcntl.h>

# include "../libft/libft.h"

// map macros
# define START_CHAR 'M'
# define FINISH_CHAR 'G'
# define NOT_NUMERIC "MG"

// surface macros
# define WATER 'W'
# define AIR 'A'
# define EARTH 'E'
# define SURF_COST "865431"

// read map offsets
# define MAP_OFFSET 1
# define SURF_OFFSET 0

// directions macros
# define LEFT 'L'
# define RIGHT 'R'
# define UP 'U'
# define DOWN 'D'

# define ABS(x) ((x) < 0 ? -(x) : (x))
# define MIN(a, b) ((a) < (b) ? (a) : (b))

typedef enum e_status {
    OK,
    ERROR
}   t_status;

typedef struct s_point {
    int row, col;
}   t_point;

typedef struct s_affects {
    int water, earth, air;
}   t_affects;

typedef struct s_map {
    // map atributtes
    char **map;
    int width;
    int height;

    // start and end
    int start;
    int end;
    int nodes;

    // graph
    int **graph;
    char **solutions;
    bool *visited;
    int *costs;
    char *surf;
    t_affects aff;
}   t_map;

// init
void setup_structure(t_map *map);
void init_graph(t_map *map);
void preinit(t_map *map);

// solvers
void print_dummy_solution(const t_map *map);
void loop_djikstra(t_map *map);

// solver utils
void init_map(t_map *map, char *filename);

// reader
void read_file(int fd, t_map *map);

// utils
int split_len(char **map);
void debug_all(t_map *map);
void debug(t_map *map);

#endif // OPENL_H

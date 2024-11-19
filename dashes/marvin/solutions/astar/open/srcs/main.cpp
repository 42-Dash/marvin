#include <cstdio>
#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <deque>
#include <set>
#include <tuple>

using namespace std;

# define WATER 'W'
# define AIR 'A'
# define EARTH 'E'
# define SURF_COST "865432"

typedef struct s_coefs {
	int water;
	int earth;
	int air;
} t_coefs;


typedef struct s_point {
	int row;
	int col;
} t_point;

static inline int get_cost(const vector<string> &map, t_point point, const t_coefs &coefs) {
	const int point_cost = map[point.row][point.col] - '0';
	switch (map[point.row][point.col - 1]) {
		case WATER:
			return (SURF_COST[coefs.water] - '0') * point_cost;
		case AIR:
			return (SURF_COST[coefs.air] - '0') * point_cost;
		case EARTH:
			return (SURF_COST[coefs.earth] - '0') * point_cost;
		default:
			return 0;
	}
}

inline bool operator<(const t_point &lhs, const t_point &rhs) {
	return (lhs.row < rhs.row || (lhs.row == rhs.row && lhs.col < rhs.col));
}

inline bool operator==(const t_point &lhs, const t_point &rhs) {
	return (lhs.row == rhs.row && lhs.col == rhs.col);
}

typedef struct s_node {
	t_point point;		// current node
	int g; 				// cost from start to current node
	double h; 			// heuristic cost from current node to end
	string	path;		// path from start to current node
} t_node;

int full(const t_node &node) {
	return node.g + node.h;
}

inline bool operator<(const t_node &lhs, const t_node &rhs) {
	return (full(lhs) < full(rhs) || (full(lhs) == full(rhs) && lhs.g > rhs.g));
}

inline bool operator==(const t_node &lhs, const t_node &rhs) {
	return (lhs.point == rhs.point);
}

double heuristic_value = 0;

const vector<string> read_file(const string &filename) {
	ifstream file(filename);
	vector<string> lines;

	if (file.is_open()) {
		string line;
		while (getline(file, line)) {
			lines.push_back(line);
		}
		file.close();
	}
	return lines;
}

bool exists(const set<t_point> &closed, const t_point &node) {
	return closed.find(node) != closed.cend();
}

void add(deque<t_node> &open, const t_node &node) {
	if (open.empty()) {
		open.push_back(node);
		return;
	}

	deque<t_node>::iterator it = lower_bound(open.begin(), open.end(), node);
	if (it != open.end() && *it == node) {
		if (node < *it) {
			*it = node;
		}
	} else {
		open.insert(it, node);
	}
}

t_point find_char(const vector<string> &map, char c) {
	t_point point = { -1, -1 };

	for (size_t i = 0; i < map.size(); i++) {
		const unsigned long col = map[i].find(c);
		if (col != string::npos) {
			point.row = i;
			point.col = col + 1;
			break;
		}
	}
	return point;
}

static inline double heuristic(const t_point &start, const t_point &end) {
	return (abs(start.row - end.row) + abs(start.col - end.col) / 2) * heuristic_value;
}

static inline char calc_path(const t_point &current, const t_point &next) {
	if (current.row < next.row) {
		return 'D';
	} else if (current.row > next.row) {
		return 'U';
	} else if (current.col < next.col) {
		return 'R';
	} else {
		return 'L';
	}
}

// A* algorithm
pair<string, int> find_path(
	const vector<string> &map,
	const t_point &start,
	const t_point &end,
	const t_coefs &coefs
) {
	vector<pair<int, int> > directions;
	deque<t_node>	open;
	set<t_point>	closed;

	t_node node = { start, 0, heuristic(start, end), "" };
	add(open, node);
	directions.push_back(make_pair(-1, 0));
	directions.push_back(make_pair(1, 0));
	directions.push_back(make_pair(0, 2));
	directions.push_back(make_pair(0, -2));

	while (!open.empty()) {
		t_node current = open.front();
		open.pop_front();

		if (current.point == end) {
			return pair<string, int>(current.path, current.g);
		}

		for (int i = 0; i < 4; i++) {
			const pair<int, int> &dir = directions[i];
			t_point next = { current.point.row + dir.first, current.point.col + dir.second };

			if (next.row < 0 || next.row >= map.size() || next.col < 0 || next.col >= map[next.row].size()) {
				continue;
			}

			t_node neighbor = {
				next,
				current.g + get_cost(map, next, coefs),
				heuristic(next, end),
				current.path + calc_path(current.point, next)
			};

			if (exists(closed, neighbor.point)) {
				continue;
			}

			add(open, neighbor);
		}
		closed.insert(current.point);
	}
	return pair<string, int>("You shall not pass!", 0); // another trust issue
}

static double count_heuristic_value(const vector<string> &map, const t_coefs &coefs) {
	double sum = 0;
	double count = 0;

	for (int i = 0; i < map.size(); i++) {
		for (int j = 1; j < map[i].size(); j += 2) {
			if (map[i][j] >= '1' && map[i][j] <= '9') {
				sum += get_cost(map, (t_point){ i, j }, coefs);
				count++;
			}
		}
	}
	return (count > 0) ? sum / count : 0.;
}

static void count_character_points(const vector<string> &map, const t_point &start, const string &path, int count[3]) {
    t_point current = start;
    for (std::string::const_iterator it = path.begin(); it != path.end(); ++it) {
        switch (*it) {
            case 'U':
                current.row--;
                break;
            case 'D':
                current.row++;
                break;
            case 'L':
                current.col -= 2;
                break;
            case 'R':
                current.col += 2;
                break;
            default:
                break;
        }
        if (map[current.row][current.col - 1] == WATER) {
            count[0] += map[current.row][current.col] - '0';
        } else if (map[current.row][current.col - 1] == EARTH) {
            count[1] += map[current.row][current.col] - '0';
        } else if (map[current.row][current.col - 1] == AIR) {
            count[2] += map[current.row][current.col] - '0';
        }
    }
}

int main(int argc, char **argv) {
	if (!(argc == 2 || argc == 3 || argc == 4)) {
		return 1;
	}
	const vector<string> &input = read_file(argv[1]);
	const t_point start = find_char(input, 'M');
	const t_point end = find_char(input, 'G');

	if (start.row == -1 || end.row == -1) { // trust issues
		return 1;
	}


    if (argc == 4) { // test purposes only
        heuristic_value = stof(argv[2]);
        t_coefs coefs = { argv[3][0] - '0', argv[3][1] - '0', argv[3][2] - '0' };
        pair<string, int> result = find_path(input, start, end, coefs);
        cout << coefs.water << coefs.earth << coefs.air << result.first << endl;
    } else if (argc == 3) {// test purposes only
        heuristic_value = stof(argv[2]);
        t_coefs coefs = { 5, 5, 5 };
        pair<string, int> result = find_path(input, start, end, coefs);
        int count[3] = { 0, 0, 0 };
        count_character_points(input, start, result.first, count);
        int best = 2147483647;
        string best_path = "";
        for (int w = 0; w <= 5; w++) {
            for (int e = 0; e <= 5; e++) {
                if (w + e < 5) {
                    continue;
                }
                coefs = (t_coefs){ w, e, 10 - w - e };
                if (count[0] * (SURF_COST[coefs.water] - '0')
                    + count[1] * (SURF_COST[coefs.earth] - '0')
                    + count[2] * (SURF_COST[coefs.air] - '0') < best) {
                    best = count[0] * (SURF_COST[coefs.water] - '0') + count[1] * (SURF_COST[coefs.earth] - '0') + count[2] * (SURF_COST[coefs.air] - '0');
                    stringstream ss;
                    ss << coefs.water << coefs.earth << coefs.air << result.first;
                    best_path = ss.str();
                }
            }
        }
        cout << best_path << endl;
    } else if (argc == 2) {
        int best = 2147483647;

        // A* is launched 21 times
        for (int w = 0; w <= 5; w++) {
            for (int e = 0; e <= 5; e++) {
                if (w + e < 5) {
                    continue;
                }
                t_coefs coefs = { w, e, 10 - w - e };
                heuristic_value = count_heuristic_value(input, coefs);
                pair<string, int> result = find_path(input, start, end, coefs);
                if (result.second < best) {
                    cout << coefs.water << coefs.earth << coefs.air << result.first << endl;
                    best = result.second;
                }
            }
        }

        // now the algo finds only one path, where heuristic is manhattan distance * heuristic_value cost
        // A* where heuristic is 0 - Dijkstra (always finds the cheapest path)
        // if you want to find all the paths, you can uncomment the following code

        // cout << "Dijkstra" << endl;
        // heuristic_value = 0;
        // for (int w = 0; w <= 5; w++) {
        //     for (int e = 0; e <= 5; e++) {
        //         if (w + e < 5) {
        //             continue;
        //         }
        // 		t_coefs coefs = { w, e, 10 - w - e };
        // 		pair<string, int> result = find_path(input, start, end, coefs);
        // 		if (result.second < best) {
        // 			cout << coefs.water << coefs.earth << coefs.air << result.first << endl;
        // 			best = result.second;
        // 		}
        // 	}
        // }
    }

	return 0;
}

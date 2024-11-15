#include <iostream>
#include <fstream>
#include <vector>

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

string find_path(const t_point &start, const t_point &end) {
	const int rows = end.row - start.row;
	const int cols = end.col - start.col;

	return string(abs(rows), rows > 0 ? 'D' : 'U') + string(abs(cols) / 2, cols > 0 ? 'R' : 'L');
}

t_coefs summarize_costs(
	const vector<string> &map,
	const t_point &start,
	const string &path
) {
	t_coefs coefs = { 0, 0, 0 };
	t_point point = start;

	for (size_t i = 0; i < path.size(); i++) {
		switch (path[i]) {
			case 'U':
				point.row--;
				break;
			case 'D':
				point.row++;
				break;
			case 'L':
				point.col -= 2;
				break;
			case 'R':
				point.col += 2;
				break;
		}
		switch (map[point.row][point.col - 1]) {
			case WATER:
				coefs.water += map[point.row][point.col] - '0';
				break;
			case AIR:
				coefs.air += map[point.row][point.col] - '0';
				break;
			case EARTH:
				coefs.earth += map[point.row][point.col] - '0';
				break;
		}
	}
	return coefs;
}

int main(int argc, char **argv) {
	if (argc != 2) {
		return 1;
	}
	const vector<string> &input = read_file(argv[1]);
	const t_point start = find_char(input, 'M');
	const t_point end = find_char(input, 'G');

	if (start.row == -1 || end.row == -1) { // trust issues
		return 1;
	}

	const string &path = find_path(start, end);
	const t_coefs total = summarize_costs(input, start, path);
	int best = 2147483647;
	int cost = 0;

    for (int w = 0; w <= 5; w++) {
        for (int e = 0; e <= 5; e++) {
            if (w + e < 5) {
                continue;
            }
			t_coefs coefs = { w, e, 10 - w - e };
			cost = total.water * (SURF_COST[coefs.water] - '0')
				+ total.earth * (SURF_COST[coefs.earth] - '0')
				+ total.air * (SURF_COST[coefs.air] - '0');
			if (cost < best) {
				cout << coefs.water << coefs.earth << coefs.air << path << endl;
				best = cost;
			}
		}
	}

	return 0;
}

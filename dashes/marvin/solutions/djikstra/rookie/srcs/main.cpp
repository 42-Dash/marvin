#include <unordered_set>
#include <iostream>
#include <fstream>
#include <vector>
#include <deque>

using namespace std;

typedef struct s_point {
	int row;
	int col;
} t_point;

template <> struct hash<s_point> {
	size_t operator()(const s_point& p) const {
		return (hash<int>()(p.row) << 1) ^ hash<int>()(p.col);
	}
};

inline bool operator==(const t_point &lhs, const t_point &rhs) {
	return (lhs.row == rhs.row && lhs.col == rhs.col);
}

typedef struct s_node {
	t_point point;		// current node
	int g; 				// cost from start to current node
	string	path;		// path from start to current node
} t_node;

inline bool operator<(const t_node &lhs, const t_node &rhs) {
	return lhs.g < rhs.g;
}

inline bool operator==(const t_node &lhs, const t_node &rhs) {
	return (lhs.point == rhs.point);
}

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

// complexity: O(1), underlaying data structure is a hash table
bool exists(const unordered_set<t_point> &closed, const t_point &node) {
	return closed.find(node) != closed.cend();
}

// complexity: O(log(n)), uses binary search
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

// complexity: O(m*n), where m is the number of rows and n is the number of columns
t_point find_char(const vector<string> &map, char c) {
	t_point point = { -1, -1 };

	for (size_t i = 0; i < map.size(); i++) {
		const unsigned long col = map[i].find(c);
		if (col != string::npos) {
			point.row = i;
			point.col = col;
			break;
		}
	}
	return point;
}

static inline int calc_g(const t_node &current, const t_point &next, const vector<string> &map) {
	if (map[next.row][next.col] == 'G') {
		return current.g;
	}
	return current.g + map[next.row][next.col] - '0';
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

// Djikstra's algorithm
string find_path(
	const vector<string> &map,
	const t_point &start,
	const t_point &end
) {
	vector<pair<int, int> >		directions;
	deque<t_node>				open;
	unordered_set<t_point>		closed;

	t_node node = { start, 0, "" };
	add(open, node);
	directions.push_back(make_pair(-1, 0));
	directions.push_back(make_pair(1, 0));
	directions.push_back(make_pair(0, 1));
	directions.push_back(make_pair(0, -1));

	while (!open.empty()) {
		t_node current = open.front();
		open.pop_front();

		if (current.point == end) {
			return current.path;
		}

		for (int i = 0; i < 4; i++) {
			const pair<int, int> &dir = directions[i];
			t_point next = { current.point.row + dir.first, current.point.col + dir.second };

			if (next.row < 0 || next.row >= map.size() || next.col < 0 || next.col >= map[next.row].size()) {
				continue;
			}

			t_node neighbor = {
				next,
				calc_g(current, next, map),
				current.path + calc_path(current.point, next)
			};

			if (exists(closed, neighbor.point)) {
				continue;
			}

			add(open, neighbor);
		}
		closed.insert(current.point);
	}
	return "You shall not pass!"; // another trust issue
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

	cout << find_path(input, start, end) << endl;
	return 0;
}

#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

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
			point.col = col;
			break;
		}
	}
	return point;
}

// Dummy approach
string find_path(const t_point &start, const t_point &end) {
	const int rows = end.row - start.row;
	const int cols = end.col - start.col;

	return string(abs(rows), rows > 0 ? 'D' : 'U')
		+ string(abs(cols), cols > 0 ? 'R' : 'L');
}

int main(int argc, char **argv) {
	if (argc != 2) {
		return 1;
	}
	const vector<string> &input = read_file(argv[1]);
	const t_point &start = find_char(input, 'M');
	const t_point &end = find_char(input, 'G');

	if (start.row == -1 || end.row == -1) { // trust issues
		return 1;
	}

	cout << find_path(start, end) << endl;
	return 0;
}

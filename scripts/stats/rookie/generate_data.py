import pandas as pd
import numpy as np
import argparse
import time
import os


NUMBER_OF_TESTS = 10
WEIGHTS_RANGE = np.arange(0, 10, 0.2)

def run_test(exec: str, map: str, weight: float) -> tuple[str, float]:
	'''Runs executables with a map passed as argument and weight for heuristic
	:param exec: The path to the executable
	:param map: The path to the map (passed as first argument)
	:param weight: The weight for the heuristic function (passed as second argument)
	:return: A tuple with the output and the time
	'''

	start = time.time()
	fileStream = os.popen(f'{exec} {map} {weight}')
	end = time.time()
	output = fileStream.read()

	return output, end - start

def get_score(output: str, inputfile: str) -> int:
	'''Get the cost of the output
	:param output: The output of the test
	:param inputfile: The file with map
	:return: The score
	'''
	lines = []
	with open(inputfile, 'r') as file:
		lines = file.readlines()

	row, col = -1, -1
	for i, line in enumerate(lines):
		for j, char in enumerate(line):
			if char == 'M':
				row, col = i, j
				break

	if row == -1 or col == -1: # trust issue
		print('The map does not have a starting point')
		return -1

	score = 0
	for char in output:
		if char == 'U':
			row -= 1
		elif char == 'D':
			row += 1
		elif char == 'L':
			col -= 1
		elif char == 'R':
			col += 1
		if lines[row][col].isdigit():
			score += int(lines[row][col])

	return score


def generate_sample(exec: str, inputfile: str, weight: int) -> pd.Series:
	'''Generate the samples
	:param exec: The path to the executable
	:param maps: The paths to the maps
	:return: A dictionary with the samples
	'''
	data = pd.Series()

	times = []
	output = ''

	for _ in range(NUMBER_OF_TESTS):
		output, time = run_test(exec, inputfile, weight)
		times.append(time)

	data['weight'] = weight
	data['output'] = output.strip()
	data['score'] = get_score(output, inputfile)
	data['time_average'] = np.mean(times)
	data['time_min'] = np.min(times)
	data['time_max'] = np.max(times)

	return data

def generate_samples(exec: str, maps: list[str]) -> pd.DataFrame:
	'''Generate the samples
	:param exec: The path to the executable
	:param maps: The maps to test
	:return: A DataFrame with the samples
	'''

	for map in maps:
		data = pd.DataFrame()
		for i, weight in enumerate(WEIGHTS_RANGE):
			data[i] = generate_sample(exec, map, weight)
		data.transpose().to_csv(f'{os.path.basename(map)}.csv')

	return data.transpose()

def arguments() -> argparse.Namespace:
	'''Parse the arguments of the script'''
	ap = argparse.ArgumentParser(description='Plot the results of the tests')
	ap.add_argument('exec', type=str, help='Path to the executable to test')
	ap.add_argument('maps', type=str, nargs='+', help='Paths to the maps to test')
	return ap.parse_args()

def file_exists(file: str) -> bool:
	'''Check if the file exists
	:param file: The file to check
	:return: True if the file exists, False otherwise
	'''
	try:
		with open(file, 'r'):
			pass
	except FileNotFoundError:
		return False
	return True

def validate_files(exec: str, maps: list[str]) -> bool:
	'''Validate the files
	:param exec: The path to the executable
	:param maps: The paths to the maps
	:return: True if the files exist, False otherwise
	'''
	valid = True
	if not file_exists(exec):
		print('The executable does not exist')
		valid = False
	for map in maps:
		if not file_exists(map):
			print(f'The map {map} does not exist')
			valid = False
	return valid

def main():
	'''Main function'''
	args = arguments()
	if not validate_files(args.exec, args.maps):
		return

	generate_samples(args.exec, args.maps)

if __name__ == '__main__':
	main()

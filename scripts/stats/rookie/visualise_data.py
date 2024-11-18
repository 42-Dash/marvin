import pandas as pd
import matplotlib.pyplot as plt
from matplotlib.gridspec import GridSpec
import statsmodels.api as sm
from PIL import Image
import numpy as np
import argparse


def plot_times(df: pd.DataFrame, axes: plt.Axes):
	'''Plots the time metrics of the dataset with smoothed curves using LOWESS
	Helper function for plot_samples
	'''
	time_metrics = ['time_average', 'time_min', 'time_max']
	time_labels = ['Average Execution Time', 'Min Execution Time', 'Max Execution Time']
	colors = ['blue', 'orange', 'red']

	for metric, label, color in zip(time_metrics, time_labels, colors):
		smoothed = sm.nonparametric.lowess(df[metric], df['weight'], frac=0.3)
		axes.plot(smoothed[:, 0], smoothed[:, 1], label=label, color=color, linewidth=2)

	axes.set_title('Time Taken to Execute the Program', fontsize=16, fontweight='bold')
	axes.set_ylabel('Time (s)', fontsize=10)
	axes.legend(title='Metrics', fontsize=8)
	axes.grid(True, which='both', linestyle='--', linewidth=0.5)


def plot_scores(df: pd.DataFrame, axes: plt.Axes):
	'''Plots the scores of the dataset with smoothed curves using LOWESS
	Helper function for plot_samples
	'''
	smoothed_score = sm.nonparametric.lowess(df['score'], df['weight'], frac=0.3)

	axes.scatter(df['weight'], df['score'], color='purple', alpha=0.6, label='Score Data Points')
	axes.plot(smoothed_score[:, 0], smoothed_score[:, 1], color='green', linewidth=2, label='Score curve')

	# Customize the third subplot
	axes.set_title('Score vs Weight', fontsize=16, fontweight='bold')
	axes.set_xlabel('Weight', fontsize=10)
	axes.set_ylabel('Score', fontsize=10)
	axes.legend(title='', fontsize=8)
	axes.grid(True, which='both', linestyle='--', linewidth=0.5)


def plot_image(picture: str, axes: plt.Axes, paths: pd.Series):
	'''Efficiently creates an image from a map'''
	def fill_map(lines: list[str], row: int, col: int, path: str, color: tuple[int, int, int]):
		for char in path:
			if char == 'U':
				row -= 1
			elif char == 'D':
				row += 1
			elif char == 'L':
				col -= 1
			elif char == 'R':
				col += 1

			if lines[row][col].isdigit():
				img_array[row, col] = color

	lines = open(picture).read().splitlines()
	rows = len(lines)
	cols = len(lines[0])

	img_array = np.zeros((rows, cols, 3), dtype=np.uint8)

	for y, line in enumerate(lines):
		for x, char in enumerate(line):
			if char == 'M':
				img_array[y, x] = [0, 0, 0]
			elif char == 'G':
				img_array[y, x] = [0, 0, 0]
			elif char.isdigit():
				intensity = int(char) * 10
				img_array[y, x] = [intensity, intensity, intensity]

	start_row, start_col = -1, -1
	for i, line in enumerate(lines):
		for j, char in enumerate(line):
			if char == 'M':
				start_row, start_col = i, j
				break

	for path, color in zip(paths, range(0, 255, 255 // len(paths))):
		fill_map(lines, start_row, start_col, path, (color, 130, 20))
		# TODO: Change the color

	img = Image.fromarray(img_array, 'RGB')

	img_resized = img.resize((rows * 2, rows * 2), Image.NEAREST)

	img_np = np.array(img_resized)

	axes.imshow(img_np)
	axes.axis('off')
	axes.set_title('Custom Map Image', fontsize=14, fontweight='bold')


def generate_image(filename: str, paths: pd.Series):
	'''Generates the image of the map'''

def plot_samples(data: str, picture: str):
	'''Plots the dataset with smoothed curves using LOWESS and adds an additional plot'''
	# Read the CSV file
	df = pd.read_csv(data)

	fig, axes = plt.subplots(2, 2, figsize=(14, 7), sharex=False)
	gs = GridSpec(2, 2, figure=fig, wspace=0.1, hspace=0.1)

	ax1 = fig.add_subplot(gs[0, 0])
	plot_times(df, ax1)

	ax2 = fig.add_subplot(gs[1, 0])
	plot_scores(df, ax2)

	ax3 = fig.add_subplot(gs[:, 1])
	plot_image(picture, ax3, df['output'])

	fig.delaxes(axes[0, 0])
	fig.delaxes(axes[0, 1])
	fig.delaxes(axes[1, 0])
	fig.delaxes(axes[1, 1])

	plt.show()


def arguments() -> argparse.Namespace:
	'''Parse the arguments of the script'''
	parser = argparse.ArgumentParser(description='Visualise the data')
	parser.add_argument('data', type=str, help='The path to the dataset')
	parser.add_argument('picture', type=str, help='The path to save the picture')
	return parser.parse_args()


def main():
	'''Main function'''
	args = arguments()
	plot_samples(args.data, args.picture)


if __name__ == '__main__':
	main()

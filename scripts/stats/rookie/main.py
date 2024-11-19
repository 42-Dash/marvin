from matplotlib import gridspec, pyplot as plt, patches
from tqdm import tqdm
from PIL import Image
import statsmodels.api as sm
import pandas as pd
import numpy as np
import argparse
import time
import os

NUMBER_OF_TESTS = 3
WEIGHTS_RANGE = np.arange(10, 2, -1)


def plot_times(df: pd.DataFrame, axes: plt.Axes):
    '''Plots the time metrics of the dataset with smoothed curves using LOWESS
    Helper function for plot_samples
    '''
    time_metrics = ['time_average', 'time_min', 'time_max']
    colors = ['blue', 'orange', 'red']
    time_labels = [
        'Average Execution Time',
        'Min Execution Time',
        'Max Execution Time'
    ]

    for metric, label, color in zip(time_metrics, time_labels, colors):
        smoothed = sm.nonparametric.lowess(df[metric], df['weight'], frac=0.3)
        axes.plot(smoothed[:, 0], smoothed[:, 1], label=label, color=color, linewidth=2)

    axes.set_title('Time Taken to Execute the Program', fontsize=16, fontweight='bold')
    axes.set_ylabel('Time (s)', fontsize=10)
    axes.set_xlabel('')
    axes.legend(title='Metrics', fontsize=8)
    axes.grid(True, which='both', linestyle='--', linewidth=0.5)
    axes.tick_params(axis='x', which='both', bottom=False, top=False, labelbottom=False)


def plot_scores(df: pd.DataFrame, axes: plt.Axes):
    '''Plots the scores of the dataset with smoothed curves using LOWESS
    Helper function for plot_samples
    '''
    smoothed_score = sm.nonparametric.lowess(
            df['score'], df['weight'], frac=0.3
        )

    axes.scatter(df['weight'], df['score'], color='purple', alpha=0.6, label='Score Data Points')
    axes.plot(smoothed_score[:, 0], smoothed_score[:, 1], color='green', linewidth=2, label='Score curve')

    # Customize the third subplot
    axes.set_title('Score vs Weight', fontsize=16, fontweight='bold')
    axes.set_xlabel('Weight', fontsize=10)
    axes.set_ylabel('Score', fontsize=10)
    axes.legend(title='', fontsize=8)
    axes.grid(True, which='both', linestyle='--', linewidth=0.5)


def plot_image(picture: str, axes: plt.Axes, paths: pd.Series, scores: pd.Series):
    '''Efficiently creates an image from a map'''
    def plot_path_to_image(lines, row, col, path, color):
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
                img_array[y, x] = [0, 200, 200]
            if char == 'G':
                img_array[y, x] = [0, 255, 0]
            elif char.isdigit():
                intensity = int(char) * 10
                img_array[y, x] = [intensity, intensity, intensity]

    start_row, start_col = -1, -1
    for i, line in enumerate(lines):
        for j, char in enumerate(line):
            if char == 'M':
                start_row, start_col = i, j
                break

    colors = plt.get_cmap('viridis', len(paths))(np.linspace(0, 1, len(paths)))
    legend_elements = [patches.Patch(facecolor=color, label=f'Cost {score}') for color, score in zip(colors, scores)]

    for path, color in zip(paths, colors):
        color_rgb = tuple(int(255 * c) for c in color[:3])
        plot_path_to_image(lines, start_row, start_col, path, color_rgb)

    img = Image.fromarray(img_array, 'RGB')
    img_resized = img.resize((cols * 5, rows * 5), Image.NEAREST)
    img_np = np.array(img_resized)

    axes.imshow(img_np)
    axes.axis('off')
    axes.set_title('Custom Map Image', fontsize=14, fontweight='bold')
    axes.legend(handles=legend_elements, loc='upper right', fontsize=8)


def plot_samples(df: pd.DataFrame, picture: str):
    '''Plots the dataset with smoothed curves using LOWESS'''

    fig, axes = plt.subplots(2, 2, figsize=(14, 7), sharex=False)
    gs = gridspec.GridSpec(2, 2, figure=fig, wspace=0.1, hspace=0.1)

    ax1 = fig.add_subplot(gs[0, 0])
    plot_times(df, ax1)

    ax2 = fig.add_subplot(gs[1, 0])
    plot_scores(df, ax2)

    ax3 = fig.add_subplot(gs[:, 1])
    df.sort_values('score', inplace=True)
    plot_image(picture, ax3, df['output'], df['score'])

    fig.delaxes(axes[0, 0])
    fig.delaxes(axes[0, 1])
    fig.delaxes(axes[1, 0])
    fig.delaxes(axes[1, 1])

    fig.canvas.manager.set_window_title(f'Pathfinding for {os.path.basename(picture)}')
    plt.show()


def run_test(exec: str, map: str, weight: float) -> tuple[str, float]:
    '''Runs executables with a map passed as argument and weight for heuristic
    :param exec: The path to the executable
    :param map: The path to the map (first argument)
    :param weight: The weight for the heuristic function (second argument)
    :return: A tuple with the output and the time
    '''

    start = time.time()
    output = os.popen(f'{exec} {map} {weight}').read()
    end = time.time()

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

    if row == -1 or col == -1:  # trust issue
        print('The map does not have a starting point')
        exit(1)

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


def generate_samples(exec: str, map: str) -> pd.DataFrame:
    '''Generate the samples
    :param exec: The path to the executable
    :param maps: The maps to test
    :return: A DataFrame with the samples
    '''

    data = pd.DataFrame()
    bar_format = "{l_bar}{bar}| {percentage:3.0f}%"
    for i, weight in enumerate(tqdm(WEIGHTS_RANGE, bar_format=bar_format)):
        data[i] = generate_sample(exec, map, weight)

    return data.transpose()


def arguments() -> argparse.Namespace:
    '''Parse the arguments of the script'''
    ap = argparse.ArgumentParser(description='Plot the results of the tests')
    ap.add_argument('exec', type=str, help='Path to the executable to test')
    ap.add_argument('inputfile', type=str, help='Path to the test map')
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


def validate_files(exec: str, map: str) -> bool:
    '''Validate the files
    :param exec: The path to the executable
    :param maps: The paths to the maps
    :return: True if the files exist, False otherwise
    '''
    if not file_exists(exec):
        print('The executable does not exist')
        exit(1)
    if not file_exists(map):
        print(f'The inputfile {map} does not exist')
        exit(1)


def store_dataframe(df: pd.DataFrame, filename: str):
    '''Backup the dataframe to a MD file'''
    df = df.sort_values('weight', ascending=True)
    df = df[['weight', 'score', 'time_average', 'time_min', 'time_max', 'output']]
    df.to_markdown(f'{os.path.basename(filename)}_results.md', index=False)


def main():
    '''Main function'''
    args = arguments()
    validate_files(args.exec, args.inputfile)
    df = generate_samples(args.exec, args.inputfile)
    store_dataframe(df, args.inputfile)
    plot_samples(df, args.inputfile)


if __name__ == '__main__':
    main()

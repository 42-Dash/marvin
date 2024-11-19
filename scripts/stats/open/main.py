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
HEURISTICS_COEFFICIENTS = 18.667


def plot_image(picture: str, df: pd.DataFrame):
    '''Efficiently creates an image from a map'''
    def plot_path_to_image(coefs, row, col, path, color):
        for char in path:
            if char == 'U':
                row -= 1
            elif char == 'D':
                row += 1
            elif char == 'L':
                col -= 1
            elif char == 'R':
                col += 1

            if coefs[row][col].isdigit():
                img_array[row, col] = color

    lines = open(picture).read().splitlines()
    map_surfaces = [''.join(line[::2]) for line in lines]
    map_coefs = [''.join(line[1::2]) for line in lines]
    rows = len(lines)
    cols = len(lines[0]) // 2

    img_array = np.zeros((rows, cols, 3), dtype=np.uint8)

    for row in range(len(map_surfaces)):
        for col in range(len(map_surfaces[row])):
            if map_surfaces[row][col] in 'WEA':
                intensity = int(map_coefs[row][col]) * 25
                match map_surfaces[row][col]:
                    case 'W':
                        img_array[row, col] = [0, 0, intensity]
                    case 'E':
                        img_array[row, col] = [0, intensity, 0]
                    case 'A':
                        img_array[row, col] = [intensity, 0, 0]
            elif map_surfaces[row][col] == 'M':
                img_array[row, col] = [0, 200, 200]
            elif map_surfaces[row][col] == 'G':
                img_array[row, col] = [0, 255, 0]

    start_row, start_col = -1, -1
    for row in range(len(map_surfaces)):
        for col in range(len(map_surfaces[row])):
            if map_surfaces[row][col] == 'M':
                start_row, start_col = row, col
                break
        if start_row != -1:
            break

    colors = plt.get_cmap('Reds', df.shape[0])(np.linspace(0, 1, df.shape[0]))
    legend_elements = [patches.Patch(facecolor=color, label=f'Char {char}') for color, char in zip(colors, df['character'])]

    for path, color in zip(df['output'], colors):
        color_rgb = tuple(int(255 * c) for c in color[:3])
        plot_path_to_image(map_coefs, start_row, start_col, path, color_rgb)

    img = Image.fromarray(img_array, 'RGB')
    img_resized = img.resize((cols * 5, rows * 5), Image.NEAREST)
    img_np = np.array(img_resized)

    _ = plt.figure(figsize=(8, 6))
    plt.imshow(img_np)
    plt.axis('off')
    plt.title('Custom Map Image', fontsize=14, fontweight='bold')
    plt.legend(handles=legend_elements, loc='upper right', fontsize=8)

    plt.show()


def run_test(exec: str, map: str, argv: str) -> tuple[str, float]:
    '''Runs executables with a map passed as argument and weight for heuristic
    :param exec: The path to the executable
    :param map: The path to the map (first argument)
    :param argv: The arguments to pass to the executable
    :return: A tuple with the output and the time
    '''

    start = time.time()
    output = os.popen(f'{exec} {map} {argv}').read()
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
        if row != -1:
            break

    if row == -1 or col == -1:  # trust issue
        print('The map does not have a starting point')
        exit(1)

    # 865432
    score = 0
    multiplicators = [4, 3, 2.5, 2, 1.5, 1]
    w, e, a = int(output[0]), int(output[1]), int(output[2])
    for char in output[3:]:
        if char == 'U':
            row -= 1
        elif char == 'D':
            row += 1
        elif char == 'L':
            col -= 2
        elif char == 'R':
            col += 2
        if lines[row][col] in "WEA":
            match lines[row][col]:
                case 'W':
                    score += int(lines[row][col + 1]) * multiplicators[w]
                case 'E':
                    score += int(lines[row][col + 1]) * multiplicators[e]
                case 'A':
                    score += int(lines[row][col + 1]) * multiplicators[a]

    return score


def generate_sample(exec: str, inputfile: str, input: str, adjusted: bool) -> pd.Series:
    '''Generate the samples
    :param exec: The path to the executable
    :param maps: The paths to the maps
    :return: A dictionary with the samples
    '''
    data = pd.Series()

    times = []
    output = ''

    for _ in range(NUMBER_OF_TESTS):
        output, time = run_test(exec, inputfile, input)
        times.append(time)

    data['score'] = get_score(output.strip(), inputfile)
    data['character'] = output.strip()[:3]
    data['time_average'] = np.mean(times)
    data['time_min'] = np.min(times)
    data['time_max'] = np.max(times)
    data['adjusted'] = adjusted
    data['output'] = output.strip()[3:]

    return data


def generate_samples(exec: str, map: str) -> pd.DataFrame:
    '''Generate the samples
    :param exec: The path to the executable
    :param maps: The maps to test
    :return: A DataFrame with the samples
    '''

    inputs = [f'{HEURISTICS_COEFFICIENTS}']
    for w in range(6):
        for e in range(6):
            if w + e < 5:
                continue
            inputs.append(f'{HEURISTICS_COEFFICIENTS} {w}{e}{10 - w - e}')

    assert len(inputs) == 22

    data = pd.DataFrame()
    bar_format = "{l_bar}{bar}| {percentage:3.0f}%"
    for i, weight in enumerate(tqdm(inputs, bar_format=bar_format)):
        data[i] = generate_sample(exec, map, weight, i == 0)

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
    df.to_markdown(f'{os.path.basename(filename)}_results.md', index=False)


def main():
    '''Main function'''
    args = arguments()
    validate_files(args.exec, args.inputfile)
    df = generate_samples(args.exec, args.inputfile)
    df = df.sort_values('score', ascending=True)
    store_dataframe(df, args.inputfile)
    plot_image(args.inputfile, df)


if __name__ == '__main__':
    main()

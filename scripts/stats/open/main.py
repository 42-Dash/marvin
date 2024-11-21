from tqdm import tqdm
import pandas as pd
import numpy as np
import argparse
import json
import time
import os

PATH_TO_STORAGE = 'results'
NUMBER_OF_TESTS = 3
HEURISTICS_COEFFICIENTS = 15.667


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


def find_starting_point(lines: list[str], symbol: str) -> tuple[int, int]:
    '''Find the starting point in the map
    :param lines: The content of the map
    :param symbol: The symbol to find
    :return: The row and column of the starting point
    '''
    for i, line in enumerate(lines):
        for j, char in enumerate(line):
            if char == 'M':
                return i, j

    print('The map does not have a starting point')
    exit(1) # trust issue


def get_score(output: str, inputfile: str) -> int:
    '''Get the cost of the output
    :param output: The output of the test
    :param inputfile: The file with map
    :return: The score
    '''
    lines = []
    with open(inputfile, 'r') as file:
        lines = file.readlines()

    row, col = find_starting_point(lines, 'M')

    # 865432 coefs
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
    data['character'] = ('optimal ' + output.strip()[:3]) if adjusted else output.strip()[:3]
    data['time_average'] = np.mean(times)
    data['time_min'] = np.min(times)
    data['time_max'] = np.max(times)
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
    generated_filename = f'{PATH_TO_STORAGE}/{os.path.basename(filename)}_results.md'
    with open(generated_filename, 'w') as f:
        f.write(f'# A* algorithm performance for {os.path.basename(filename)} map\n\n')
        df.to_markdown(f, index=False)
        f.write(f'\n\n### Average time: {df["time_average"].mean():.6f}\n')
        f.write(f'### Average cost: {df["score"].mean():.2f}\n')
        f.write(f'### Heuristic coef: {HEURISTICS_COEFFICIENTS}\n')

    print(f'Generated markdown file: {generated_filename}')


def store_results_json(df: pd.DataFrame, filename: str):
    '''Store the results in a JSON file ready to be consumed by the frontend'''
    generated_filename = f'{PATH_TO_STORAGE}/{os.path.basename(filename)}_results.json'
    content = {
        "league": f"Open league {HEURISTICS_COEFFICIENTS}",
        "levels": [
            {
                "lvl": os.path.basename(filename),
                "map": open(filename, 'r').read().strip().split('\n'),
                "groups": [
                    {
                        "name": team['character'],
                        "status": "valid",
                        "score": int(team['score']),
                        "path": team['output'],
                    } for team in df.to_dict(orient='records')
                ]
            }
        ]
    }
    print(json.dumps(content, indent=4), file=open(generated_filename, 'w'))
    print(f'Generated JSON file: {generated_filename}')

def main():
    '''Main function'''
    args = arguments()
    validate_files(args.exec, args.inputfile)
    df = generate_samples(args.exec, args.inputfile)
    df = df.sort_values('score', ascending=True)
    store_dataframe(df, args.inputfile)
    store_results_json(df, args.inputfile)


if __name__ == '__main__':
    main()

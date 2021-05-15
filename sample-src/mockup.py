import argparse
parser = argparse.ArgumentParser(description = 'my description')
parser.add_argument('-n', '--name')
args = parser.parse_args()
print('hello ' + args.name)
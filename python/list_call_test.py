import fileinput

fp = fileinput.input()

for line in fp:
    print(line.strip())
#! /usr/bin/python3

import sys
from word2number import w2n

def check_number(text):
  for i in range(1,len(text)+1):
    try:
      w2n.word_to_num(text[:i])
      return w2n.word_to_num(text[:i]) 
    except ValueError:
      continue
    return None

if __name__ == "__main__":

  file_name = sys.argv[1]

  dat = None
  with open(file_name, "r") as f:
    dat = f.readlines()

  dat = [d.replace("\n", "") for d in dat]

  calib = 0
  new_line = []

  for line in dat:
    for i,l in enumerate(line):
      print(l, line[i:])
      number = check_number(line[i:]) 
      if l.isnumeric():
        new_line.append(l)
      elif number is not None:
        new_line.append(str(number))

    line = new_line
    new_line = []
    print(line)

    if len(line) == 1:
      calib += int(line[0] + line[0])
    else:
      calib += int(line[0] + line[-1])

  print(calib)

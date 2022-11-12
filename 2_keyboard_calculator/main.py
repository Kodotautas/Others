import time
import sys
import argparse
import os

cwd = os.getcwd()
print(cwd)
#set interval and session time in seconds in arguments
# parser = argparse.ArgumentParser()
# parser.add_argument("-i", "--interval", help="interval in seconds", type=int, default=10)
# parser.add_argument("-s", "--session", help="session time in seconds", type=int, default=20)

# args = parser.parse_args()
interval = 10
session = 30

#set path to keyboard events in logs folder
path = cwd + "\logs\keyboard_events.txt"
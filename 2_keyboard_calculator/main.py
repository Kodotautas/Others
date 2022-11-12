import msvcrt
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

#store keypresses in dictionary
def store_keypresses():
    """store keypresses in dictionary
    Returns:
        _type_: dictionary with keypresses
    """
    keypresses = {}
    while True:
        if msvcrt.kbhit():
            key = msvcrt.getch()
            if key in keypresses:
                keypresses[key] += 1
            else:
                keypresses[key] = 1
        return keypresses


#count how many keypresses in interval and store in dictionary
def count_keypresses(store_keypresses):
    """count how many keypresses in interval
    Returns:
        _type_: sum of keypresses in interval
    """
    keypresses = store_keypresses()
    start_time = time.time()
    while (time.time() - start_time) < interval:
        if msvcrt.kbhit():
            key = msvcrt.getch()
            if key in keypresses:
                keypresses[key] += 1
            else:
                keypresses[key] = 1
    #sum of keypresses and top 3 keypresses
    return sum(keypresses.values())

#take store_keypresses dictionary and return top n keypresses
def top_n_keypresses(store_keypresses, n):
    """return top n keypresses
    Returns:
        _type_: top n keypresses
    """
    keypresses = store_keypresses()
    sorted_keypresses = sorted(keypresses.items(), key=lambda x: x[1], reverse=True)
    return sorted_keypresses[:n]

def main():
    #start session
    start_time = time.time()
    print("Program started")
    while (time.time() - start_time) < session:
        keypresses = count_keypresses(store_keypresses)
        print("Keypresses in {} seconds interval: {}".format(interval, keypresses))
        top = top_n_keypresses(store_keypresses, 3)
        print("Top 3 keypresses in {} seconds interval: {}".format(interval, top))
    print("{} seconds session ended".format(session))

if __name__ == "__main__":
    main()
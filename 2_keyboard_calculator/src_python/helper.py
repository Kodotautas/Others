import time
import msvcrt
import argparse

#set interval and session time in seconds in arguments
parser = argparse.ArgumentParser()
parser.add_argument("-i", "--interval", help="interval in seconds", type=int, default=10)
parser.add_argument("-s", "--session", help="session time in seconds", type=int, default=30)
parser.add_argument("-n", "--top_n", help="top n keypresses", type=int, default=3)

args = parser.parse_args()
interval = args.interval
session = args.session
n = args.top_n

keypresses = {} #for keypresses in interval
keypress_total_dict = {} #for top n keypresses of the session
keypresses_total = [] #for otal count of keypresses
keypress_speed = [] #for keypresses speed


class Keys_Tracker:
    def __init__(self):
        self.keypresses = {}
        self.keypress_total_dict = {}
        self.start_time = time.time()
        self.interval = interval
        self.session = session

    #store keypresses in dictionary
    def store_keypresses(self):
        """store keypresses in dictionaries
        Returns:
            _type_: dictionary and total dictionary of keypresses
        """
        while True:
            if msvcrt.kbhit():
                key = msvcrt.getch()
                #track keypresses in interval and session
                if key in keypresses:
                    keypresses[key] += 1
                else:
                    keypresses[key] = 1
            return keypresses

    #store total keypresses in dictionary
    def store_total_keypresses(self, dictionary):
        """store total keypresses in dictionary
        Returns:
            _type_: total keypresses
        """
        for key, value in dictionary.items():
            if key in keypress_total_dict:
                keypress_total_dict[key] += value
            else:
                keypress_total_dict[key] = value
        return keypress_total_dict

    #count how many keypresses in interval and store in dictionary
    def count_keypresses(self, dictionary):
        """count how many keypresses in interval, append for speed calculation
        Returns:
            _type_: sum of keypresses in interval
        """
        start_time = time.time()
        while (time.time() - start_time) < interval:
            if msvcrt.kbhit():
                key = msvcrt.getch()
                if key in dictionary:
                    dictionary[key] += 1
                else:
                    dictionary[key] = 1
        sum_keypresses = sum(keypresses.values())
        keypresses_total.append(sum_keypresses)
        keypress_speed.append(sum_keypresses)
        return sum_keypresses

    #return top n keypresses from store_keypresses dictionary
    def top_n_keypresses(self, dictionary, n):
        """return top n keypresses
        Returns:
            _type_: top n keypresses
        """
        sorted_keypresses = sorted(dictionary.items(), key=lambda x: x[1], reverse=True)
        #clean up dictionary
        sorted_keypresses = [str(keypress[0]).replace("b'","") for keypress in sorted_keypresses]
        sorted_keypresses = [keypress.replace("'","") for keypress in sorted_keypresses]
        return sorted_keypresses[:n]


    #calculate average keypresses speed per minute
    def calculate_average_speed(self, keypress_speed):
        """calculate average keypresses speed per session
        Returns:
            _type_: average keypresses speed
        """
        speed = int(sum(keypress_speed) / (session/60))
        return speed
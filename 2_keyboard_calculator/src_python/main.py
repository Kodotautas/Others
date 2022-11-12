import time
import sys
import os
from helper import *

#main function
def main ():
    keys_tracker = Keys_Tracker()
    print(f'Started, keypresses will be tracked every {interval} seconds for {session} seconds.')
    print('')
    print('------------INTERVAL STATS----------')
    while (time.time() - keys_tracker.start_time) < keys_tracker.session:
        print(f'Keypress in {interval} seconds: ', keys_tracker.count_keypresses(keypresses))
        print(f'Top 3 keypresses: ', keys_tracker.top_n_keypresses(keypresses, n))
        print(f'Current session speed:', keys_tracker.calculate_average_speed(keypress_speed), 'characters per minute')
        keys_tracker.store_total_keypresses(keypresses)
        keypresses.clear()
        print('------------------------------------')
    
    print('')
    print('------------SESSION STATS----------')
    print(f'Total keypresses: ', sum(keypresses_total))
    print(f'Top 3 keypresses: ', keys_tracker.top_n_keypresses(keypress_total_dict, n))
    print(f'Final session speed:', keys_tracker.calculate_average_speed(keypress_speed), 'characters per minute')
    print('')

#if session less than interval, print error
if session < interval:
    print('Session time must be greater than interval time. Check it and try again.')
    sys.exit()
else:
    main()
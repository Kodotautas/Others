import os
import subprocess
import sys 

location = os.path.dirname(os.path.abspath(__file__))

#main function of the program
def main():
    while True:
        #interval input and check
        interval = input('Enter the interval in seconds: ')
        if not interval.isdigit() or int(interval) <= 0:
            print('Interval must be a number and greater than 0. Program exits.')
            sys.exit()
        #session input and check    
        session = input('Enter the session in seconds: ')
        if not session.isdigit() or int(session) <= 0:
            print('Session must be a number and greater than 0. Program exits.')
            sys.exit()
        #compare interval and session
        if session < interval:
            print('Session time must be greater than interval time. Program exits.')
        task = input("Enter 'start' / 'stop' to manage app: ")
        if task == 'start':
            subprocess.call(f'python calculator.py -i {interval} -s {session}', cwd=location)
            sys.exit()
        elif task == 'stop':
            sys.exit()
        else:
            print('Invalid input. Try again.')

if __name__ == '__main__':
    main()
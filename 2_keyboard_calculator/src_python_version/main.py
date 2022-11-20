import os
import subprocess
import sys 

location = os.path.dirname(os.path.abspath(__file__))


# ------------------------------- MAIN FUNCTION ------------------------------ #
def main():
    task = input("Enter 'start' / 'stop' to manage app: ")
    if task == 'start':
        while True:
            #interval input and check
            interval = input('Enter the interval in seconds: ')
            if not interval.isdigit() or int(interval) <= 0:
                print(f'Interval must be a number and greater than 0.\nStart / stop app to try again.')
                main()
            #session input and check    
            session = input('Enter the session in seconds: ')
            if not session.isdigit() or int(session) <= 0:
                print(f'Interval must be a number and greater than 0.\nStart / stop app to try again.')
                main()
            #compare interval and session
            elif int(interval) >= int(session):
                print(f'Interval must be less than session.\nStart / stop app to try again.')
                main()
            #run program with previous users inputs
            subprocess.call(f'python calculator.py -i {interval} -s {session}', cwd=location)
            sys.exit()
    elif task == 'stop':
        sys.exit('Program stopped.')
    else:
        print('Invalid input. Try again.')
        main()


if __name__ == '__main__':
    main()
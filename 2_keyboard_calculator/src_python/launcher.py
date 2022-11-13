import os
import subprocess 

location = os.path.dirname(os.path.abspath(__file__))

#main function of the program
def main():
    while True:
        interval = input('Enter the interval in seconds: ')
        session = input('Enter the session in seconds: ')
        task = input("Enter 'start' / 'stop' to manage app: ")
        if task == 'start':
            subprocess.call(f'python main.py -i {interval} -s {session}', cwd=location)
        elif task == 'stop':
            subprocess.call(['taskkill', '/F', '/IM', 'python.exe'])
        else:
            print('Invalid input. Try again.')

if __name__ == '__main__':
    main()
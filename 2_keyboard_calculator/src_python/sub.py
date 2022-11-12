import os
import subprocess 

location = os.path.dirname(os.path.abspath(__file__))

print(location)
print('-------------------------')
#main function
def main():
    while True:
        task = input("Enter 'start' / 'stop' to manage app: ")
        if task == 'start':
            #call main.py with subprocess
            subprocess.call(['python', 'main.py'], cwd=location)
        elif task == 'stop':
            subprocess.call(['taskkill', '/F', '/IM', 'python.exe'])
        else:
            print('Invalid input. Try again.')

if __name__ == '__main__':
    main()
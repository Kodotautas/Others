# Keypress tracker

Program tracks keypresses during the intervals of the session.

## More details
1. When program is started it tracks keypresses
2. In each interval of session it returns: how many keys pressed, what is top keys and characters per minute speed of session.
3. At session ending it returns 


## Main arguments
Arguments:
```Python
parser.add_argument("-i", "--interval", help="interval in seconds", type=int, default=10)
parser.add_argument("-s", "--session", help="session time in seconds", type=int, default=30)
parser.add_argument("-n", "--top_n", help="top n keypresses", type=int, default=3)

```

## How to use it?
### Lauch program with default arguments or set yours via CLI, like:
`"python main.py -i 20 -s 80 -n 5'"`
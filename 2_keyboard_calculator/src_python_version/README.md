# Keypress tracker

Program tracks keypresses during the intervals of the session.

## More details
1. When program is started (input = "start") it tracks keypresses (in background)

2. In each interval of session it returns: 
(a) how many keys pressed, 
(b) what is top keys and (c) characters per minute speed of session.

3. At session ending it returns the same parameters as during the interval.


## Main arguments:
```Python
parser.add_argument("-i", "--interval", help="interval in seconds", type=int, default=10)
parser.add_argument("-s", "--session", help="session time in seconds", type=int, default=30)
parser.add_argument("-n", "--top_n", help="top n keypresses", type=int, default=3)

```

## How to use it?
### Write your inputs after program is launched
`"py main.py'"`
### or set it via CLI, example:
`"py main.py -i 20 -s 80 -n 5"`




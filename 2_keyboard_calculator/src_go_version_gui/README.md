# Keypress tracker (Go version) adding GUI (using fyne.io)

Program tracks keypresses during the intervals of the session.

## More details
1. When program is started (input="start") it tracks keypresses. Need to press Enter after input(s).

2. After each interval of session it returns: 
(a) how many keys pressed, 
(b) what is top keys and (c) current speed of session (characters per minute).

3. At session ending it returns the same parameters as during the interval.

## Main parameters:
interval(int), session(int) - duration in seconds

## How to use it?
### Write your inputs in the code after launching program
`"go run ."`

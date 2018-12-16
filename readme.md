ws2js - Websockets to Javascript
========================================

Send javascript commands to a local webpage from any program.  Make
live graphics in your web browser from any programming language!


Current Status
------------------------------

~~~bash
myprogram | ws2js
~~~

Upon running this command, your browswer will open to
`localhost:8080`.  The output from `myprogram` will be read by `ws2js`
and run as javascript code in that webpage.

You might need to use `unbuffer` to prevent the unix pipe from holding
onto all of the commands:

~~~bash
unbuffer myprogram | ws2js
~~~





Goals
------------------------------

The original purpose is to provide an easy way to create live
graphics/animations to visualize algorithms or to run experiments. The
web browswer seemed like an ideal target because it's cross-platform
and there are a lot of easy-to-use graphics that can be dynamically
displayed with it.

A future goal might be to turn this into an online game client.  Since
it uses websockets, the javascript graphics directives could be sent
from a game server written in a language like python.  It might not be
the fastest way to do graphics, but it would be flexible for
prototypes.



How it Works
------------------------------

1. ws2js serves a webpage to `http://localhost:8080`
1. ws2js serves a websocket to `ws://localhost:8080/ws`
1. web browswer opens and goes to `localhost:8080`
1. ws2js reads user input from stdin.
1. ws2js ends user input to the websocket.
1. browser recieves text from the websocket.
1. browser runs the text as javascript.

~~~
user input --> stdin --> websockets --> webpage
~~~

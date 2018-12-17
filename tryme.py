"""
Example of ws2js Usage
==============================
"""

import random
import time

def out(string):
    s = string.replace('\n', ';')
    print(s)


createCanvas = '''
let canvas = document.createElement('canvas')
canvas.setAttribute('id', 'maincanvas')
canvas.style.border = 'solid 1px lightgray'
document.body.appendChild(canvas)
'''

context = '''
let ctx = document.querySelector('#maincanvas').getContext('2d')
'''

def randomColorHex():
    return "%03x" % random.randint(0, 0xFFF)

    
def randomCanvasBackground():
    s = "canvas.style.background = '#{}'"
    # s = "document.querySelector('#maincanvas').style.background = '#{}'"
    s = s.format(randomColorHex())
    return s


def main():
    for _ in range(100):
        out(randomCanvasBackground())
        time.sleep(0.2)
    
main()


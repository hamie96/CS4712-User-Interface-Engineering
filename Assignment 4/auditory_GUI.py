import tkinter as tk
import pyttsx3

def text_to_speech():
    engine = pyttsx3.init()
    engine.say(e1.get())
    engine.runAndWait()

master = tk.Tk()
tk.Label(master, text = "Text-To-Speech").grid(row=0)
e1 = tk.Entry(master)
e1.grid(row = 0, column = 0)
tk.Button(master, text = "Convert", command=text_to_speech).grid(row=0,column=3,sticky=tk.W,pady=4)

tk.mainloop()

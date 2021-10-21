# Divyanshu Lohani: https://github.com/DivyanshuLohani

from tkinter import *

import tkinter.ttk
root = Tk()
root.title("Calculator - Divyanshu")
root.resizable(False, False)

button_texts = "C-%÷789×456-123+0.d="

first_col = 'C_%÷'
second_col = '789×'
third_col = '456-'
fourth_col = '123+'
fifth_col = '0.d='

first_bg_color_set = ["#2c3338", "#2c3338", "#2c3338", "#725fcc"]
second_bg_color_set = ["#2c3338", "#2c3338", "#2c3338", "#725fcc"]
third_bg_color_set = ["#2c3338", "#2c3338", "#2c3338", "#725fcc"]
fourth_bg_color_set = ["#2c3338", "#2c3338", "#2c3338", "#725fcc"]
fifth_bg_color_set = ["#2c3338", "#2c3338", "#2c3338", "#fac40a"]

main_ans = ""
display_label_value = ""
history_label_value = ""

history_label = Label(root, text=history_label_value, height=2, bg="white", font="lucida 20", anchor='e')
history_label.pack(fill=X)
display_label = Label(root, text=history_label_value, height=2, width=2, pady=5,bg="white", font="lucida 30 bold", anchor='e')
display_label.pack(fill=X)
Frame(root, height=1, bg='white').pack(fill=X)
ends_with_op = True

def click(event):
    global display_label_value, history_label_value, main_ans, ends_with_op
    btn_text = event.widget.cget('text')
    # print(btn_text)
    
    if btn_text == "C":
        display_label_value = ""
        history_label_value = ""
        main_ans = ""
    elif btn_text == "-":
        if ends_with_op:
            return
        main_ans += display_label_value + "-"
        history_label_value = main_ans
        display_label_value = ""
        ends_with_op = True
    elif btn_text == "+":
        if ends_with_op:
            return
        main_ans += display_label_value + "+"
        history_label_value = main_ans
        display_label_value = ""
        ends_with_op = True
    elif btn_text == "÷":
        if ends_with_op:
            return
        main_ans += display_label_value + "/"
        history_label_value = main_ans.replace("/", "÷")
        display_label_value = ""
        ends_with_op = True
    elif btn_text == "×":
        if ends_with_op:
            return
        main_ans += display_label_value + "*"
        history_label_value = main_ans.replace("*", "×")
        display_label_value = ""
        ends_with_op = True
    elif btn_text == "CE":
        display_label_value = display_label_value[:-1]
    elif btn_text == "%":
        try:
            percent = ''
            if not percent == display_label_value:
                percent += display_label_value
            percent = percent + "/100"
            
            percent = str(eval(percent))
            display_label_value = percent
        except:
            
            pass
        finally:
            percent = ''
    elif btn_text == "=":
        main_ans += display_label_value
        if display_label_value == main_ans: return
        history_label_value += display_label_value
        display_label_value = ""
        try:
            main_ans = str(eval(main_ans))
            display_label_value = main_ans
            main_ans = ''
        except ZeroDivisionError as e:
            display_label_value = "Cannot Devide By Zero"
            main_ans = ''
        except SyntaxError as e:
            pass
    elif btn_text == ".":
        
        
        if display_label_value.find('.') == -1:
            if display_label_value == "":
                display_label_value = "0" + display_label_value + btn_text
            else:
                display_label_value += btn_text
    elif btn_text == "+/-":
        
        if display_label_value.startswith("-"):
        
            display_label_value = display_label_value.replace("-", "")
        else:
            display_label_value = "-" + display_label_value
    else:    
        display_label_value += btn_text
        ends_with_op = False
    display_label['text'] = display_label_value
    history_label['text'] = history_label_value
    # print(main_ans)



for row in range(5):
    f1 = Frame(root, height=55)
    # Setup the buttons
    for column in range(4):
        text = ""
        if row == 0:
            text = first_col
            bg_color = first_bg_color_set
        if row == 1:
            text = second_col
            bg_color = second_bg_color_set
        if row == 2:
            text = third_col
            bg_color = third_bg_color_set
        if row == 3:
            text = fourth_col
            bg_color = fourth_bg_color_set
        if row == 4:
            text = fifth_col
            bg_color = fifth_bg_color_set
        symbol = text[column]
        # Make the respective symbols 
        if symbol == "_": symbol = "+/-"
        elif symbol == "d": symbol = "CE"
        fg_color = "#dddddd"
        if row == 4 and column == 3:
            fg_color = "black"
        b1 = Button(f1, text=symbol, font="comic 20 bold", bg=bg_color[column], fg=fg_color,  padx=30, width=2, height=2, )
        b1.pack(side=LEFT)
        b1.bind("<Button 1>", click)


    f1.pack(fill=X)

root.mainloop()
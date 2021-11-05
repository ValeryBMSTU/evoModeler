import PySimpleGUI as sg
import requests

sg.theme("DarkTeal2")

def makeMainWindow():
    layout = [
        [sg.Text(text="url:", font=14),
         sg.In(size=(30, 1), enable_events=True, key="-URL-")],
        [sg.Combo(["POST", "GET", "PUT", "DELETE"], key="-METHOD-")],
        [sg.Button(button_text="Send request", size=(20, 1), key="-CLICK-"),
         sg.Text(text="", size=(30, 1), font=14, key="-SEND_REQUEST-")],
        [sg.Text(text="Info:", font=14),
         sg.Text(text="", size=(30, 1), font=14, key="-INFO-")]
    ]
    return sg.Window(title="evoModeler GUI", layout=layout, margins=(160, 90))

mainWindow = makeMainWindow()

while True:
    event, value = mainWindow.read()
    if event == "OK" or event == sg.WIN_CLOSED:
        mainWindow.close()
        print("Closing....")
        break
    elif event == "-CLICK-":
        url = mainWindow["-URL-"].get()  
        method = mainWindow["-METHOD-"].get()
        print(method)
        resp = ""
        if method == "GET":
            resp = requests.get(url)
        elif method == "POST":
            resp = requests.post(url)
        elif method == "PUT":
            resp = requests.put(url)
        elif method == "DELETE":
            resp = requests.delete(url)
        else:
            resp = requests.get(url)
            
        mainWindow["-INFO-"].update(\
            str(resp.status_code)+\
            str(resp.json()))



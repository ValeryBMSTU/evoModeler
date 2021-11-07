import PySimpleGUI as sg
import requests

sg.theme("DarkTeal2")

def makeMainWindow():
    layout = [
        [sg.Text(text="url:", size=(10, 1), font=14), sg.In(default_text="http://127.0.0.1:8080/ping", enable_events=True, key="-URL-")],
        [sg.Text(text="method:", size=(10, 1), font=14), sg.Combo(["POST", "GET", "PUT", "DELETE"], default_value="GET", key="-METHOD-")],
        [sg.Button(button_text="Send request", key="-CLICK-"),
         sg.Text(text="", font=14, key="-SEND_REQUEST-")],
        [sg.Text(text="Info:", font=14),
         sg.Text(text="", font=14, key="-INFO-")],
        [sg.Multiline(default_text='This is the default Text should you decide not to type anything', size=(100, 10), key="-LOG-")]
    ]
    return sg.Window(title="evoModeler GUI", layout=layout, margins=(20, 20))

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
            
        
        mainWindow["-LOG-"].update(mainWindow["-LOG-"].get()+"\n"+\
            str(resp.status_code)+\
            str(resp.json()))



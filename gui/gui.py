import PySimpleGUI as sg
import requests
import datetime

class User:
    def __init__(self, isAuth, sessionID):
        self.isAuth = isAuth
        self.sessionID = sessionID

class Server:
    def __init__(self, address):
        self.address = address

class State:
    def __init__(self, isAuth, sessionID, address):
        self.user = User(isAuth, sessionID)
        self.server = Server(address)

sg.theme("DarkTeal2")

def makeExpertWindow():
    layout = [
        [sg.Text(text="url:", size=(10, 1), font=14), sg.In(default_text="http://127.0.0.1:8080/ping", enable_events=True, key="-URL-")],
        [sg.Text(text="method:", size=(10, 1), font=14), sg.Combo(["POST", "GET", "PUT", "DELETE"], default_value="GET", key="-METHOD-")],
        [sg.Button(button_text="Send request", key="-CLICK-"),
         sg.Text(text="", font=14, key="-SEND_REQUEST-")],
         [sg.Button(button_text="Return to home page", key="-HOME-")],
        [sg.Multiline(default_text='This is the default Text should you decide not to type anything', size=(100, 10), key="-LOG-")]]
    return sg.Window(title="Expert Mode", layout=layout, margins=(20, 20))

def makeAuthWindow():
    layout = [
        [sg.Text(text="Login:", size=(10, 1), font=14), sg.In(enable_events=True, key="-LOGIN-")],
        [sg.Text(text="Password:", size=(10, 1), font=14), sg.In(enable_events=True, key="-PASS-")],
        [sg.Button(button_text="Log in", key="-LOG_IN-")],
        [sg.Button(button_text="Registration", key="-REG-")]]
    return sg.Window(title="Auth Window", layout=layout, margins=(20, 20))

def makeRegWindow():
    layout = [
        [sg.Text(text="Login:", size=(10, 1), font=14), sg.In(enable_events=True, key="-LOGIN-")],
        [sg.Text(text="Password:", size=(10, 1), font=14), sg.In(enable_events=True, key="-PASS-")],
        [sg.Text(text="Repeat Password:", size=(10, 1), font=14), sg.In(enable_events=True, key="-REP_PASS-")],
        [sg.Button(button_text="Sing Up", key="-SING_UP-")]]
    return sg.Window(title="Registration Window", layout=layout, margins=(20, 20))

def makeHomeWindow():
    left_col = [
        [sg.Button(button_text="Expert mode", key="-EXPERT-")],
        [sg.Button(button_text="Create new task", key="-CREATE-")]]
    right_col = [
        [sg.Text(text="Hi. I am here", font=14)]]
    layout = [
        [sg.Column(left_col, element_justification='c'),
         sg.VSeperator(),
         sg.Column(right_col, element_justification='c')]]
    return sg.Window(title="Home window", layout=layout, margins=(20, 20))

def makeCreateWindow():
    layout = [
        [sg.Text(text="Name:", size=(10, 1), font=14), sg.In(enable_events=True, key="-NAME-")],
        [sg.Text(text="Type:", size=(10, 1), font=14), sg.Combo(["Ant"], key="-TYPE-")],
        [sg.Button(button_text="Create", key="-CREATE-"), sg.Button(button_text="Cancel", key="-CANCEL-")]]
    return sg.Window(title="Create Window", layout=layout, margins=(20, 20))

def expertWindowProcess():
    expertWindow = makeExpertWindow()

    while True:
        event, value = expertWindow.read()
        if event == "OK" or event == "-HOME-" or event == sg.WIN_CLOSED:
            expertWindow.close()
            print("Closing expert window....")
            break
        elif event == "-CLICK-":
            url = expertWindow["-URL-"].get()
            method = expertWindow["-METHOD-"].get()
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

            print(resp.json())
            expertWindow["-LOG-"].update(expertWindow["-LOG-"].get()+"\n"+\
                str(datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S'))+" "+\
                str(resp.status_code)+" "+\
                str(resp.json()))

    expertWindow.close()

def regWindowProcess():
    regWindow = makeRegWindow()

    while True:
        event, value = regWindow.read()
        if event == "OK" or event == sg.WIN_CLOSED:
            regWindow.close()
            print("Closing reg window....")
            break
        elif event == "-SING_UP-":
             url = state.server.address + "singup?login=" + regWindow["-LOGIN-"].get() + "&pass=" + regWindow["-PASS-"].get()
             print(url)
             resp = requests.post(url)
             if resp.json()["meta"]["info"] == "OK":
                state.user.isAuth = True
                state.user.sessionID = resp.json()["data"]["session_id"]
                regWindow.close()
                print("Closing reg window....")
             else:
                print(resp.json()["meta"]["err"])

def homeWindowProcess():
    homeWindow = makeHomeWindow()

    while True:
        event, value = homeWindow.read()
        if event == "OK" or event == sg.WIN_CLOSED:
            homeWindow.close()
            state.user.isAuth = False
            state.user.sessionID = -1
            print("Closing home window....")
            break
        elif event == "-EXPERT-":
            homeWindow.Hide()
            expertWindowProcess()
            homeWindow.UnHide()
        elif event == "-CREATE-":
            homeWindow.Hide()
            createWindowProcess()
            homeWindow.UnHide()

def createWindowProcess():
    createWindow = makeCreateWindow()

    while True:
        event, value = createWindow.read()
        if event == "OK" or event == "-CANCEL-" or sg.WIN_CLOSED:
            createWindow.close()
            print("Closing create window....")
            break
        elif event == "-CREATE-": # заглушка
            createWindow.close()
            print("Closing create window....")
            break

state = State(False, -1, "http://127.0.0.1:8080/")
authWindow = makeAuthWindow()
while True:
    event, value = authWindow.read()
    if event == "OK" or event == sg.WIN_CLOSED:
        authWindow.close()
        print("Closing auth window....")
        break
    elif event == "-REG-":
        authWindow.Hide()
        regWindowProcess()
        authWindow.UnHide()
        if state.user.isAuth == True:
            authWindow.Hide()
            homeWindowProcess()
            authWindow.UnHide()
    elif event == "-LOG_IN-":
        url = state.server.address + "login?login=" + authWindow["-LOGIN-"].get() + "&pass=" + authWindow["-PASS-"].get()
        print(url)
        resp = requests.post(url)
        if resp.json()["meta"]["info"] == "OK":
            state.user.isAuth = True
            state.user.sessionID = resp.json()["data"]["session_id"]
            authWindow.Hide()
            homeWindowProcess()
            authWindow.UnHide()
        else:
            print(resp.json()["meta"]["err"])



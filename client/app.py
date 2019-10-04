from flask import Flask, render_template
from twilio.rest import Client
app = Flask(__name__)

@app.route('/')
def login():
    return render_template('login.html')

@app.route('/index')
def index():
    return render_template('index.html')

@app.route('/register')
def register():
    return render_template('register.html')

@app.route('/marketplace')
def marketplace():
    return render_template('marketplace.html')

@app.route('/devices')
def devices():
    return render_template('devices.html')

@app.route('/unlock')
def unlock():
    print("sljd")#return render_template('index.html')

@app.route('/lock')
def lock():
    print("locked")#return render_template('index.html')


if __name__ == '__main__':
    app.run()


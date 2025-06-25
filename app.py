from flask import Flask, make_response, redirect, request, render_template
import sqlite3
from dotenv import load_dotenv
import os
import base64
from hashlib import shake_256

load_dotenv()
DB_FILE = os.getenv("DB_FILE")
SERVER_PORT = os.getenv("SERVER_PORT")
FORWARD_DOMAIN = os.getenv("FORWARD_DOMAIN")

init_sql = """
CREATE TABLE IF NOT EXISTS shorts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    path TEXT NOT NULL UNIQUE,
    link TEXT NOT NULL
);
"""

sqlcon = sqlite3.connect(DB_FILE)
sqlcon.execute(init_sql)

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/get/<link>')
def getLink(link):
    sqlcon = sqlite3.connect(DB_FILE)
    cursor = sqlcon.cursor()
    cursor.execute("SELECT link FROM shorts WHERE path=?", (link,))
    data = cursor.fetchone()
    cursor.close()
    sqlcon.close()
    return redirect(data[0], code=302)

@app.route('/newlink', methods=["POST"])
def addLink():
    link = request.json['url']
    path = base64.b64encode(shake_256(link.encode()).digest(5)).decode()
    # path = shake_256(link["url"]).digest(3).decode()
    sqlcon = sqlite3.connect(DB_FILE)
    sqlcon.execute("INSERT INTO shorts (path, link) VALUES (?, ?)", (path, link))
    sqlcon.commit()
    sqlcon.close()
    return make_response({'status': 'success', 'shortUrl': FORWARD_DOMAIN+path}, 200)

app.run(port=int(SERVER_PORT))
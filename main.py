from flask import Flask, make_response, redirect, request
import sqlite3
from dotenv import load_dotenv
import os
import secrets
import string

def random_code(length: int) -> str:
    random_string = ''.join(secrets.choice(string.ascii_letters + string.digits) for _ in range(length))
    return random_string

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

@app.route('/get/<link>')
def getLink(link):
    sqlcon = sqlite3.connect(DB_FILE)
    cursor = sqlcon.cursor()
    cursor.execute("SELECT link FROM shorts WHERE path=?", (link,))
    data = cursor.fetchone()
    cursor.close()
    sqlcon.close()
    return redirect(data[0], code=302)

@app.route('/newlink')
def addLink():
    link = request.args.get("link")
    path = random_code(7)
    sqlcon = sqlite3.connect(DB_FILE)
    sqlcon.execute("INSERT INTO shorts (path, link) VALUES (?, ?)", (path, link))
    sqlcon.commit()
    sqlcon.close()
    return make_response({'status': 'success', 'url': FORWARD_DOMAIN+path}, 200)

app.run(port=int(SERVER_PORT))
from flask import Flask, make_response, redirect, request, render_template, abort
import sqlite3
from dotenv import load_dotenv
import os
import secrets
import string

def random_code(length: int) -> str:
    random_string = ''.join(secrets.choice(string.ascii_letters + string.digits) for _ in range(length))
    return random_string

def get_safe_random_code(cursor: sqlite3.Cursor, length: int) -> str:
    data = []
    while data is not None:
        code = random_code(length)
        cursor.execute('SELECT id FROM shorts WHERE path = ?', (code, ))
        data = cursor.fetchone()
    return code

load_dotenv()
DB_FILE = os.getenv("DB_FILE")
SERVER_PORT = os.getenv("SERVER_PORT")
TLS_CERT = os.getenv("TLS_CERT_FILE")
TLS_KEY = os.getenv("TLS_KEY_FILE")
ADDRESS_RESOLUTION = os.getenv("ADDRESS_RESOLUTION")

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

@app.errorhandler(404)
def pagenotfound(error):
    return render_template('wrong.html', code="404", message="Not Found", details="The following link maybe be broken"), 404

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/g/<link>')
def getLink(link):
    sqlcon = sqlite3.connect(DB_FILE)
    cursor = sqlcon.cursor()
    cursor.execute("SELECT link FROM shorts WHERE path=?", (link,))
    data = cursor.fetchone()
    if data is None:
        return pagenotfound(404)
    cursor.close()
    sqlcon.close()
    return redirect(data[0], code=302)

@app.route('/newlink', methods=["POST"])
def addLink():
    if 'url' not in request.json or not request.json['url']:
        return make_response({'status': 'failed', 'message': 'URL is empty'}, 500)
    sqlcon = sqlite3.connect(DB_FILE)
    link = str(request.json['url'])
    cursor = sqlcon.cursor()
    cursor.execute('SELECT path FROM shorts WHERE link = ?', (link, ))
    path = cursor.fetchone()
    if path is None:
        path = get_safe_random_code(cursor, 6)
        sqlcon.execute('INSERT INTO shorts (path, link) VALUES (?, ?)', (path, link))
        sqlcon.commit()
    else: path = path[0]
    cursor.close()
    sqlcon.close()
    return make_response({'status': 'success', 'shortUrl': path}, 200)

try:
    app.run(ssl_context=(TLS_CERT, TLS_KEY), host=ADDRESS_RESOLUTION, port=int(SERVER_PORT))
except:
    print('TLS data missing')
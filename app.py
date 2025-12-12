from flask import Flask, make_response, redirect, request, render_template
import sqlite3
from dotenv import load_dotenv
import os
import secrets
import string
import redis

def random_code(length: int) -> str:
    random_string = ''.join(secrets.choice(string.ascii_letters + string.digits) for _ in range(length))
    return random_string

def get_safe_random_code(cursor: sqlite3.Cursor, length: int) -> str:
    data = []
    code = ''
    while data is not None:
        code = random_code(length)
        cursor.execute('SELECT id FROM shorts WHERE path = ?', (code, ))
        data = cursor.fetchone()
    if code == '':
        raise ValueError('generated empty string')
    return code

load_dotenv()
DB_FILE = str(os.getenv("DB_FILE"))
SERVER_PORT = int(str(os.getenv("SERVER_PORT")))
ADDRESS_RESOLUTION = str(os.getenv("ADDRESS_RESOLUTION"))
REDIS_ADDRESS = str(os.getenv('REDIS_SERVER'))
REDIS_PORT = int(str(os.getenv('REDIS_PORT')))

r = redis.Redis(host=REDIS_ADDRESS, port=REDIS_PORT, db=0)

init_sql = """
CREATE TABLE IF NOT EXISTS shorts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    path TEXT NOT NULL UNIQUE,
    link TEXT NOT NULL
);
"""

sqlcon = sqlite3.connect(DB_FILE)
sqlcon.execute(init_sql)
sqlcon.close()

app = Flask(__name__)

@app.errorhandler(404)
def pagenotfound(error):
    return render_template('wrong.html', code="404", message=error, details="The following link maybe be broken"), 404

@app.route('/')
def index():
    return render_template('index.html')

def get_url_from_db(link: str) -> str | None:
    conn = sqlite3.connect(DB_FILE)
    cursor = conn.cursor()
    cursor.execute("SELECT link FROM shorts WHERE path=?", (link,))
    data = cursor.fetchone()
    cursor.close()
    conn.close()
    return data[0] if data else None

def get_url(link: str) -> str | None:
    key = f'url:{link}'

    cached = r.get(key)
    if cached and type(cached) == bytes:
        return cached.decode()

    url = get_url_from_db(link)
    if url != None:
        r.setex(key, 600, url)
    return url


@app.route('/<link>')
def getLink(link):
    redirect_url = get_url(link)
    if redirect_url == None:
        return pagenotfound(404)
    
    return redirect(redirect_url, code=302)

@app.route('/newlink', methods=["POST"])
def addLink():
    if request.json == None:
        return make_response({'status': 'failed', 'message': 'Invalid request'}, 400)
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

app.run(host=ADDRESS_RESOLUTION, port=SERVER_PORT)

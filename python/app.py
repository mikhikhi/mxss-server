from html_sanitizer import Sanitizer
from flask import Flask, request
sanitizer = Sanitizer()
app = Flask(__name__)

@app.route('/', methods=['POST'])
def index():
    return sanitizer.sanitize(request.form.get('js'))

app.run(host='0.0.0.0', port=8181)
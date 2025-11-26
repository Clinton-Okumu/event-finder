from flask import Blueprint, jsonify

welcome = Blueprint('welcome', __name__, url_prefix='/')

@welcome.route('/')
def index():
    return jsonify({'message': 'Welcome to the Event Finder API!'})



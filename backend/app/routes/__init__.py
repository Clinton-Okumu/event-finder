from .welcome import welcome
from .auth import auth


def register_routes(app):
    app.register_blueprint(welcome)
    app.register_blueprint(auth)

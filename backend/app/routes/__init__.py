from .welcome import welcome

def register_routes(app):
    app.register_blueprint(welcome)

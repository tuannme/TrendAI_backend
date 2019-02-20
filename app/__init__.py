from flask import Flask
from .core import db
from logging.config import dictConfig
from flask_jwt_extended import JWTManager

def create_app():
    """Create Flask app
    """

    # Config app logging
    dictConfig({
        'version': 1,
        'formatters': {'default': {
            'format': '[%(asctime)s] %(levelname)s in %(module)s: %(message)s',
        }},
        'handlers': {'wsgi': {
            'class': 'logging.StreamHandler',
            'stream': 'ext://flask.logging.wsgi_errors_stream',
            'formatter': 'default'
        }},
        'root': {
            'level': 'INFO',
            'handlers': ['wsgi']
        }
    })

    # Init app
    app = Flask(__name__)
    app.config.from_object('config')

    # Init database connection
    with app.app_context():
        db.init_app(app)
        db.get_db()

    # Init JWT
    JWTManager(app)

    # Default route
    @app.route('/')
    def default():
        return 'Do you see me?'

    # Register blueprints
    from .routers import auth, user
    app.register_blueprint(auth.bp, url_prefix='/auth')
    app.register_blueprint(user.bp, url_prefix='/user')

    return app

from flask import Blueprint, current_app as app
from ..core.db import get_db

bp = Blueprint('login', __name__)

@bp.route('/login', methods=('GET', 'POST'))
def login():
    db = get_db()
    app.logger.info("db instance")
    app.logger.info(db)
    app.logger.info(db.users.count())

    return 'Login page'

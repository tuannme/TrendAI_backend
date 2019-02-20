import click
from flask import current_app, g
from flask.cli import with_appcontext
from mongoengine import *


def get_db():
    """Connect to the application's configured database. The connection
    is unique for each request and will be reused if this is called
    again.
    """
    if 'db' not in g:
        g.db = connect(current_app.config['APP_NAME'], host=current_app.config['MONGO_URI'], alias='default')

    return g.db


def close_db(e=None):
    """If this request connected to the database, close the
    connection.
    """
    g.pop('db', None)


def init_db():
    """Init db instance"""
    get_db()


@click.command('init-db')
@with_appcontext
def init_db_command():
    """Command for init db instance."""
    init_db()
    click.echo('Initialized the database.')


def init_app(app):
    """Register database functions with the Flask app. This is called by
    the application factory.
    """
    app.teardown_appcontext(close_db)
    app.cli.add_command(init_db_command)

from flask import Blueprint, current_app as app, request, jsonify
import twitter
from twitter import error
from ..models.User import User
from ..models.ExternalUser import ExternalUser
import datetime
from mongoengine import errors

bp = Blueprint('login', __name__)


@bp.route('/login', methods=('GET', 'POST'))
def login():
    if request.json.get('access_token') is None:
        return jsonify({
            'error': True,
            'message': 'Missing access_token'
        }), 522

    if request.json.get('access_token_secret') is None:
        return jsonify({
            'error': True,
            'message': 'Missing access_token_secret'
        }), 522

    try:
        api = twitter.Api(consumer_key=app.config['TWITTER_CONSUMER_KEY'],
                          consumer_secret=app.config['TWITTER_CONSUMER_SECRET'],
                          access_token_key=request.json.get('access_token'),
                          access_token_secret=request.json.get('access_token_secret'), )

        twitter_user = api.VerifyCredentials(include_email=True)

        if twitter_user is None:
            return jsonify({
                'error': True,
                'message': 'Invalid credentials'
            }), 401

        twitter_user_dict = twitter_user.AsDict()
        external_user = ExternalUser(appId=app.config['TWITTER_CONSUMER_KEY'],
                                     userId=str(twitter_user_dict.get('id')))

        user = User()
        user.name = twitter_user_dict.get('name')
        user.email = twitter_user_dict.get('email')
        user.createdAt = datetime.datetime.utcnow()
        user.externalUsers = [external_user]
        user.save()

        return jsonify(twitter_user_dict)
    except error.TwitterError as ex:
        app.logger.error('Twitter authentication error %s' % ex)
        return jsonify({
            'error': True,
            'message': 'Unauthorized'
        }), 401
    except errors.ValidationError as ex:
        app.logger.error('Mongo document validation error %s' % ex)
        return jsonify({
            'error': True,
            'message': 'Unauthorized'
        }), 401

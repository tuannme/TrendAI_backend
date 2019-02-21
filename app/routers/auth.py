from flask import Blueprint, current_app as app, request, jsonify
import twitter
from twitter import error
from ..models.User import User
from ..models.ExternalUser import ExternalUser
from mongoengine.errors import ValidationError

bp = Blueprint('login', __name__)


@bp.route('/login', methods=['POST'])
def login():
    """Route login"""

    # Validate access_token
    if request.json.get('access_token') is None:
        return jsonify({
            'msg': 'Missing access_token'
        }), 522

    # Validate access_token_secret
    if request.json.get('access_token_secret') is None:
        return jsonify({
            'msg': 'Missing access_token_secret'
        }), 522

    try:
        # If all parameters passed the validation, try to connect to twitter
        api = twitter.Api(consumer_key=app.config['TWITTER_CONSUMER_KEY'],
                          consumer_secret=app.config['TWITTER_CONSUMER_SECRET'],
                          access_token_key=request.json.get('access_token'),
                          access_token_secret=request.json.get('access_token_secret'), )

        # Validate credentials
        twitter_user = api.VerifyCredentials(include_email=True)

        # Check credentials is valid
        if twitter_user is None:
            return jsonify({
                'msg': 'Invalid credentials'
            }), 401

        # Get twitter user's data as dictionary
        twitter_user_dict = twitter_user.AsDict()

        # External app's info
        external_app_id = app.config['TWITTER_APP_ID']
        external_user_id = str(twitter_user_dict.get('id'))
        external_user_name = twitter_user_dict.get('name')
        external_user_email = twitter_user_dict.get('email')
        external_user_lang = twitter_user_dict.get('lang')
        external_user_location = twitter_user_dict.get('location')
        external_user_followers_count = twitter_user_dict.get('followers_count')
        external_user_friends_count = twitter_user_dict.get('friends_count')
        external_user_statuses_count = twitter_user_dict.get('statuses_count')

        # Init external user
        external_user = ExternalUser(appId=external_app_id, userId=external_user_id)
        external_user.lang = external_user_lang
        external_user.location = external_user_location
        external_user.followers_count = external_user_followers_count
        external_user.friends_count = external_user_friends_count
        external_user.statuses_count = external_user_statuses_count

        # Get or create internal user
        try:
            user = User.objects.get(email=external_user_email)
        except User.DoesNotExist:
            user = User()

        # Check current external user is available
        if len(user.externalUsers) <= 0:
            user.externalUsers = [external_user]
        else:
            found = False
            for i, u in enumerate(user.externalUsers):
                if u.appId == external_app_id and u.userId == external_user_id:
                    found = True
                    # If external user was found in the list, update it
                    user.externalUsers[i] = external_user
                    break
            # If external user not found, add it to the list
            if found is False:
                user.externalUsers.append(external_user)

        # Assign external user's data to internal user's data
        user.name = external_user_name
        user.email = external_user_email
        user.save()

        return jsonify(user.auth_response())
    except ValidationError as ex:
        app.logger.error('Mongo document validation error %s' % ex)
        return jsonify({
            'msg': 'Unauthorized'
        }), 401
    except error.TwitterError as ex:
        app.logger.error('Twitter authentication error %s' % ex)
        return jsonify({
            'msg': 'Unauthorized'
        }), 401

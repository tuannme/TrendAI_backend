from functools import wraps
from flask import jsonify, current_app as app
from flask_jwt_extended import verify_jwt_in_request, get_jwt_identity
from ..models.User import User
from mongoengine.errors import ValidationError


# Custom decorator that verifies the JWT from request
def auth_required(fn):
    @wraps(fn)
    def wrapper(*args, **kwargs):
        verify_jwt_in_request()
        user_id = get_jwt_identity()
        try:
            # Get user from database
            user = User.get_by_pk(user_id)
            # Pass user data to next process
            return fn(user, *args, **kwargs)
        except User.DoesNotExist:
            return jsonify(msg='Unauthorized'), 401
        except ValidationError as ex:
            app.logger.error('Mongoengine validation error: ' + ex.message)
            return jsonify(msg='Unauthorized'), 401

    return wrapper

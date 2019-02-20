from flask import Blueprint, jsonify
from ..core.authentication import auth_required

bp = Blueprint('user', __name__)


@bp.route('/', methods=['GET'])
@auth_required
def get_user(user):
    """Get current user information"""
    return jsonify(user.response())

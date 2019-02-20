import datetime
from mongoengine import Document, StringField, ListField, DateTimeField, EmbeddedDocumentField
from .ExternalUser import ExternalUser
from flask_jwt_extended import create_access_token


class User(Document):
    name = StringField(required=True, max_length=255)
    email = StringField(required=True, unique=True, max_length=255)
    externalUsers = ListField(EmbeddedDocumentField(ExternalUser))
    createdAt = DateTimeField(default=datetime.datetime.utcnow)

    meta = {
        'indexes': [
            'email',
            'externalUsers.userId',
            'externalUsers.appId',
        ]
    }

    @staticmethod
    def get_by_pk(pk):
        """Get user by primary key"""
        return User.objects.get(pk=pk)

    def response(self):
        """Get user's data for response"""
        return {
            'id': str(self.id),
            'name': self.name,
            'email': self.email,
            'createdAt': self.createdAt
        }

    def auth_response(self):
        """Get user's data for response that need authentication information"""
        return {
            'user': self.response(),
            'token': {
                'access_token': create_access_token(identity=str(self.id))
            }
        }

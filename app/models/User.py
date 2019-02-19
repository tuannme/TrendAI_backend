from mongoengine import Document, StringField, ListField, DateTimeField, EmbeddedDocumentField
import datetime
from .ExternalUser import ExternalUser


class User(Document):
    name = StringField(required=True, max_length=255)
    #email = StringField(required=True, unique=True, max_length=255)
    externalUsers = ListField(EmbeddedDocumentField(ExternalUser))
    createdAt = DateTimeField(default=datetime.datetime.utcnow)

    meta={
        'indexes': [
            'externalUsers.userId',
            'externalUsers.appId',
        ]
    }

from mongoengine import EmbeddedDocument, StringField


class ExternalUser(EmbeddedDocument):
    appId = StringField(required=True, max_length=255)
    userId = StringField(required=True, max_length=255)

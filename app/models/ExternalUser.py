from mongoengine import EmbeddedDocument, StringField, IntField


class ExternalUser(EmbeddedDocument):
    appId = StringField(required=True, max_length=255)
    userId = StringField(required=True, max_length=255)
    lang = StringField(max_length=255)
    location = StringField(max_length=255)
    followers_count = IntField(min_value=0)
    friends_count = IntField(min_value=0)
    statuses_count = IntField(min_value=0)

import os

APP_NAME = os.getenv('APP_NAME', 'PythonApp')
FLASK_DEBUG = os.getenv('FLASK_DEBUG', False)
FLASK_ENV = os.getenv('FLASK_ENV', 'production')
MONGO_URI = os.getenv('MONGO_URI', 'mongodb://localhost:27017/trendAi')
TWITTER_APP_ID = os.getenv('TWITTER_APP_ID')
TWITTER_CONSUMER_KEY = os.getenv('TWITTER_CONSUMER_KEY')
TWITTER_CONSUMER_SECRET = os.getenv('TWITTER_CONSUMER_SECRET')

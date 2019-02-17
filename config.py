import os

FLASK_DEBUG = os.getenv('FLASK_DEBUG', False)
FLASK_ENV = os.getenv('FLASK_ENV', 'production')
MONGO_URI = os.getenv('MONGO_URI', 'mongodb://localhost:27017/trendAi')

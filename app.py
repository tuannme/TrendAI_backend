from dotenv import load_dotenv
from pathlib import Path  # Python3 only
from app import create_app

# Load environment variables from .env file
env_path = Path('.') / '.env'
load_dotenv(dotenv_path=env_path)

# Run your application
app = create_app()
app.run(host='0.0.0.0', port=5000)

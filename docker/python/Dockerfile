FROM python:3.7.2-stretch

# Make app directory
RUN mkdir /usr/src/app

# Change working directory
WORKDIR /usr/src/app

# Copy requirements to working directory
COPY ./requirements.txt .

# Install necessary packages
RUN pip install -r requirements.txt

# Copy source code to working directory
COPY . .

# Run your app
CMD python app.py

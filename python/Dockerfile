FROM python:3.13-rc-alpine
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
EXPOSE 8181
CMD ["python", "app.py"]
# syntax=docker/dockerfile:1.16
FROM python:3.13-alpine

COPY nap-dos/flaskserver.py .
COPY nap-dos/requirements.txt .

RUN pip3 install -r requirements.txt

EXPOSE 8000
CMD [ "gunicorn", "-w", "15", "-b", "0.0.0.0", "flaskserver:app", "--access-logfile", "-", "--error-logfile", "-"]

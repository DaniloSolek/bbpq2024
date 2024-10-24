FROM python:3.13

WORKDIR /app

COPY main.py /
COPY random-image.png /
COPY encrypted1.png /

RUN pip install opencv-python-headless numpy scipy
RUN apt-get update
RUN apt-get install -y libglib2.0-0 libgl1-mesa-glx

CMD ["python", "main.py"]
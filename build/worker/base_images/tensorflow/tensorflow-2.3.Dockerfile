FROM tensorflow/tensorflow:2.3.0

RUN apt update \
  && apt install -y libgl1-mesa-glx
COPY ./lib/requirements.txt /home
RUN pip install -r /home/requirements.txt

ENV PYTHONPATH "/home/lib"

WORKDIR /home/work
COPY ./lib /home/lib

ENTRYPOINT ["python"]

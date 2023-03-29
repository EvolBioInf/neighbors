FROM debian:stable-slim
RUN useradd -m -p NbqDBxZy0F.tE -s /bin/bash jdoe
RUN usermod -aG sudo jdoe
COPY install.sh .
COPY tutorial.txt /home/jdoe/
COPY README.md /home/jdoe/
RUN ["chmod", "+x", "./install.sh"]
RUN ./install.sh
RUN chown -R jdoe /home/jdoe/
RUN chgrp -R jdoe /home/jdoe/
WORKDIR /home/jdoe
USER jdoe
ENV HOME /home/jdoe
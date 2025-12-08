FROM postgis/postgis:17-3.4

USER root

RUN apt-get update && apt-get install -y \
    git \
    postgresql-server-dev-17 \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/pgpartman/pg_partman.git /tmp/pg_partman \
    && cd /tmp/pg_partman \
    && make && make install \
    && rm -rf /tmp/pg_partman

USER postgres

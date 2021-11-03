# Exmple of creating a container image that will run as a user 'mssql' instead of root
# This is example is based on the official image from Microsoft and effectively changes the user that SQL Server runs as
# and allows for dumps to generate as a non-root user


FROM mcr.microsoft.com/mssql/server

# Create non-root user and update permissions
#
RUN mkdir -p /opt/mssql-scripts
COPY script.sql /opt/mssql-scripts

ENV SA_PASSWORD=aB345678
ENV ACCEPT_EULA=Y
ENV MSSQL_PID=Express

RUN ( /opt/mssql/bin/sqlservr --accept-eula & sleep 20  ) | grep -q "Service Broker manager has started" \
    && /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P "aB345678" -i /opt/mssql-scripts/script.sql \
    && pkill sqlservr
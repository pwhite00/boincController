FROM ubuntu:16.04
MAINTAINER Peter White <samael00@gmail.com>

# Install the Boinc client
RUN apt-get update && apt-get install -y boinc-client

# Set the working directory
WORKDIR /var/lib/boinc-client

# Run boinc_client
CMD mkdir /var/lib/boinc-client/slots &&chown -R boinc /var/lib/boinc-client && chmod 775 /var/lib/boinc-client/slots && /etc/init.d/boinc-client start; sleep 15; /usr/bin/boinccmd --project_attach ${boincurl} ${boinckey}; tail -f /var/lib/boinc-client/std*.txt



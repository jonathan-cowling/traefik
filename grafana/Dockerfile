FROM grafana/grafana:6.2.5

COPY grafana.ini $WORKING_DIR/conf/custom.ini
COPY provisioning /etc/grafana/provisioning
COPY dashboards /var/lib/grafana/dashboards

ENV GF_PATHS_CONFIG="$WORKING_DIR/conf/custom.ini"
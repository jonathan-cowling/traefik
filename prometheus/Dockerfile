FROM prom/prometheus
COPY prometheus.yml /etc/prometheus/
ENTRYPOINT [ "prometheus", "--web.external-url=https://localhost:9090/prometheus/", "--config.file=/etc/prometheus/prometheus.yml" ]
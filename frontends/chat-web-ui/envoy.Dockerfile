FROM envoyproxy/envoy-dev:v1.25.0
COPY envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml
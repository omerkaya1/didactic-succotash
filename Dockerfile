FROM golang as dep_builder

FROM dep_builder as app_builder

FROM scratch

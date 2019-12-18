#!/usr/bin/env bash

swagger generate client --target ./servicios/camacho --name camachoAPIClient --spec ./servicios/camacho/swagger/swagger.yaml

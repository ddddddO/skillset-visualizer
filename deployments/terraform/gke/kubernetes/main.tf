provider kubernetes {
}

resource kubernetes_service sv_api_svc {
  metadata {
    name = "sv-api-svc"
    labels = {
      svc = "sv-api-svc"
    }
  }

  spec {
    type = "NodePort"
    selector = {
      api = "skillset-visualizer"
    }

    port {
      port        = 8081
      target_port = 8081
    }
  }
}

resource kubernetes_pod sv_api_db {
  metadata {
    name = "sv-api-db"
    labels = {
      api = "skillset-visualizer"
    }
  }
  spec {
    container {
      name  = "sv-db"
      image = "us.gcr.io/work1111/sv-db"
      port {
        container_port = 5432
      }
    }
    container {
      name  = "sv-api"
      image = "us.gcr.io/work1111/sv-api"
      port {
        container_port = 8081
      }
      command = ["ash", "-c", "sleep 10; ./api"]
      readiness_probe {
        http_get {
          path = "/health"
          port = 8081
        }
        initial_delay_seconds = 15
        period_seconds        = 120
        timeout_seconds       = 20
      }
    }
  }
}


# ref: https://www.terraform.io/docs/providers/kubernetes/r/service.html
resource kubernetes_service sv_app_svc {
  metadata {
    name = "sv-app-svc"
    labels = {
      svc = "sv-app-svc"
    }
  }

  spec {
    type = "NodePort"
    selector = {
      app = "skillset-visualizer"
    }

    port {
      port        = 80
      target_port = 8080
    }
  }
}

resource kubernetes_pod sv_app {
  metadata {
    name = "sv-app"
    labels = {
      app = "skillset-visualizer"
    }
  }
  spec {
    container {
      name  = "sv-app"
      image = "us.gcr.io/work1111/sv-app"
      port {
        container_port = 8080
      }
      command = ["ash", "-c", "sleep 10; http-server dist"]
    }
  }
}

# ref: https://www.terraform.io/docs/providers/kubernetes/r/ingress.html
resource kubernetes_ingress sv_app_svc_ingress {
  metadata {
    name = "sv-app-svc-ingress"
    annotations = {
      "kubernetes.io/ingress.global-static-ip-name" = data.google_compute_global_address.dododo_ip.name
    }
  }

  spec {
    rule {
      host = "dododo.site"
      http {
        path {
          path = "/*"
          backend {
            service_name = "${kubernetes_service.sv_app_svc.metadata.0.labels.svc}"
            service_port = 80
          }
        }
        path {
          path = "/fetch"
          backend {
            service_name = "${kubernetes_service.sv_api_svc.metadata.0.labels.svc}"
            service_port = 8081
          }
        }
        path {
          path = "/put/*"
          backend {
            service_name = "${kubernetes_service.sv_api_svc.metadata.0.labels.svc}"
            service_port = 8081
          }
        }
      }
    }
  }
}


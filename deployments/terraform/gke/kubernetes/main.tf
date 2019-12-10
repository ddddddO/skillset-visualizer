provider kubernetes {
}

resource kubernetes_pod sv_all_in_one {
  metadata {
    name = "sv-all-in-one"
    labels = {
      app = "skillset-visualizer"
    }
  }

  spec {
    container {
      name = "sv-db"
      image = "us.gcr.io/work1111/sv-db"
      port {
        container_port = 5432
      }
    }
    container {
      name = "sv-api"
      image = "us.gcr.io/work1111/sv-api"
      port {
        container_port = 8081
      }
      command = ["ash", "-c", "sleep 10; ./api"]
    }
    container {
      name = "sv-app"
      image = "us.gcr.io/work1111/sv-app"
      port {
        container_port = 8080
      }
      command = ["ash", "-c", "sleep 10; http-server dist"]
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
      #app = "${kubernetes_pod.sv_all_in_one.metadata.0.labels.app}" NOTE: This map does not have an element with the key "app".
      app = "skillset-visualizer"
    }

    port {
      port = 80
      target_port = 8080
    }
  }
}

# ref: https://www.terraform.io/docs/providers/kubernetes/r/ingress.html
resource kubernetes_ingress sv_app_svc_ingress {
  metadata {
    name = "sv-app-svc-ingress"
    annotations = {
      "kubernetes.io/ingress.global-static-ip-name" = "work-skillset-visualizer"
    }
  }

  spec {
    rule {
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
      port = 8081
      target_port = 8081
    }
  }
}

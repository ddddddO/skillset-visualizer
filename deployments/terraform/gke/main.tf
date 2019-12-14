provider google {
  credentials = "${file("~/.config/gcloud/legacy_credentials/lbfdeatq@gmail.com/adc.json")}"

  project = "work1111"
  region  = "us-central1"
  zone    = "us-central1-a"
}

# ネットワーク
resource google_compute_network work_net {
  name = "work-net"
}

# ファイアウォール
resource google_compute_firewall work_net_fw {
  name = "work-net-fw"
  network = "${google_compute_network.work_net.name}"

  allow {
    protocol = "tcp"
    ports = ["80", "8081"]
  }
}

resource google_container_cluster gcc {
  name = "work-gcc"
  location = "us-central1"
  network = "${google_compute_network.work_net.self_link}" 

  remove_default_node_pool = true
  initial_node_count = 1

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = true
    }
  }
}

resource google_container_node_pool gc_node_pool {
  name = "gc-node-pool"
  location   = "us-central1"
  cluster = "${google_container_cluster.gcc.name}"
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"

    metadata = {
      disable-legacy-endpoints = "true"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/devstorage.read_only",
    ]
  }
}

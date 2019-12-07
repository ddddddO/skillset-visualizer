provider google {
  # 認証情報
  credentials = "${file("~/.config/gcloud/legacy_credentials/lbfdeatq@gmail.com/adc.json")}"  

  project = "work1111" # 一旦
  region  = "us-central1"
  zone    = "us-central1-a"
}

# ネットワーク
resource google_compute_network work_net {
  name = "work-net"
}

# VM
resource google_compute_instance fst_vm {
  name = "fst-vm-1"

  # マシンタイプ
  machine_type = "n1-standard-1" # NOTE: 徐々に性能を落としていく

  # OS
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  # vmが存在するネットワークの指定
  network_interface {
    network = "${google_compute_network.work_net.self_link}"
  }
}

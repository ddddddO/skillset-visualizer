provider google {
  credentials = "${file("~/.config/gcloud/legacy_credentials/lbfdeatq@gmail.com/adc.json")}"

  project = "work1111"
}

# 手動登録済みのリソースを参照する
data google_compute_global_address dododo_ip {
  name = "work-skillset-visualizer"
}

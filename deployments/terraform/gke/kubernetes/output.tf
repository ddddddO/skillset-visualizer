# 確認用
output my_global_ip {
  value = data.google_compute_global_address.dododo_ip
  description = "For confirming my global ip (dododo ip)"
}

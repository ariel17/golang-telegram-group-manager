output "host_public_ip" {
  description = "Host public IP"
  value = aws_instance.bot_server[0].public_ip
}

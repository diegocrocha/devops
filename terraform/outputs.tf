output "instance_public_ip" {
  description = "IP público da instância EC2"
  value       = aws_instance.ec2.public_ip
}

output "private_key_path" {
  description = "Caminho da chave privada gerada"
  value       = local_file.private_key.filename
}

output "kafka_public_ip" {
  description = "IP público do servidor Kafka"
  value       = aws_instance.kafka.public_ip
}

output "kafka_bootstrap_servers" {
  description = "Endpoint para conexão com o Kafka"
  value       = "${aws_instance.kafka.public_ip}:9092"
}

resource "aws_instance" "ec2" {
  ami                         = "ami-0c02fb55956c7d316"
  instance_type               = var.instance_type
  subnet_id                   = aws_subnet.public.id
  vpc_security_group_ids      = [aws_security_group.ec2_sg.id]
  key_name                    = aws_key_pair.generated_key.key_name
  associate_public_ip_address = true

  tags = {
    Name = "ec2-linux"
  }
}

resource "aws_instance" "kafka" {
  ami                         = "ami-0c02fb55956c7d316"
  instance_type               = var.instance_type
  subnet_id                   = aws_subnet.public.id
  vpc_security_group_ids      = [aws_security_group.ec2_sg.id]
  key_name                    = aws_key_pair.generated_key.key_name
  associate_public_ip_address = true

  tags = {
    Name = "kafka-broker"
  }
}

resource "aws_security_group" "main" {
  egress = [
    {
      cidr_blocks      = ["0.0.0.0/0", ]
      description      = ""
      from_port        = 0
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "-1"
      security_groups  = []
      self             = false
      to_port          = 0
    }
  ]
  ingress = [
    {
      cidr_blocks      = ["0.0.0.0/0", ]
      description      = ""
      from_port        = 22
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
      to_port          = 22
    }
  ]
}

resource "aws_instance" "bot_server" {
  ami                    = var.AWS_AMI  # See available AMIs here: https://aws.amazon.com/amazon-linux-ami/
  instance_type          = "t2.nano"
  count                  = 1
  key_name               = var.PRIVATE_KEY_NAME
  vpc_security_group_ids = [aws_security_group.main.id]

  connection {
    type        = "ssh"
    user        = var.REMOTE_USER 
    private_key = file(var.PRIVATE_KEY_PATH)
    host        = self.public_ip
  }

  provisioner "file" {
    source      = "../.env"
    destination = "/home/${var.REMOTE_USER}/.env"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install docker -y",
      "sudo service docker start",
      "sudo docker run --pull=always --env-file=/home/${var.REMOTE_USER}/.env -d ariel17/golang-telegram-group-manager:latest"
    ]
  }

  tags = {
    name = "bot-server"
  }
}

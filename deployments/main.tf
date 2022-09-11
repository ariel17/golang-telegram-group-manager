terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.63.0"
    }
  }

  required_version = "~> 1.0.5"
}

provider "aws" {
  profile = "ariel17"
  region  = "us-east-2"
}

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
  ami                    = "ami-0b59bfac6be064b78" # See available AMIs here: https://aws.amazon.com/amazon-linux-ami/
  instance_type          = "t2.nano"
  count                  = 1
  key_name               = "bot-server"
  vpc_security_group_ids = [aws_security_group.main.id]

  connection {
    type        = "ssh"
    user        = "ec2-user"
    private_key = file("~/.ssh/bot-server.pem")
    host        = self.public_ip
  }

  provisioner "file" {
    source      = "../.env"
    destination = "/home/ec2-user/.env"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install docker -y",
      "sudo service docker start",
      "sudo docker run --pull=always --env-file=/home/ec2-user/.env -d ariel17/golang-telegram-group-manager:latest"
    ]
  }

  tags = {
    name = "bot-server"
  }
}

resource "aws_cloudwatch_log_group" "auth" {
  name = "${var.default_tags.project}-auth-logs"
}

resource "aws_cloudwatch_log_group" "file" {
  name = "${var.default_tags.project}-file-logs"
}

resource "aws_cloudwatch_log_group" "notification" {
  name = "${var.default_tags.project}-notification-logs"
}

resource "aws_cloudwatch_log_group" "post" {
  name = "${var.default_tags.project}-post-logs"
}

resource "aws_cloudwatch_log_group" "user" {
  name = "${var.default_tags.project}-user-logs"
}

resource "aws_cloudwatch_log_group" "router" {
  name = "${var.default_tags.project}-router-logs"
}

locals {
  auth_logs_configuration = {
    logDriver = "awslogs"
    options = {
      awslogs-group         = aws_cloudwatch_log_group.auth.name
      awslogs-region        = var.region
      awslogs-stream-prefix = "${var.default_tags.project}-auth"
    }
  }
  file_logs_configuration = {
    logDriver = "awslogs"
    options = {
      awslogs-group         = aws_cloudwatch_log_group.file.name
      awslogs-region        = var.region
      awslogs-stream-prefix = "${var.default_tags.project}-file"
    }
  }
  notification_logs_configuration = {
    logDriver = "awslogs"
    options = {
      awslogs-group         = aws_cloudwatch_log_group.notification.name
      awslogs-region        = var.region
      awslogs-stream-prefix = "${var.default_tags.project}-notification"
    }
  }
  post_logs_configuration = {
    logDriver = "awslogs"
    options = {
      awslogs-group         = aws_cloudwatch_log_group.post.name
      awslogs-region        = var.region
      awslogs-stream-prefix = "${var.default_tags.project}-post"
    }
  }
  user_logs_configuration = {
    logDriver = "awslogs"
    options = {
      awslogs-group         = aws_cloudwatch_log_group.user.name
      awslogs-region        = var.region
      awslogs-stream-prefix = "${var.default_tags.project}-user"
    }
  }
  router_logs_configuration = {
    logDriver = "awslogs"
    options = {
      awslogs-group         = aws_cloudwatch_log_group.router.name
      awslogs-region        = var.region
      awslogs-stream-prefix = "${var.default_tags.project}-router"
    }
  }
}
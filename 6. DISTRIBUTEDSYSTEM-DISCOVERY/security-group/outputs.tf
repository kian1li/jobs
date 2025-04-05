output "sg_apollo_router_service" {
  value       = aws_security_group.sg_apollo_router_service
  description = "Security Groupo of router loadbalancer"
}

output "sg_apollo_router_alb" {
  value       = aws_security_group.sg_apollo_router_alb
  description = "Security Groupo of router loadbalancer"
}

output "sg_ecs_auth_service" {
  value       = aws_security_group.sg_ecs_auth_service
  description = "Security Groupo of elastic container service Authorization"
}

output "sg_ecs_file_service" {
  value       = aws_security_group.sg_ecs_file_service
  description = "Security Groupo of elastic container service File"
}

output "sg_ecs_notification_service" {
  value       = aws_security_group.sg_ecs_notification_service
  description = "Security Groupo of elastic container service Notification"
}

output "sg_ecs_post_service" {
  value       = aws_security_group.sg_ecs_post_service
  description = "Security Groupo of elastic container service Post"
}

output "sg_ecs_user_service" {
  value       = aws_security_group.sg_ecs_user_service
  description = "Security Groupo of elastic container service User"
}
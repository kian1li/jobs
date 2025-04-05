resource "aws_route53_record" "router_alias" {
    zone_id = data.aws_route53_zone.micro.zone_id
    name = "router.micro.dev"
    type = "A"

    alias {
      name = data.aws_lb.router_alb.dns_name
      zone_id = data.aws_lb.router_alb.zone_id
      evaluate_target_health = true
    }
}
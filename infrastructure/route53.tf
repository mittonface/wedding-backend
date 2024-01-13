data "aws_route53_zone" "mydomain" {
 name = "brent.click"
}

resource "aws_route53_record" "wedding_backend" {
 zone_id = data.aws_route53_zone.mydomain.zone_id
 name   = "wedding-backend"
 type   = "A"

 alias {
   name                 = aws_alb.application_load_balancer.dns_name
   zone_id               = aws_alb.application_load_balancer.zone_id
   evaluate_target_health = true
 }
}

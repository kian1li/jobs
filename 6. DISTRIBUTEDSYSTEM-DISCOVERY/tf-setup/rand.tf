resource "random_id" "id" {
  byte_length = 8
}

output "tfid" {
  value = random_id.id.hex
}
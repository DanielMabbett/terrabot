resource "null_resource" "test_resource_1" {
  provisioner "local-exec" {
    command = "echo Nothing here to do"
  }
}

resource "null_resource" "test_resource_2" {
  provisioner "local-exec" {
    command = "echo Nothing here to do"
  }
}

resource "null_resource" "test_resource_1" {
  # Triggers can be used to detect changes in the configuration
  triggers = {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "echo The null resource was executed at $(date)"
  }
}

resource "null_resource" "test_resource_2" {
  # Triggers can be used to detect changes in the configuration
  triggers = {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "echo The null resource was executed at $(date)"
  }
}

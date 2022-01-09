resource "commercetools_subscription" "commercetools_category_subscription" {
  key = "commercetools-category-subscription"

  destination = {
    type       = "google_pubsub"
    project_id = var.gcp_project_id
    topic      = var.gcp_category_topic
  }

  changes {
    resource_type_ids = ["category"]
  }
}
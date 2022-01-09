terraform {
  required_providers {
    commercetools = {
      source = "labd/commercetools"
    }
  }
}

provider "google" {
  project = var.gcp_project_id
  region  = var.gcp_region
}

provider "commercetools" {
  client_id     = var.commercetools_client_id
  client_secret = var.commercetools_client_secret
  project_key   = var.commercetools_project_key
  scopes        = var.commercetools_scopes
  api_url       = var.commercetools_api_url
  token_url     = var.commercetools_token_url
}
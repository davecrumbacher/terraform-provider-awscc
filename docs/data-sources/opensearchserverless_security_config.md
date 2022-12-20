---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_opensearchserverless_security_config Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::OpenSearchServerless::SecurityConfig
---

# awscc_opensearchserverless_security_config (Data Source)

Data Source schema for AWS::OpenSearchServerless::SecurityConfig



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `description` (String) Security config description
- `name` (String) The friendly name of the security config
- `saml_options` (Attributes) Describes saml options in form of key value map (see [below for nested schema](#nestedatt--saml_options))
- `type` (String) Config type for security config

<a id="nestedatt--saml_options"></a>
### Nested Schema for `saml_options`

Read-Only:

- `group_attribute` (String) Group attribute for this saml integration
- `metadata` (String) The XML saml provider metadata document that you want to use
- `session_timeout` (Number) Defines the session timeout in minutes
- `user_attribute` (String) Custom attribute for this saml integration


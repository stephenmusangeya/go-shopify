package access

// APIScope represents the available Shopify API scopes
type APIScope string

const (
	// APIScopeReadAllOrders read all orders.
	APIScopeReadAllOrders APIScope = "read_all_orders"
	//  Read assigned fulfillment orders.
	APIScopeReadAssignedFulfillmentOrders APIScope = "read_assigned_fulfillment_orders"
	//  Write assigned fulfillment orders.
	APIScopeWriteAssignedFulfillmentOrders APIScope = "write_assigned_fulfillment_orders"
	//  Read checkouts.
	APIScopeReadCheckouts APIScope = "read_checkouts"
	//  Write checkouts.
	APIScopeWriteCheckouts APIScope = "write_checkouts"
	//  Read content.
	APIScopeReadContent APIScope = "read_content"
	//  Write content.
	APIScopeWriteContent APIScope = "write_content"
	//  Read customer merge.
	APIScopeReadCustomerMerge APIScope = "read_customer_merge"
	//  Write customer merge.
	APIScopeWriteCustomerMerge APIScope = "write_customer_merge"
	//  Read customers.
	APIScopeReadCustomers APIScope = "read_customers"
	//  Write customers.
	APIScopeWriteCustomers APIScope = "write_customers"
	//  Read customer payment methods.
	APIScopeReadCustomerPaymentMethods APIScope = "read_customer_payment_methods"
	//  Read discounts.
	APIScopeReadDiscounts APIScope = "read_discounts"
	//  Write discounts.
	APIScopeWriteDiscounts APIScope = "write_discounts"
	//  Read draft orders.
	APIScopeReadDraftOrders APIScope = "read_draft_orders"
	//  Write draft orders.
	APIScopeWriteDraftOrders APIScope = "write_draft_orders"
	//  Read files.
	APIScopeReadFiles APIScope = "read_files"
	//  Write files.
	APIScopeWriteFiles APIScope = "write_files"
	//  Read fulfillments.
	APIScopeReadFulfillments APIScope = "read_fulfillments"
	//  Write fulfillments.
	APIScopeWriteFulfillments APIScope = "write_fulfillments"
	//  Read gift cards.
	APIScopeReadGiftCards APIScope = "read_gift_cards"
	//  Write gift cards.
	APIScopeWriteGiftCards APIScope = "write_gift_cards"
	//  Read inventory.
	APIScopeReadInventory APIScope = "read_inventory"
	//  Write inventory.
	APIScopeWriteInventory APIScope = "write_inventory"
	//  Read legal policies.
	APIScopeReadLegalPolicies APIScope = "read_legal_policies"
	//  Read locales.
	APIScopeReadLocales APIScope = "read_locales"
	//  Write locales.
	APIScopeWriteLocales APIScope = "write_locales"
	//  Read locations.
	APIScopeReadLocations APIScope = "read_locations"
	//  Read metaobject definitions.
	APIScopeReadMetaobjectDefinitions APIScope = "read_metaobject_definitions"
	//  Write metaobject definitions.
	APIScopeWriteMetaobjectDefinitions APIScope = "write_metaobject_definitions"
	//  Read metaobjects.
	APIScopeReadMetaobjects APIScope = "read_metaobjects"
	//  Write metaobjects.
	APIScopeWriteMetaobjects APIScope = "write_metaobjects"
	//  Read marketing events.
	APIScopeReadMarketingEvents APIScope = "read_marketing_events"
	//  Write marketing events.
	APIScopeWriteMarketingEvents APIScope = "write_marketing_events"
	//  Read merchant approval signals.
	APIScopeReadMerchantApprovalSignals APIScope = "read_merchant_approval_signals"
	//  Read merchant managed fulfillment orders.
	APIScopeReadMerchantManagedFulfillmentOrders APIScope = "read_merchant_managed_fulfillment_orders"
	//  Write merchant managed fulfillment orders.
	APIScopeWriteMerchantManagedFulfillmentOrders APIScope = "write_merchant_managed_fulfillment_orders"
	//  Read orders.
	APIScopeReadOrders APIScope = "read_orders"
	//  Write orders.
	APIScopeWriteOrders APIScope = "write_orders"
	//  Read payment mandate.
	APIScopeReadPaymentMandate APIScope = "read_payment_mandate"
	//  Write payment mandate.
	APIScopeWritePaymentMandate APIScope = "write_payment_mandate"
	//  Read payment terms.
	APIScopeReadPaymentTerms APIScope = "read_payment_terms"
	//  Write payment terms.
	APIScopeWritePaymentTerms APIScope = "write_payment_terms"
	//  Read price rules.
	APIScopeReadPriceRules APIScope = "read_price_rules"
	//  Write price rules.
	APIScopeWritePriceRules APIScope = "write_price_rules"
	//  Read products.
	APIScopeReadProducts APIScope = "read_products"
	//  Write products.
	APIScopeWriteProducts APIScope = "write_products"
	//  Read product listings.
	APIScopeReadProductListings APIScope = "read_product_listings"
	//  Read publications.
	APIScopeReadPublications APIScope = "read_publications"
	//  Write publications.
	APIScopeWritePublications APIScope = "write_publications"
	//  Read purchase options.
	APIScopeReadPurchaseOptions APIScope = "read_purchase_options"
	//  Write purchase options.
	APIScopeWritePurchaseOptions APIScope = "write_purchase_options"
	//  Read reports.
	APIScopeReadReports APIScope = "read_reports"
	//  Write reports.
	APIScopeWriteReports APIScope = "write_reports"
	//  Read resource feedbacks.
	APIScopeReadResourceFeedbacks APIScope = "read_resource_feedbacks"
	//  Write resource feedbacks.
	APIScopeWriteResourceFeedbacks APIScope = "write_resource_feedbacks"
	//  Read script tags.
	APIScopeReadScriptTags APIScope = "read_script_tags"
	//  Write script tags.
	APIScopeWriteScriptTags APIScope = "write_script_tags"
	//  Read shipping.
	APIScopeReadShipping APIScope = "read_shipping"
	//  Write shipping.
	APIScopeWriteShipping APIScope = "write_shipping"
	//  Read Shopify Payments disputes.
	APIScopeReadShopifyPaymentsDisputes APIScope = "read_shopify_payments_disputes"
	//  Read Shopify Payments payouts.
	APIScopeReadShopifyPaymentsPayouts APIScope = "read_shopify_payments_payouts"
	//  Read own subscription contracts.
	APIScopeReadOwnSubscriptionContracts APIScope = "read_own_subscription_contracts"
	//  Write own subscription contracts.
	APIScopeWriteOwnSubscriptionContracts APIScope = "write_own_subscription_contracts"
	//  Read returns.
	APIScopeReadReturns APIScope = "read_returns"
	//  Write returns.
	APIScopeWriteReturns APIScope = "write_returns"
	//  Read themes.
	APIScopeReadThemes APIScope = "read_themes"
	//  Write themes.
	APIScopeWriteThemes APIScope = "write_themes"
	//  Read translations.
	APIScopeReadTranslations APIScope = "read_translations"
	//  Write translations.
	APIScopeWriteTranslations APIScope = "write_translations"
	//  Read third-party fulfillment orders.
	APIScopeReadThirdPartyFulfillmentOrders APIScope = "read_third_party_fulfillment_orders"
	//  Write third-party fulfillment orders.
	APIScopeWriteThirdPartyFulfillmentOrders APIScope = "write_third_party_fulfillment_orders"
	//  Read users.
	APIScopeReadUsers APIScope = "read_users"
	//  Read order edits.
	APIScopeReadOrderEdits APIScope = "read_order_edits"
	//  Write order edits.
	APIScopeWriteOrderEdits APIScope = "write_order_edits"
	//  Write payment gateways.
	APIScopeWritePaymentGateways APIScope = "write_payment_gateways"
	//  Write payment sessions.
	APIScopeWritePaymentSessions APIScope = "write_payment_sessions"
)

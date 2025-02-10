package constants

const MAX_DESCRIPTION_LEN int = 50
const TREASURY_API_URL string = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=record_date,country,currency,exchange_rate&filter=record_date:gt:%s,country:eq:%s"

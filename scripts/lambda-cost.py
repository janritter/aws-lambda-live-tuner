import boto3
import json

def pricing_for_region_list(response):
    results = []
    for price in response["PriceList"]:
        product = json.loads(price)['product']
        terms  = json.loads(price)['terms']

        if product['attributes']['regionCode'] == '':
            print("Empty region, skipping processing of price info: ")
            print(price)
            continue

        on_demand = terms['OnDemand']

        for _, element in on_demand.items():
            price_dimensions = element['priceDimensions']

            for _, dimension in price_dimensions.items():
                if dimension['beginRange'] == '0' and 'USD' in dimension['pricePerUnit']:
                    results.append({
                        "region": product['attributes']['regionCode'],
                        "price": dimension['pricePerUnit']['USD']
                    })

    return results

def print_results(results):
    for result in results:
        print('"' + result['region'] + '": '+result['price']+',')

#### X86

client = boto3.client('pricing', region_name="us-east-1")

response = client.get_products(
    ServiceCode = "AWSLambda",
    Filters=[
        {
            'Field': 'group',
            'Type': 'TERM_MATCH',
            'Value': 'AWS-Lambda-Duration',
        },
    ],
    FormatVersion='aws_v1',
    MaxResults=100,
)

x86_results = pricing_for_region_list(response)

#### ARM

response = client.get_products(
    ServiceCode = "AWSLambda",
    Filters=[
        {
            'Field': 'group',
            'Type': 'TERM_MATCH',
            'Value': 'AWS-Lambda-Duration-ARM',
        },
    ],
    FormatVersion='aws_v1',
    MaxResults=100,
)

arm_results = pricing_for_region_list(response)

####

print("------ x86 ------")
print_results(x86_results)


print("------ ARM ------")
print_results(arm_results)

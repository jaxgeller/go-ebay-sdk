import textwrap

import requests
from lxml import etree, html

typemap = {
    "string": 'string',
    "boolean": "bool",
    "int": "int",
    "dateTime": "*time.Time",
    "time": "*time.Time",
    "anyURI": "string",
    "token": "string",
    "double": "float64",
    "float": "float64",
    "long": "float64",
    "decimal": "float64",
    "struct": "struct"
}
methods = {}
items = {}
structs = ""


def converted_type(ebay_type, length):
    if ebay_type.endswith(')'):
        ebay_type = ebay_type.split('(')[1].replace(')', '')

    if ebay_type in typemap.keys():
        return typemap[ebay_type]

    # is typed enum like DisputeExplanationCodeType
    if length == 0:
        return 'string'

    return ebay_type


def print_type(val):
    if val.endswith('Type'):
        if 'Array' in val:
            val = '[]' + val
        else:
            val = '*' + val
    return val


def get_remaining_types():
    global items
    data = ""
    for k, v in items.items():
        data += 'type ' + k + \
            ' struct {\nXMLName   xml.Name`xml:"' + \
                k.replace('Type', '') + '"`\n'
        for kk, vv in v.items():
            vv = print_type(vv)
            data += kk + ' ' + vv + '`xml:",omitempty"`\n'
        data += '}\n\n'

    return data


# not my proudest function
def generate_types(el, struct, depth=0, switch='struct', search=None):
    global items
    for kid in el.getchildren():

        if '...' in kid.text:
            continue

        if ' ' in kid.text.strip() and 'Type' not in kid.text:
            continue

        key = kid.tag
        ebay_type = converted_type(kid.text.strip(), len(kid))

        if ebay_type.endswith('Type') and len(kid) > 0:
            items[ebay_type] = {}
            generate_types(kid, struct, depth=depth + 1,
                           switch='type', search=ebay_type)

        if depth == 0 and switch != 'type':
            ebay_type = print_type(ebay_type)
            struct += key + " " + ebay_type + '`xml:",omitempty"`\n'
        if switch == 'type' and search is not None:
            items[search][key] = ebay_type

        generate_types(kid, struct, depth=depth + 1)

    return struct + "}"


def get_typed(content):
    global structs
    struct = "\ntype " + content.tag + " struct {\nXMLName   xml.Name\n"
    if content.tag.endswith('Request'):
        struct += 'RequesterCredentials *RequesterCredentialsType\n'
    generated = generate_types(content, struct)
    structs += generated


def get_tree(page):
    r = requests.get(page)
    return html.fromstring(r.text)


def parse_call_index():
    call_index = 'https://developer.ebay.com/Devzone/XML/docs/Reference/eBay/index.html#CallIndex'
    doc = get_tree(call_index)
    table = doc.xpath('//table')[0]

    for method in table.xpath('//tr/td/a'):
        parse_page_for_xml(method.text)


def try_xml_from_pre(p):
    content = p.text_content()

    # no soap please
    if 'http://schemas.xmlsoap.org/soap/envelope/' in content:
        return None

    # wrong example pre
    if '<?xml version="1.0" encoding="utf-8"?>' not in content:
        return None

    content = content.replace('xmlns="urn:ebay:apis:eBLBaseComponents"', '')

    try:
        return etree.fromstring(content.strip().encode())
    except:
        return None


def parse_page_for_xml(method):
    page = 'https://developer.ebay.com/Devzone/XML/docs/Reference/eBay/{}.html'.format(
        method)
    doc = get_tree(page)

    methods[method] = {
        'desc': doc.xpath('//*[@id="doc"]/p[1]')[0].text_content()
    }

    for p in doc.xpath('//pre'):
        tree = try_xml_from_pre(p)
        if tree is not None:
            get_typed(tree)


def write_types():
    parse_call_index()
    data = get_remaining_types()
    f = open('trading_types.go', 'w')
    f.write('''package ebay

import (
    "encoding/xml"
    "time"
)''')
    f.write(structs)
    f.write('\n')
    f.write(data)
    f.write("""
type RequesterCredentialsType struct {
    EbayAuthToken string `xml:"eBayAuthToken"`
}
""")
    f.close()


def get_func(method, description):
    comment = "{} {}".format(method, description)
    comment = textwrap.wrap(comment, 80)
    comment = '// ' + '\n// '.join(comment)
    comment += '\n// https://developer.ebay.com/Devzone/XML/docs/Reference/eBay/{}.html'.format(
        method)
    f = '''{comment}
func (c EbayClient) {method}(payload *{method}Request) (*{method}Response, error) {{
        payload.RequesterCredentials = &RequesterCredentialsType{{EbayAuthToken: c.token}}
        response := &{method}{{}}
        err := c.execute("{method}", payload, response)
        return response, err
}}\n\n'''.format(method=method, comment=comment)
    return f


def write_methods():
    global methods
    f = open('trading_calls.go', 'w')
    for meth in methods.keys():
        func = get_func(meth, methods[meth]['desc'])
        f.write(func)
    f.close()


write_types()
write_methods()

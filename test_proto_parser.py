import proto_parser

with open("/home/lordpenguin/Work/codes/go/etcd/api/etcdserverpb/rpc.proto", 'r') as f:
    file = f.read()

parser = proto_parser.Interpreter().load_file(file)


# package = parser.package()
# print(package)
#
# imports = parser.imports()
# print(imports)
#
# options = parser.options()
# print(options)
#
# services = parser.services()
# # print(services)

# rpc, req, resp = parser.rpc('KV', 'Range')
# print(rpc)

# basic = parser.basic_info()
# print(basic)

fields = parser.message_elements('RangeRequest')